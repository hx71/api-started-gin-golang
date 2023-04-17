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

//UserMenuService is a ....
type UserMenuService interface {
	Create(req []dto.UserMenuCreateValidation) error
	Show(id string) models.UserMenus
	Update(id string, req dto.UserMenuCreateValidation) error
	Delete(model models.UserMenus) error
	Pagination(ctx *gin.Context, pagination *helpers.Pagination) response.Response
}

type userMenuService struct {
	userMenuRepository repository.UserMenuRepository
}

//NewUserMenuService .....
func NewUserMenuService(userMenuRepo repository.UserMenuRepository) UserMenuService {
	return &userMenuService{
		userMenuRepository: userMenuRepo,
	}
}

func (service *userMenuService) Create(req []dto.UserMenuCreateValidation) error {
	userMenu := models.UserMenu{}
	for _, v := range req {
		userMenu.ID = uuid.NewString()
		err := smapping.FillStruct(&userMenu, smapping.MapFields(&v))
		if err != nil {
			log.Fatalf("Failed map %v: ", err)
		}
		return service.userMenuRepository.Create(userMenu)
	}
	return nil
}

func (service *userMenuService) Show(id string) models.UserMenus {
	return service.userMenuRepository.Show(id)
}

func (service *userMenuService) Update(id string, req dto.UserMenuCreateValidation) error {
	userMenu := models.UserMenu{}
	userMenu.ID = id
	err := smapping.FillStruct(&userMenu, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	return service.userMenuRepository.Update(userMenu)
}

func (service *userMenuService) Delete(model models.UserMenus) error {
	return service.userMenuRepository.Delete(model)
}

func (r *userMenuService) Pagination(context *gin.Context, pagination *helpers.Pagination) response.Response {

	operationResult, totalPages := r.userMenuRepository.Pagination(pagination)

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
