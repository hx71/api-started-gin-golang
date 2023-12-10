package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/app/role"
	"github.com/hx71/api-started-gin-golang/config"
	"github.com/hx71/api-started-gin-golang/helpers"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
)

type RoleHandler interface {
	Index(ctx *gin.Context)
	Create(ctx *gin.Context)
	Show(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type roleHandler struct {
	Usecase role.Usecase
}

func NewRoleHandler(usecase role.Usecase) RoleHandler {
	return &roleHandler{
		Usecase: usecase,
	}
}

func (u *roleHandler) Index(ctx *gin.Context) {
	pagination := response.GeneratePaginationRequest(ctx)
	res := u.Usecase.Pagination(ctx, pagination)
	if !res.Status {
		response := response.ResponseError("failed to get data roles", res.Message)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.ResponseSuccess("list of roles", res.Data)
	ctx.JSON(http.StatusOK, response)
}

func (u *roleHandler) Create(ctx *gin.Context) {
	var req dto.RoleCreateValidation
	req.ID = uuid.NewString()
	err := ctx.ShouldBind(&req)
	if err != nil {
		response := response.ResponseError(config.MessageErr.FailedProcess, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = u.Usecase.Create(req)
	if err != nil {
		go helpers.CreateLogError(uuid.NewString(), helpers.GetIP(ctx), "roles", "created roles", err.Error())
		response := response.ResponseError("failed to process created", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	go helpers.CreateLogInfo(uuid.NewString(), helpers.GetIP(ctx), "roles", "created roles", "created success")
	response := response.ResultSuccess("created success")
	ctx.JSON(http.StatusCreated, response)
}

func (u *roleHandler) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	var req models.Role = u.Usecase.Show(id)
	if req.ID == "" {
		res := response.ResponseError("Data not found", "No data with given id")
		ctx.JSON(http.StatusNotFound, res)
	} else {
		response := response.ResponseSuccess("detail role", req)
		ctx.JSON(http.StatusOK, response)
	}
}

func (u *roleHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var req models.Role = u.Usecase.Show(id)
	if req.ID == "" {
		res := response.ResponseError("data not found", "no data with given id")
		ctx.JSON(http.StatusNotFound, res)
	} else {
		var req dto.RoleCreateValidation
		req.ID = id
		err := ctx.ShouldBind(&req)
		if err != nil {
			response := response.ResponseError(config.MessageErr.FailedProcess, err.Error())
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		err = u.Usecase.Update(req)
		if err != nil {
			go helpers.CreateLogError(uuid.NewString(), helpers.GetIP(ctx), "roles", "updated roles", err.Error())
			response := response.ResponseError("failed to process updated", err.Error())
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		go helpers.CreateLogInfo(uuid.NewString(), helpers.GetIP(ctx), "roles", "updated roles", "created success")
		response := response.ResultSuccess("updated success")
		ctx.JSON(http.StatusCreated, response)
	}
}

func (u *roleHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var req models.Role = u.Usecase.Show(id)
	if req.ID == "" {
		response := response.ResponseError("data not found", "no data with given id")
		ctx.JSON(http.StatusNotFound, response)
	} else {
		err := u.Usecase.Delete(req)
		if err != nil {
			go helpers.CreateLogError(uuid.NewString(), helpers.GetIP(ctx), "roles", "deleted roles", err.Error())
			response := response.ResponseError("failed to process deleted", err.Error())
			ctx.JSON(http.StatusNotFound, response)
			return
		}
		go helpers.CreateLogInfo(uuid.NewString(), helpers.GetIP(ctx), "roles", "deleted roles", "deleted success")
		response := response.ResultSuccess("deleted success")
		ctx.JSON(http.StatusOK, response)
	}
}
