package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hx71/api-started-gin-golang/app/auditlog"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/config"
	"github.com/hx71/api-started-gin-golang/helpers"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
)

type AuditLogHandler interface {
	Index(ctx *gin.Context)
	Create(ctx *gin.Context)
	Show(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type auditLogHandler struct {
	Usecase auditlog.Usecase
}

func NewAuditLogHandler(usecase auditlog.Usecase) AuditLogHandler {
	return &auditLogHandler{
		Usecase: usecase,
	}
}

func (u *auditLogHandler) Index(ctx *gin.Context) {
	pagination := response.GeneratePaginationRequest(ctx)
	res := u.Usecase.Pagination(ctx, pagination)
	if !res.Status {
		response := response.ResponseError("failed to get data audit-log", res.Message)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.ResponseSuccess("list of audit-log", res.Data)
	ctx.JSON(http.StatusOK, response)
}

func (s *auditLogHandler) Create(ctx *gin.Context) {
	var req dto.AuditLogCreateValidation
	req.ID = uuid.NewString()
	err := ctx.ShouldBind(&req)
	if err != nil {
		response := response.ResponseError(config.MessageErr.FailedProcess, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = s.Usecase.Create(req)
	if err != nil {
		go helpers.CreateLogError(uuid.NewString(), helpers.GetIP(ctx), "audit-log", "created audit-log", err.Error())
		response := response.ResponseError("failed to process created", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := response.ResultSuccess("created success")
	ctx.JSON(http.StatusCreated, response)
}

func (s *auditLogHandler) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	var role models.AuditLog = s.Usecase.Show(id)
	if role.ID == "" {
		res := response.ResponseError("Data not found", "No data with given id")
		ctx.JSON(http.StatusNotFound, res)
	} else {
		response := response.ResponseSuccess("detail audit-log", role)
		ctx.JSON(http.StatusOK, response)
	}
}

func (s *auditLogHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var auditlog models.AuditLog = s.Usecase.Show(id)
	if auditlog.ID == "" {
		response := response.ResponseError("data not found", "no data with given id")
		ctx.JSON(http.StatusNotFound, response)
	} else {
		err := s.Usecase.Delete(auditlog)
		if err != nil {
			go helpers.CreateLogError(uuid.NewString(), helpers.GetIP(ctx), "audit-log", "deleted audit-log", err.Error())
			response := response.ResponseError("failed to process deleted", err.Error())
			ctx.JSON(http.StatusNotFound, response)
			return
		}
		go helpers.CreateLogInfo(uuid.NewString(), helpers.GetIP(ctx), "audit-log", "deleted audit-log", "deleted success")
		response := response.ResultSuccess("deleted success")
		ctx.JSON(http.StatusOK, response)
	}

}
