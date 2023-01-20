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
// @Router /api/v1/version [get]
func ApiVersion(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Todo
// API Todo : passing specific parameter to DBaaS from Service Portal
// @Accept  json
// @Produce  json
// @Success 200 {object} engine.ResponseSuccess
// @Failure 400 {object} engine.ResponseStatus
// / @Failure 404 {object} engine.ResponseStatus
// / @Failure 500 {object} engine.ResponseStatus
// @Router /api/v1/todo [get]
func IndexTodo(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Todo
// API Todo : passing specific parameter to DBaaS from Service Portal
// @Accept  json
// @Produce  json
// @Success 200 {object} engine.ResponseSuccess
// @Failure 400 {object} engine.ResponseStatus
// / @Failure 404 {object} engine.ResponseStatus
// / @Failure 500 {object} engine.ResponseStatus
// @Router /api/v1/todo [post]
func CreateTodo(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Todo
// API Todo : passing specific parameter to DBaaS from Service Portal
// @Accept  json
// @Produce  json
// @Success 200 {object} engine.ResponseSuccess
// @Failure 400 {object} engine.ResponseStatus
// / @Failure 404 {object} engine.ResponseStatus
// / @Failure 500 {object} engine.ResponseStatus
// @Router /api/v1/todo/{id} [get]
func ShowTodo(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Todo
// API Todo : passing specific parameter to DBaaS from Service Portal
// @Accept  json
// @Produce  json
// @Success 200 {object} engine.ResponseSuccess
// @Failure 400 {object} engine.ResponseStatus
// / @Failure 404 {object} engine.ResponseStatus
// / @Failure 500 {object} engine.ResponseStatus
// @Router /api/v1/todo/{id} [put]
func UpdateTodo(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

// @Tags Todo
// API Todo : passing specific parameter to DBaaS from Service Portal
// @Accept  json
// @Produce  json
// @Success 200 {object} engine.ResponseSuccess
// @Failure 400 {object} engine.ResponseStatus
// / @Failure 404 {object} engine.ResponseStatus
// / @Failure 500 {object} engine.ResponseStatus
// @Router /api/v1/todo/{id} [delete]
func DeleteTodo(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
