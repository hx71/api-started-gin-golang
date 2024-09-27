package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hx71/api-started-gin-golang/app/auditlog"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/config"
	"github.com/hx71/api-started-gin-golang/helpers"
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

func (h *auditLogHandler) Index(ctx *gin.Context) {
	pagination := response.GeneratePaginationRequest(ctx)
	data := h.Usecase.Pagination(ctx, pagination)

	if !data.Status {
		errResponse := response.ResponseError("failed to get data audit-log", data.Message)
		ctx.JSON(http.StatusBadRequest, errResponse)
		return
	}

	ctx.JSON(http.StatusOK, response.ResponseSuccess("list of audit-log", data.Data))
}

func (s *auditLogHandler) Create(ctx *gin.Context) {
	req := dto.AuditLogCreateValidation{
		ID: uuid.NewString(),
	}

	if err := ctx.ShouldBind(&req); err != nil {
		response := response.ResponseError(config.MessageErr.FailedProcess, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if err := s.Usecase.Create(req); err != nil {
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
	role := s.Usecase.Show(id)

	if role.ID == "" {
		res := response.ResponseError("Data not found", "No data with given id")
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	ctx.JSON(http.StatusOK, response.ResponseSuccess("detail audit-log", role))
}

func (s *auditLogHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	auditlog := s.Usecase.Show(id)

	if auditlog.ID == "" {
		response := response.ResponseError("data not found", "no data with given id")
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	err := s.Usecase.Delete(auditlog)
	if err != nil {
		go helpers.CreateLogError(uuid.NewString(), helpers.GetIP(ctx), "audit-log", "deleted audit-log", err.Error())

		response := response.ResponseError("failed to process deleted", err.Error())
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	go helpers.CreateLogInfo(uuid.NewString(), helpers.GetIP(ctx), "audit-log", "deleted audit-log", "deleted success")

	ctx.JSON(http.StatusOK, response.ResultSuccess("deleted success"))
}
