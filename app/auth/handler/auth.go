package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/hx71/api-started-gin-golang/app/auth"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/config"
	"github.com/hx71/api-started-gin-golang/middleware"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
)

type AuthHandler interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	Register(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
}

type authHandler struct {
	Usecase auth.Usecase
	// Jwt     jwt.Usecase
}

func NewAuthHandler(usecase auth.Usecase) AuthHandler {
	return &authHandler{
		Usecase: usecase,
	}
}

func (u *authHandler) Login(ctx *gin.Context) {
	var credentials dto.LoginValidation
	errCredentials := ctx.ShouldBind(&credentials)
	if errCredentials != nil {
		response := response.ResponseError(config.MessageErr.FailedProcess, errCredentials.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := u.Usecase.VerifyCredential(credentials.Email, credentials.Password)
	if user, ok := authResult.(models.User); ok {
		generatedToken := middleware.GenerateToken(user.Email)
		tokenResponse := response.JwtResponse{
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

func (u *authHandler) Register(ctx *gin.Context) {
	var req dto.RegisterValidation
	req.ID = uuid.NewString()
	errRequest := ctx.ShouldBind(&req)
	if errRequest != nil {
		response := response.ResponseError(config.MessageErr.FailedProcess, errRequest.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !u.Usecase.FindByEmail(req.Email) {
		response := response.ResponseError(config.MessageErr.FailedProcess, "Duplicate email")
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := u.Usecase.CreateUser(req)
		response := response.ResponseSuccess("register successfull!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (u *authHandler) RefreshToken(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	extractedToken := strings.Split(authHeader, "Bearer ")
	authHeader = strings.TrimSpace(extractedToken[1])
	token, err := middleware.ValidateToken(authHeader)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		email := fmt.Sprintf("%s", claims["email"])
		refreshToken := middleware.GenerateToken(email)
		ctx.JSON(http.StatusOK, gin.H{"refresh_token": refreshToken})
	} else {
		response := response.ResponseError("Token is not valid", err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}
}

func (u *authHandler) Logout(ctx *gin.Context) {
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
