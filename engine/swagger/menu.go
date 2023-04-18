package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags Menus
// API Menus : passing specific parameter to DBaaS from Service
// @Security BearerAuth
// @Param filter body engine.Menu true "Pass session information to DBaaS Parameter"
// @Accept  json
// @Produce  json
// @Router /api/v1/menus [post]
func CreateMenus(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Menus
// API Menus : passing specific parameter to DBaaS from Service
// @Security BearerAuth
// @Param id path string true "Pass session information to DBaaS Parameter"
// @Accept  json
// @Produce  json
// @Router /api/v1/menus/{id} [get]
func ShowMenus(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Menus
// API Menus : passing specific parameter to DBaaS from Service
// @Security BearerAuth
// @Param id path string true "Pass session information to DBaaS Parameter"
// @Param filter body engine.Menu true "Pass session information to DBaaS Parameter"
// @Accept  json
// @Produce  json
// @Router /api/v1/menus/{id} [put]
func UpdateMenus(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Menus
// API Menus : passing specific parameter to DBaaS from Service
// @Security BearerAuth
// @Param id path string true "Pass session information to DBaaS Parameter"
// @Accept  json
// @Produce  json
// @Router /api/v1/menus/{id} [delete]
func DeleteMenus(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
