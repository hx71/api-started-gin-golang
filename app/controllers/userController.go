package controllers

import (
	"net/http"
	"srp-golang/app/models"
	"srp-golang/app/request"
	"srp-golang/helper"
	"srp-golang/service"
	"strconv"

	"github.com/gin-gonic/gin"
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

func (c *userController) Index(ctx *gin.Context) {
	var users []models.User = c.userService.Index()
	res := helper.BuildResponse(true, "List of user", users)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Create(ctx *gin.Context) {
	var userValidation request.UserCreateValidation
	errRequest := ctx.ShouldBind(&userValidation)
	if errRequest != nil {
		response := helper.BuildErrorResponse("Failed to process request", errRequest.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.userService.IsDuplicateEmail(userValidation.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		user := c.userService.Create(userValidation)
		response := helper.BuildResponse(true, "Created Success!", user)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *userController) Show(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var user models.User = c.userService.Show(id)
	if (user == models.User{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "Detail user", user)
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *userController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var user models.User = c.userService.Show(id)
	if (user == models.User{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		var userValidation request.UserUpdateValidation
		userValidation.ID = id
		errValidation := ctx.ShouldBind(&userValidation)
		if errValidation != nil {
			res := helper.BuildErrorResponse("Failed to process request", errValidation.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, res)
		}
		result := c.userService.Update(userValidation)
		res := helper.BuildResponse(true, "Updated success", result)
		ctx.JSON(http.StatusCreated, res)
	}
}

func (c *userController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var user models.User = c.userService.Show(id)
	if (user == models.User{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		result := c.userService.Delete(user)
		res := helper.BuildResponse(true, "Deleted success", result)
		ctx.JSON(http.StatusCreated, res)
	}
}
