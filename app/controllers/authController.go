package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/dto"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/service"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/models"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/response"
)

// AuthController interface is a contract what this controller can do
type AuthController interface {
	Version(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
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

// NewAuthController creates a new instance of AuthController
func NewAuthController(authServ service.AuthService, jwtServ service.JWTService) AuthController {
	return &authController{
		authService: authServ,
		jwtService:  jwtServ,
	}
}

func (c *authController) Version(ctx *gin.Context) {
	response := response.ResponseSuccess("api version", "version 1.0.0")
	ctx.JSON(http.StatusOK, response)
}

func (s *authController) Login(ctx *gin.Context) {
	var credentials dto.LoginValidation
	errCredentials := ctx.ShouldBind(&credentials)
	if errCredentials != nil {
		response := response.ResponseError("Failed to process request", errCredentials.Error())
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
		response := response.ResponseSuccess("login successfull!", tokenResponse)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := response.ResponseError("Please check again your credential", "Invalid Credential")
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (s *authController) Register(ctx *gin.Context) {
	var req dto.RegisterValidation
	req.ID = uuid.NewString()
	errRequest := ctx.ShouldBind(&req)
	if errRequest != nil {
		response := response.ResponseError("Failed to process request", errRequest.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !s.authService.FindByEmail(req.Email) {
		response := response.ResponseError("Failed to process request", "Duplicate email")
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := s.authService.CreateUser(req)
		response := response.ResponseSuccess("register successfull!", createdUser)
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
		response := response.ResponseError("Token is not valid", err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}
}

func (s *authController) Logout(ctx *gin.Context) {
	cookies := http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-7 * 24 * time.Hour),
		MaxAge:  -1,
	}
	http.SetCookie(ctx.Writer, &cookies)

	res := response.ResponseSuccess("Successfully logged out!", cookies)
	ctx.JSON(http.StatusCreated, res)
}
