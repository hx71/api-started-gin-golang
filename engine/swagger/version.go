package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags Api Version
// API Version : passing specific parameter to DBaaS from Service
// @Accept  json
// @Produce  json
// @Success 200 {object} engine.ResponseSuccess
// @Failure 400 {object} engine.ResponseStatus
// / @Failure 404 {object} engine.ResponseStatus
// / @Failure 500 {object} engine.ResponseStatus
// @Router /api/v1/version [get]
func ApiVersion(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
