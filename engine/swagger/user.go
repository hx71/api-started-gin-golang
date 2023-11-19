package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags Users
// API Users : passing specific parameter to DBaaS from Service
// @Security BearerAuth
// @Param limit query integer false "Limit Per Page"
// @Param page query integer false "Page Number"
// @Param sort query string false "Sort By {ex: created_at asc | desc}"
// @Param id.equals query string false "Seraching by column {ex: id} action {ex: equals | contains | in}"
// @Accept  json
// @Produce  json
// @Router /api/v1/users [get]
func IndexUsers(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Users
// API Users : passing specific parameter to DBaaS from Service
// @Security BearerAuth
// @Param filter body engine.User true "Pass session information to DBaaS Parameter"
// @Accept  json
// @Produce  json
// @Router /api/v1/users [post]
func CreateUsers(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Users
// API Users : passing specific parameter to DBaaS from Service
// @Security BearerAuth
// @Param id path string true "Pass session information to DBaaS Parameter"
// @Accept  json
// @Produce  json
// @Router /api/v1/users/{id} [get]
func ShowUsers(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Users
// API Users : passing specific parameter to DBaaS from Service
// @Security BearerAuth
// @Param id path string true "Pass session information to DBaaS Parameter"
// @Param filter body engine.User true "Pass session information to DBaaS Parameter"
// @Accept  json
// @Produce  json
// @Router /api/v1/users/{id} [put]
func UpdateUsers(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Users
// API Users : passing specific parameter to DBaaS from Service
// @Security BearerAuth
// @Param id path string true "Pass session information to DBaaS Parameter"
// @Accept  json
// @Produce  json
// @Router /api/v1/users/{id} [delete]
func DeleteUsers(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
