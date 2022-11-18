package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/dto"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/service"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/helpers"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/models"
)

//AuthController interface is a contract what this controller can do
type AuthController interface {
	Version(ctx *gin.Context)
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}
type LoginResponse struct {
	ID          string `json:"id"`
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

func (c *authController) Version(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "version 1.0.0")
}

func (s *authController) Login(ctx *gin.Context) {
	var credentials dto.LoginValidation
	errCredentials := ctx.ShouldBind(&credentials)
	if errCredentials != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errCredentials.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := s.authService.VerifyCredential(credentials.Email, credentials.Password)
	if user, ok := authResult.(models.User); ok {
		generatedToken := s.jwtService.GenerateToken(user.Email)
		tokenResponse := LoginResponse{
			ID:          user.ID,
			Name:        user.Name,
			Username:    user.Username,
			Email:       user.Email,
			AccessToken: generatedToken,
		}
		response := helpers.BuildResponse(true, "login successfull!", tokenResponse)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helpers.BuildErrorResponse("Please check again your credential", "Invalid Credential", helpers.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (s *authController) Register(ctx *gin.Context) {
	var req dto.RegisterValidation
	req.ID = uuid.NewString()
	errRequest := ctx.ShouldBind(&req)
	if errRequest != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errRequest.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if s.authService.FindByEmail(req.Email) != nil {
		response := helpers.BuildErrorResponse("Failed to process request", "Duplicate email", helpers.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := s.authService.CreateUser(req)
		response := helpers.BuildResponse(true, "register successfull!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (s *authController) RefreshToken(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	extractedToken := strings.Split(authHeader, "Bearer ")
	authHeader = strings.TrimSpace(extractedToken[1])
	token, err := s.jwtService.ValidateToken(authHeader)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		email := fmt.Sprintf("%s", claims["email"])
		refresh_token := s.jwtService.GenerateToken(email)
		ctx.JSON(http.StatusOK, gin.H{"refresh_token": refresh_token})
	} else {
		log.Println(err)
		response := helpers.BuildErrorResponse("Token is not valid", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}
}
