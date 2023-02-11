package service

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/dto"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/repository"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/helpers"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/models"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/response"
	"github.com/mashingan/smapping"
)

//RoleService is a ....
type RoleService interface {
	Create(req dto.RoleCreateValidation) error
	Show(id string) models.Roles
	Update(req dto.RoleCreateValidation) error
	Delete(id string) error
	Pagination(ctx *gin.Context, pagination *helpers.Pagination) response.Response
}

type roleService struct {
	roleRepository repository.RoleRepository
}

//NewRoleService .....
func NewRoleService(roleRepo repository.RoleRepository) RoleService {
	return &roleService{
		roleRepository: roleRepo,
	}
}

func (service *roleService) Create(req dto.RoleCreateValidation) error {
	role := models.Role{}
	role.ID = uuid.NewString()
	err := smapping.FillStruct(&role, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	return service.roleRepository.Create(role)
}

func (service *roleService) Show(id string) models.Roles {
	return service.roleRepository.Show(id)
}

func (service *roleService) Update(req dto.RoleCreateValidation) error {
	role := models.Role{}
	err := smapping.FillStruct(&role, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	return service.roleRepository.Update(role)
}

func (service *roleService) Delete(id string) error {
	return service.roleRepository.Delete(id)
}

func (r *roleService) Pagination(context *gin.Context, pagination *helpers.Pagination) response.Response {

	operationResult, totalPages := r.roleRepository.Pagination(pagination)

	if operationResult.Error != nil {
		return response.Response{Status: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*helpers.Pagination)

	// get current url path
	urlPath := context.Request.URL.Path

	// search query params
	searchQueryParams := ""

	for _, search := range pagination.Searchs {
		searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
	}

	// set first & last page pagination response
	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, 1, pagination.Sort) + searchQueryParams
	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, totalPages, pagination.Sort) + searchQueryParams

	if data.Page > 1 {
		// set previous page pagination response
		data.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page, pagination.Sort) + searchQueryParams
	}

	if data.Page < totalPages {
		// set next page pagination response
		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page+1, pagination.Sort) + searchQueryParams
	}

	if data.Page > totalPages {
		// reset previous page
		data.PreviousPage = ""
	}

	return response.Response{Status: true, Data: data}
}
