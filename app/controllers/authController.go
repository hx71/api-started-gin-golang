package controllers

import (
	"fmt"
	"log"
	"net/http"
	"srp-golang/app/models"
	"srp-golang/app/request"
	"srp-golang/helper"
	"srp-golang/service"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}
type LoginResponse struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
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

func (c *authController) RefreshToken(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	extractedToken := strings.Split(authHeader, "Bearer ")
	authHeader = strings.TrimSpace(extractedToken[1])
	token, err := c.jwtService.ValidateToken(authHeader)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		email := fmt.Sprintf("%s", claims["email"])
		refresh_token := c.jwtService.GenerateToken(email)
		ctx.JSON(http.StatusOK, gin.H{"refresh_token": refresh_token})
	} else {
		log.Println(err)
		response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}
}
