package service

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/dto"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/repository"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/helpers"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/models"
	"github.com/mashingan/smapping"
)

//UserService is a ....
type UserService interface {
	Index() []models.User
	Create(model dto.UserCreateValidation) models.User
	Show(id string) models.User
	Update(model dto.UserUpdateValidation) models.User
	Delete(model models.User) models.User
	FindByEmail(email string) models.User
	IsDuplicateEmail(email string) bool

	PaginationUser(ctx *gin.Context, pagination *helpers.Pagination) models.Response
}

type userService struct {
	userRepository repository.UserRepository
}

//NewUserService .....
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Index() []models.User {
	return service.userRepository.Index()
}

func (service *userService) Create(model dto.UserCreateValidation) models.User {
	user := models.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&model))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.userRepository.Create(user)
	return res
}

func (service *userService) Show(id string) models.User {
	return service.userRepository.Show(id)
}

func (service *userService) Update(model dto.UserUpdateValidation) models.User {
	user := models.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&model))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.userRepository.Update(user)
	return res
}

func (service *userService) Delete(user models.User) models.User {
	return service.userRepository.Delete(user)
}

func (service *userService) FindByEmail(email string) models.User {
	return service.userRepository.FindByEmail(email)
}

func (service *userService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func (r *userService) PaginationUser(context *gin.Context, pagination *helpers.Pagination) models.Response {

	operationResult, totalPages := r.userRepository.PaginationUser(pagination)

	if operationResult.Error != nil {
		return models.Response{Success: false, Message: operationResult.Error.Error()}
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

	return models.Response{Success: true, Data: data}
}
