package controllers

import (
	"net/http"
	"srp-golang/app/request"
	"srp-golang/helper"
	"srp-golang/service"

	"github.com/gin-gonic/gin"
)

//AuthController interface is a contract what this controller can do
type AuthController interface {
	// Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	// jwtService  service.JWTService
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

func (c *authController) Register(ctx *gin.Context) {
	var RegisterRequest request.RegisterValidation
	errDTO := ctx.ShouldBind(&RegisterRequest)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// if !c.authService.IsDuplicateEmail(RegisterRequest.Email) {
	// 	response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
	// 	ctx.JSON(http.StatusConflict, response)
	// } else {
	createdUser := c.authService.CreateUser(RegisterRequest)
	// token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
	// createdUser.Token = token
	response := helper.BuildResponse(true, "register successfull!", createdUser)
	ctx.JSON(http.StatusCreated, response)
	// }
}
