package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags Roles
// API Roles : passing specific parameter to DBaaS from Service
// @Security BearerAuth
// @Param limit query integer false "Limit Per Page"
// @Param page query integer false "Page Number"
// @Param sort query string false "Sort By {ex: created_at asc | desc}"
// @Param id.equals query string false "Seraching by column {ex: id} action {ex: equals | contains | in}"
// @Accept  json
// @Produce  json
// @Router /api/v1/roles [get]
func IndexRoles(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Roles
// API Roles : passing specific parameter to DBaaS from Service
// @Security BearerAuth
// @Param filter body engine.Role true "Pass session information to DBaaS Parameter"
// @Accept  json
// @Produce  json
// @Router /api/v1/roles [post]
func CreateRoles(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Roles
// API Roles : passing specific parameter to DBaaS from Service
// @Security BearerAuth
// @Param id path string true "Pass session information to DBaaS Parameter"
// @Accept  json
// @Produce  json
// @Router /api/v1/roles/{id} [get]
func ShowRoles(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Roles
// API Roles : passing specific parameter to DBaaS from Service
// @Security BearerAuth
// @Param id path string true "Pass session information to DBaaS Parameter"
// @Param filter body engine.Role true "Pass session information to DBaaS Parameter"
// @Accept  json
// @Produce  json
// @Router /api/v1/roles/{id} [put]
func UpdateRoles(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Roles
// API Roles : passing specific parameter to DBaaS from Service
// @Security BearerAuth
// @Param id path string true "Pass session information to DBaaS Parameter"
// @Accept  json
// @Produce  json
// @Router /api/v1/roles/{id} [delete]
func DeleteRoles(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
