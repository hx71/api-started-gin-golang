package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VersionHandler interface {
	Version(ctx *gin.Context)
}

type versionHandler struct {
}

func NewVersionHandler() VersionHandler {
	return &versionHandler{}
}

func (h *versionHandler) Version(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"version": "1.0.0",
	})
}
