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

//UserService is a ....
type UserService interface {
	Index() []models.User
	Create(model dto.UserCreateValidation) error
	Show(id string) models.User
	Update(model dto.UserUpdateValidation) error
	Delete(model models.User) error
	FindByEmail(email string) bool
	Pagination(ctx *gin.Context, pagination *helpers.Pagination) response.Response
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

func (service *userService) Create(model dto.UserCreateValidation) error {
	user := models.User{}
	user.ID = uuid.NewString()
	err := smapping.FillStruct(&user, smapping.MapFields(&model))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	return service.userRepository.Create(user)
}

func (service *userService) Show(id string) models.User {
	return service.userRepository.Show(id)
}

func (service *userService) Update(model dto.UserUpdateValidation) error {
	user := models.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&model))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	return service.userRepository.Update(user)
}

func (service *userService) Delete(user models.User) error {
	return service.userRepository.Delete(user)
}

func (service *userService) FindByEmail(email string) bool {
	return service.userRepository.FindByEmail(email)
}

func (r *userService) Pagination(context *gin.Context, pagination *helpers.Pagination) response.Response {

	operationResult, totalPages := r.userRepository.Pagination(pagination)

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
