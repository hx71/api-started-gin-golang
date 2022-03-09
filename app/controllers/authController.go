package controllers

import (
	"net/http"
	"srp-golang/app/models"
	"srp-golang/app/request"
	"srp-golang/helper"
	"srp-golang/service"

	"github.com/gin-gonic/gin"
)

//AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}
type LoginResponse struct {
	ID          uint64
	Name        string
	Username    string
	Email       string
	AccessToken string `json:"access_token"`
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(authServ service.AuthService, jwtServ service.JWTService) AuthController {
	return &authController{
		authService: authServ,
		jwtService:  jwtServ,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var credentials request.LoginValidation
	errCredentials := ctx.ShouldBind(&credentials)
	if errCredentials != nil {
		response := helper.BuildErrorResponse("Failed to process request", errCredentials.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(credentials.Email, credentials.Password)
	if user, ok := authResult.(models.User); ok {
		generatedToken := c.jwtService.GenerateToken(user.Email)
		// generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(user.ID, 10))
		tokenResponse := LoginResponse{
			ID:          user.ID,
			Name:        user.Name,
			Username:    user.Username,
			Email:       user.Email,
			AccessToken: generatedToken,
		}
		response := helper.BuildResponse(true, "OK!", tokenResponse)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var RegisterRequest request.RegisterValidation
	errRequest := ctx.ShouldBind(&RegisterRequest)
	if errRequest != nil {
		response := helper.BuildErrorResponse("Failed to process request", errRequest.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(RegisterRequest.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(RegisterRequest)
		response := helper.BuildResponse(true, "register successfull!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
