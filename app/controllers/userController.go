package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/dto"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/service"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/helpers"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/models"
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
	code := http.StatusOK
	pagination := helpers.GeneratePaginationRequest(ctx)
	response := s.userService.PaginationUser(ctx, pagination)
	if !response.Success {
		code = http.StatusBadRequest
	}
	ctx.JSON(code, response)
}

func (s *userController) Create(ctx *gin.Context) {
	var userValidation dto.UserCreateValidation
	errRequest := ctx.ShouldBind(&userValidation)
	if errRequest != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errRequest.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !s.userService.IsDuplicateEmail(userValidation.Email) {
		response := helpers.BuildErrorResponse("Failed to process request", "Duplicate email", helpers.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		user := s.userService.Create(userValidation)
		response := helpers.BuildResponse(true, "Created Success!", user)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (s *userController) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User = s.userService.Show(id)
	if (user == models.User{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "Detail user", user)
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *userController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User = c.userService.Show(id)
	if (user == models.User{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		var userValidation dto.UserUpdateValidation
		userValidation.ID = id
		errValidation := ctx.ShouldBind(&userValidation)
		if errValidation != nil {
			res := helpers.BuildErrorResponse("Failed to process request", errValidation.Error(), helpers.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, res)
		}
		result := c.userService.Update(userValidation)
		res := helpers.BuildResponse(true, "Updated success", result)
		ctx.JSON(http.StatusCreated, res)
	}
}

func (c *userController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	var user models.User = c.userService.Show(id)
	if (user == models.User{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		result := c.userService.Delete(user)
		res := helpers.BuildResponse(true, "Deleted success", result)
		ctx.JSON(http.StatusCreated, res)
	}
}
