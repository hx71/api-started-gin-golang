package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/config"
	"github.com/hx71/api-started-gin-golang/response"
)

// PanicRecoveryHandler is a custom recovery handler for Gin that returns a standardized JSON response
func PanicRecoveryHandler(c *gin.Context, recovered interface{}) {
	if err, ok := recovered.(string); ok {
		res := response.ResponseError(config.MessageErr.FailedProcess, fmt.Sprintf("Internal Server Error: %s", err))
		c.JSON(http.StatusInternalServerError, res)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	res := response.ResponseError(config.MessageErr.FailedProcess, fmt.Sprintf("Internal Server Error: %v", recovered))
	c.JSON(http.StatusInternalServerError, res)
	c.AbortWithStatus(http.StatusInternalServerError)
}
