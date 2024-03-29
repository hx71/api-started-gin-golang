package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/app/jwtauth"
	"github.com/hx71/api-started-gin-golang/config"
	"github.com/hx71/api-started-gin-golang/response"
)

// AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT(jwtService jwtauth.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := response.ResponseError(config.MessageErr.FailedProcess, "token not found")
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		extractedToken := strings.Split(authHeader, "Bearer ")
		authHeader = strings.TrimSpace(extractedToken[1])
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			// claims := token.Claims.(jwt.MapClaims)
			// log.Println("Claim[user_id]: ", claims["user_id"])
			// log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			// log.Println(err)
			response := response.ResponseError("token is not valid", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
