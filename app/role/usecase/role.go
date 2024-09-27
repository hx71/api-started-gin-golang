package usecase

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/app/role"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
	"github.com/mashingan/smapping"
)

type roleUsecase struct {
	repo role.Repository
}

func NewRoleUsecase(repo role.Repository) role.Usecase {
	return &roleUsecase{
		repo: repo,
	}
}

func (r *roleUsecase) Create(req dto.RoleCreateValidation) error {
	role := models.Role{
		ID: uuid.NewString(),
	}
	err := smapping.FillStruct(&role, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	return r.repo.Create(role)
}

func (r *roleUsecase) Show(id string) models.Role {
	return r.repo.Show(id)
}

func (r *roleUsecase) Delete(model models.Role) error {
	return r.repo.Delete(model)
}

func (r *roleUsecase) Update(req dto.RoleCreateValidation) error {
	role := models.Role{}
	err := smapping.FillStruct(&role, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	return r.repo.Update(role)
}

func (r *roleUsecase) Pagination(context *gin.Context, pagination *response.Pagination) response.Response {

	operationResult, totalPages := r.repo.Pagination(pagination)

	if operationResult.Error != nil {
		return response.Response{Status: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*response.Pagination)

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
