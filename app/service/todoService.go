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

//TodoService is a ....
type TodoService interface {
	Create(model dto.TodoCreateValidation) error
	Show(id string) models.Todos
	Update(model dto.TodoCreateValidation) error
	Delete(id string) error
	Pagination(ctx *gin.Context, pagination *helpers.Pagination) response.Response
}

type todoService struct {
	todoRepository repository.TodoRepository
}

//NewTodoService .....
func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: todoRepo,
	}
}

func (service *todoService) Create(model dto.TodoCreateValidation) error {
	todo := models.Todo{}
	todo.ID = uuid.NewString()
	err := smapping.FillStruct(&todo, smapping.MapFields(&model))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	return service.todoRepository.Create(todo)
}

func (service *todoService) Show(id string) models.Todos {
	return service.todoRepository.Show(id)
}

func (service *todoService) Update(model dto.TodoCreateValidation) error {
	todo := models.Todo{}
	err := smapping.FillStruct(&todo, smapping.MapFields(&model))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	return service.todoRepository.Update(todo)
}

func (service *todoService) Delete(id string) error {
	return service.todoRepository.Delete(id)
}

func (r *todoService) Pagination(context *gin.Context, pagination *helpers.Pagination) response.Response {

	operationResult, totalPages := r.todoRepository.Pagination(pagination)

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
