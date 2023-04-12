package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags Api Version
// API Version : passing specific parameter to DBaaS from Service Portal
// @Accept  json
// @Produce  json
// @Success 200 {object} engine.ResponseSuccess
// @Failure 400 {object} engine.ResponseStatus
// / @Failure 404 {object} engine.ResponseStatus
// / @Failure 500 {object} engine.ResponseStatus
// @Router /version [get]
func ApiVersion(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Roles
// API Roles : passing specific parameter to DBaaS from Service Portal
// @Accept  json
// @Produce  json
// @Router /api/v1/roles [get]
func IndexRoles(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Roles
// API Roles : passing specific parameter to DBaaS from Service Portal
// @Accept  json
// @Produce  json
// @Router /api/v1/roles [post]
func CreateRoles(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Roles
// API Roles : passing specific parameter to DBaaS from Service Portal
// @Accept  json
// @Produce  json
// @Router /api/v1/roles/{id} [get]
func ShowRoles(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Roles
// API Roles : passing specific parameter to DBaaS from Service Portal
// @Accept  json
// @Produce  json
// @Router /api/v1/roles/{id} [put]
func UpdateRoles(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Roles
// API Roles : passing specific parameter to DBaaS from Service Portal
// @Accept  json
// @Produce  json
// @Router /api/v1/roles/{id} [delete]
func DeleteRoles(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
