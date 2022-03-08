package middleware

import (
	"log"
	"net/http"
	"srp-golang/helper"
	"srp-golang/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//AuthorizeJWT validates the token user given, return 401 if not valid
// func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		clientToken := c.Request.Header.Get("Authorization")
// 		if clientToken == "" {
// 			c.JSON(403, "No Authorization header provided")
// 			c.Abort()
// 			return
// 		}

// 		// token, err := jwtService.ValidateToken(clientToken)
// 		// if err != nil {
// 		// 	panic(err)
// 		// }
// 		// c.JSON(http.StatusOK, gin.H{"token": clientToken})

// 		extractedToken := strings.Split(clientToken, "Bearer ")

// 		if len(extractedToken) == 2 {
// 			clientToken = strings.TrimSpace(extractedToken[1])
// 		} else {
// 			c.JSON(400, "Incorrect Format of Authorization Token")
// 			c.Abort()
// 			return
// 		}

// 		token, err := jwtService.ValidateToken(clientToken)
// 		// if token.Valid {
// 		// 	claims := token.Claims.(jwt.MapClaims)
// 		// 	log.Println("Claim[user_id]: ", claims["user_id"])
// 		// 	log.Println("Claim[issuer] :", claims["issuer"])
// 		// } else {
// 		// 	log.Println(err)
// 		// 	response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
// 		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 		// }

// 		// err := jwtService.ValidateToken(clientToken)
// 		if err != nil {
// 			c.JSON(401, err.Error())
// 			c.Abort()
// 			return
// 		}

// 		// c.Next()

// 		// c.JSON(http.StatusOK, gin.H{"token": clientToken})
// 		c.JSON(http.StatusOK, gin.H{"token": token})
// 		// return

// 		// jwtWrapper := service.JwtWrapper{
// 		// 	SecretKey: "verysecretkey",
// 		// 	Issuer:    "AuthService",
// 		// }

// 		// claims, err := jwtWrapper.ValidateToken(clientToken)
// 		// if err != nil {
// 		// 	c.JSON(401, err.Error())
// 		// 	c.Abort()
// 		// 	return
// 		// }

// 		// c.Set("email", claims.Email)

// 		// c.Next()

// 		// if token.Valid {
// 		// 	claims := token.Claims.(jwt.MapClaims)
// 		// 	log.Println("Claim[user_id]: ", claims["user_id"])
// 		// 	log.Println("Claim[issuer] :", claims["issuer"])
// 		// } else {
// 		// 	log.Println(err)
// 		// 	response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
// 		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 		// }
// 	}
// }
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
