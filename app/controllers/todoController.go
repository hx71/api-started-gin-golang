package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/dto"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/service"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/helpers"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/models"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/response"
)

// TodoController is a contract what this controller can do
type TodoController interface {
	Index(ctx *gin.Context)
	Create(ctx *gin.Context)
	Show(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type todoController struct {
	todoService service.TodoService
	jwtService  service.JWTService
}

// NewTodoController create a new instances of TodoController
func NewTodoController(todoServ service.TodoService, jwtServ service.JWTService) TodoController {
	return &todoController{
		todoService: todoServ,
		jwtService:  jwtServ,
	}
}

func (s *todoController) Index(ctx *gin.Context) {
	pagination := helpers.GeneratePaginationRequest(ctx)
	res := s.todoService.Pagination(ctx, pagination)
	if !res.Status {
		response := response.ResponseError("failed to get data todo", res.Message)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.ResponseSuccess("list of todo", res.Data)
	ctx.JSON(http.StatusOK, response)
}

func (s *todoController) Create(ctx *gin.Context) {
	var req dto.TodoCreateValidation
	req.ID = uuid.NewString()
	err := ctx.ShouldBind(&req)
	if err != nil {
		response := response.ResponseError("failed to process request", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user := s.todoService.Create(req)
	response := response.ResponseSuccess("created succeess", user)
	ctx.JSON(http.StatusCreated, response)
}

func (s *todoController) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	var todo models.Todos = s.todoService.Show(id)
	if todo.ID == "" {
		res := response.ResponseError("Data not found", "No data with given id")
		ctx.JSON(http.StatusNotFound, res)
	} else {
		response := response.ResponseSuccess("detail user", todo)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *todoController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var todo models.Todos = c.todoService.Show(id)
	if todo.ID == "" {
		res := response.ResponseError("data not found", "no data with given id")
		ctx.JSON(http.StatusNotFound, res)
	} else {
		var req dto.TodoCreateValidation
		req.ID = id
		err := ctx.ShouldBind(&req)
		if err != nil {
			response := response.ResponseError("failed to process request", err.Error())
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		err = c.todoService.Update(req)
		if err != nil {
			response := response.ResponseError("update failed", err.Error())
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		response := response.ResponseSuccess("update success", nil)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *todoController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response := response.ResponseError("data not found", "no data with given id")
		ctx.JSON(http.StatusNotFound, response)
	} else {
		todo := c.todoService.Delete(id)
		response := response.ResponseSuccess("deleted success", todo)
		ctx.JSON(http.StatusOK, response)
	}
}
