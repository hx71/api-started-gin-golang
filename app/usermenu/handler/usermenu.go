package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/app/usermenu"
	"github.com/hx71/api-started-gin-golang/config"
	"github.com/hx71/api-started-gin-golang/helpers"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
)

type UserMenuHandler struct {
	Usecase usermenu.Usecase
}

func (u *UserMenuHandler) Index(ctx *gin.Context) {
	pagination := helpers.GeneratePaginationRequest(ctx)
	res := u.Usecase.Pagination(ctx, pagination)
	if !res.Status {
		response := response.ResponseError("failed to get data user menu", res.Message)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.ResponseSuccess("list of user menu", res.Data)
	ctx.JSON(http.StatusOK, response)
}

func (u *UserMenuHandler) Create(ctx *gin.Context) {
	var req []dto.UserMenuCreateValidation
	// req.ID = uuid.NewString()
	err := ctx.ShouldBind(&req)
	if err != nil {
		response := response.ResponseError(config.MessageErr.FailedProcess, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = u.Usecase.Create(req)
	if err != nil {
		go helpers.CreateLogError(uuid.NewString(), helpers.GetIP(ctx), "user menu", "created user menu", err.Error())
		response := response.ResponseError("failed to process created", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	go helpers.CreateLogInfo(uuid.NewString(), helpers.GetIP(ctx), "user menu", "created user menu", "created success")
	response := response.ResultSuccess("created success")
	ctx.JSON(http.StatusCreated, response)
}

func (u *UserMenuHandler) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	var userMenu models.UserMenus = u.Usecase.Show(id)
	if userMenu.ID == "" {
		res := response.ResponseError("Data not found", "No data with given id")
		ctx.JSON(http.StatusNotFound, res)
	} else {
		response := response.ResponseSuccess("detail user menu", userMenu)
		ctx.JSON(http.StatusOK, response)
	}
}

func (u *UserMenuHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var userMenu models.UserMenus = u.Usecase.Show(id)
	if userMenu.ID == "" {
		res := response.ResponseError("data not found", "no data with given id")
		ctx.JSON(http.StatusNotFound, res)
	} else {
		var req dto.UserMenuCreateValidation
		err := ctx.ShouldBind(&req)
		if err != nil {
			response := response.ResponseError(config.MessageErr.FailedProcess, err.Error())
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		err = u.Usecase.Update(id, req)
		if err != nil {
			go helpers.CreateLogError(uuid.NewString(), helpers.GetIP(ctx), "user menu", "updated user menu", err.Error())
			response := response.ResponseError("failed to process updated", err.Error())
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		go helpers.CreateLogInfo(uuid.NewString(), helpers.GetIP(ctx), "user menu", "updated user menu", "updated success")
		response := response.ResultSuccess("updated success")
		ctx.JSON(http.StatusCreated, response)
	}
}

func (u *UserMenuHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var userMenu models.UserMenus = u.Usecase.Show(id)
	if userMenu.ID == "" {
		response := response.ResponseError("data not found", "no data with given id")
		ctx.JSON(http.StatusNotFound, response)
	} else {
		err := u.Usecase.Delete(userMenu)
		if err != nil {
			go helpers.CreateLogError(uuid.NewString(), helpers.GetIP(ctx), "user menu", "deleted user menu", err.Error())
			response := response.ResponseError("failed to process deleted", err.Error())
			ctx.JSON(http.StatusNotFound, response)
			return
		}
		go helpers.CreateLogInfo(uuid.NewString(), helpers.GetIP(ctx), "user menu", "deleted user menu", "deleted success")
		response := response.ResultSuccess("deleted success")
		ctx.JSON(http.StatusOK, response)
	}

}
