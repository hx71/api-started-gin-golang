package service

import (
	"fmt"
	"log"
	"srp-golang/app/models"
	"srp-golang/app/request"
	"srp-golang/repository"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

//UserService is a ....
type UserService interface {
	Index() []models.User
	Create(model request.UserCreateValidation) models.User
	Show(id uint64) models.User
	Update(model request.UserUpdateValidation) models.User
	Delete(model models.User) models.User
	FindByEmail(email string) models.User
	IsDuplicateEmail(email string) bool
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

func (service *userService) Create(model request.UserCreateValidation) models.User {
	user := models.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&model))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.userRepository.Create(user)
	return res
}

func (service *userService) Show(id uint64) models.User {
	return service.userRepository.Show(id)
}

func (service *userService) Update(model request.UserUpdateValidation) models.User {
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

func PaginationUser(repository repository.UserRepository, context *gin.Context, pagination *request.Pagination) request.Response {

	operationResult, totalPages := repository.PaginationUser(pagination)

	if operationResult.Error != nil {
		return request.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*request.Pagination)

	// get current url path
	urlPath := context.Request.URL.Path

	// search query params
	searchQueryParams := ""

	for _, search := range pagination.Searchs {
		searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
	}

	// set first & last page pagination response
	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, 0, pagination.Sort) + searchQueryParams
	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, totalPages, pagination.Sort) + searchQueryParams

	if data.Page > 0 {
		// set previous page pagination response
		data.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page-1, pagination.Sort) + searchQueryParams
	}

	if data.Page < totalPages {
		// set next page pagination response
		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page+1, pagination.Sort) + searchQueryParams
	}

	if data.Page > totalPages {
		// reset previous page
		data.PreviousPage = ""
	}

	return request.Response{Success: true, Data: data}
}
