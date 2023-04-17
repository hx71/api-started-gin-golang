package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags UserMenu
// API UserMenu : passing specific parameter to DBaaS from Service
// @Security BearerAuth
// @Param limit query integer false "Limit Per Page"
// @Param page query integer false "Page Number"
// @Param sort query string false "Sort By {ex: created_at asc | desc}"
// @Param id.equals query string false "Seraching by column {ex: id} action {ex: equals | contains | in}"
// @Accept  json
// @Produce  json
// @Router /api/v1/user-menus [get]
func IndexUserMenu(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags UserMenu
// API UserMenu : passing specific parameter to DBaaS from Service
// @Security BearerAuth
// @Param filter body []engine.UserMenu true "Pass session information to DBaaS Parameter"
// @Accept  json
// @Produce  json
// @Router /api/v1/user-menus [post]
func CreateUserMenu(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags UserMenu
// API UserMenu : passing specific parameter to DBaaS from Service
// @Security BearerAuth
// @Param id path string true "Pass session information to DBaaS Parameter"
// @Accept  json
// @Produce  json
// @Router /api/v1/user-menus/{id} [get]
func ShowUserMenu(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags UserMenu
// API UserMenu : passing specific parameter to DBaaS from Service
// @Security BearerAuth
// @Param id path string true "Pass session information to DBaaS Parameter"
// @Param filter body engine.UserMenu true "Pass session information to DBaaS Parameter"
// @Accept  json
// @Produce  json
// @Router /api/v1/user-menus/{id} [put]
func UpdateUserMenu(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags UserMenu
// API UserMenu : passing specific parameter to DBaaS from Service
// @Security BearerAuth
// @Param id path string true "Pass session information to DBaaS Parameter"
// @Accept  json
// @Produce  json
// @Router /api/v1/user-menus/{id} [delete]
func DeleteUserMenu(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
