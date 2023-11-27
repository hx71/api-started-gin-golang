package service

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/app/repository"
	"github.com/hx71/api-started-gin-golang/helpers"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
	"github.com/mashingan/smapping"
)

//MenuService is a ....
type MenuService interface {
	Create(req dto.MenuCreateValidation) error
	Show(id string) models.Menus
	Update(req dto.MenuCreateValidation) error
	Delete(id string) error
	Pagination(ctx *gin.Context, pagination *helpers.Pagination) response.Response
}

type menuService struct {
	menuRepository repository.MenuRepository
}

//NewMenuService .....
func NewMenuService(menuRepo repository.MenuRepository) MenuService {
	return &menuService{
		menuRepository: menuRepo,
	}
}

func (service *menuService) Create(req dto.MenuCreateValidation) error {
	menu := models.Menu{}
	menu.ID = uuid.NewString()
	err := smapping.FillStruct(&menu, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	return service.menuRepository.Create(menu)
}

func (service *menuService) Show(id string) models.Menus {
	return service.menuRepository.Show(id)
}

func (service *menuService) Update(req dto.MenuCreateValidation) error {
	menu := models.Menu{}
	err := smapping.FillStruct(&menu, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	return service.menuRepository.Update(menu)
}

func (service *menuService) Delete(id string) error {
	return service.menuRepository.Delete(id)
}

func (r *menuService) Pagination(context *gin.Context, pagination *helpers.Pagination) response.Response {

	operationResult, totalPages := r.menuRepository.Pagination(pagination)

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
