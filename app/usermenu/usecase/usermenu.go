package r

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/app/usermenu"
	"github.com/hx71/api-started-gin-golang/helpers"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
	"github.com/mashingan/smapping"
)

type userMenuUsecase struct {
	repo usermenu.Repository
}

func NewUserMenuUsecase(repo usermenu.Repository) usermenu.Usecase {
	return &userMenuUsecase{
		repo: repo,
	}
}

func (r *userMenuUsecase) Create(req []dto.UserMenuCreateValidation) error {
	userMenu := models.UserMenu{}
	for _, v := range req {
		userMenu.ID = uuid.NewString()
		err := smapping.FillStruct(&userMenu, smapping.MapFields(&v))
		if err != nil {
			log.Fatalf("Failed map %v: ", err)
		}
		return r.repo.Create(userMenu)
	}
	return nil
}

func (r *userMenuUsecase) Show(id string) models.UserMenus {
	return r.repo.Show(id)
}

func (r *userMenuUsecase) Update(id string, req dto.UserMenuCreateValidation) error {
	userMenu := models.UserMenu{}
	userMenu.ID = id
	err := smapping.FillStruct(&userMenu, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	return r.repo.Update(userMenu)
}

func (r *userMenuUsecase) Delete(model models.UserMenus) error {
	return r.repo.Delete(model)
}

func (r *userMenuUsecase) Pagination(context *gin.Context, pagination *helpers.Pagination) response.Response {

	operationResult, totalPages := r.repo.Pagination(pagination)

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
