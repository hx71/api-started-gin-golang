package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/dto"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/service"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/helpers"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/models"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/response"
)

type Paginatex struct {
	Meta map[string]interface{} `json:"meta"`
	Data interface{}            `json:"data"`
}

//UserController is a contract what this controller can do
type UserController interface {
	Index(ctx *gin.Context)
	Create(ctx *gin.Context)
	Show(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

//NewUserController create a new instances of UserController
func NewUserController(userServ service.UserService, jwtServ service.JWTService) UserController {
	return &userController{
		userService: userServ,
		jwtService:  jwtServ,
	}
}

func (s *userController) Index(ctx *gin.Context) {
	pagination := helpers.GeneratePaginationRequest(ctx)
	res := s.userService.PaginationUser(ctx, pagination)
	if !res.Status {
		response := response.ResponseError("failed to get data user", res.Message, response.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.ResponseSuccess("list of user", res.Data)
	ctx.JSON(http.StatusOK, response)
}

func (s *userController) Create(ctx *gin.Context) {
	var userValidation dto.UserCreateValidation
	err := ctx.ShouldBind(&userValidation)
	if err != nil {
		response := response.ResponseError("failed to process request", err.Error(), response.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if !s.userService.IsDuplicateEmail(userValidation.Email) {
		response := response.ResponseError("Failed to process request", "Duplicate email", helpers.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		user := s.userService.Create(userValidation)
		response := response.ResponseSuccess("created succeess", user)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (s *userController) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User = s.userService.Show(id)
	if (user == models.User{}) {
		res := response.ResponseError("Data not found", "No data with given id", helpers.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		response := response.ResponseSuccess("detail user", user)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *userController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User = c.userService.Show(id)
	if (user == models.User{}) {
		res := response.ResponseError("Data not found", "No data with given id", helpers.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		var userValidation dto.UserUpdateValidation
		userValidation.ID = id
		err := ctx.ShouldBind(&userValidation)
		if err != nil {
			response := response.ResponseError("Failed to process request", err.Error(), response.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		user := c.userService.Update(userValidation)
		response := response.ResponseSuccess("update success", user)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *userController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User = c.userService.Show(id)
	if (user == models.User{}) {
		response := response.ResponseError("data not found", "No data with given id", response.EmptyObj{})
		ctx.JSON(http.StatusNotFound, response)
	} else {
		user := c.userService.Delete(user)
		response := response.ResponseSuccess("deleted success", user)
		ctx.JSON(http.StatusOK, response)
	}
}
