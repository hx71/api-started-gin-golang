package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/app/user"
	"github.com/hx71/api-started-gin-golang/config"
	"github.com/hx71/api-started-gin-golang/helpers"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
)

type UserHandler interface {
	Index(ctx *gin.Context)
	Create(ctx *gin.Context)
	Show(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type userHandler struct {
	Usecase user.Usecase
}

func NewUserHandler(usecase user.Usecase) UserHandler {
	return &userHandler{
		Usecase: usecase,
	}
}

func (u *userHandler) Index(ctx *gin.Context) {
	pagination := helpers.GeneratePaginationRequest(ctx)
	res := u.Usecase.Pagination(ctx, pagination)
	if !res.Status {
		response := response.ResponseError("failed to get data user", res.Message)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.ResponseSuccess("list of user", res.Data)
	ctx.JSON(http.StatusOK, response)
}

func (u *userHandler) Create(ctx *gin.Context) {
	var req dto.UserCreateValidation
	err := ctx.ShouldBind(&req)
	if err != nil {
		response := response.ResponseError(config.MessageErr.FailedProcess, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// if !u.Usecase.FindByEmail(req.Email) {
	// 	response := response.ResponseError(config.MessageErr.FailedProcess, "duplicate email")
	// 	ctx.JSON(http.StatusConflict, response)
	// } else {
	err = u.Usecase.Create(req)
	if err != nil {
		go helpers.CreateLogError(uuid.NewString(), helpers.GetIP(ctx), "users", "created users", err.Error())
		response := response.ResponseError("failed to process created", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	go helpers.CreateLogInfo(uuid.NewString(), helpers.GetIP(ctx), "users", "created users", "created success")
	response := response.ResultSuccess("created success")
	ctx.JSON(http.StatusCreated, response)
	// }
}

func (u *userHandler) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User = u.Usecase.Show(id)
	if (user == models.User{}) {
		res := response.ResponseError("Data not found", "No data with given id")
		ctx.JSON(http.StatusNotFound, res)
	} else {
		response := response.ResponseSuccess("detail user", user)
		ctx.JSON(http.StatusOK, response)
	}
}

func (u *userHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User = u.Usecase.Show(id)
	if user.ID == "" {
		res := response.ResponseError("data not found", "no data with given id")
		ctx.JSON(http.StatusNotFound, res)
	} else {
		var userValidation dto.UserUpdateValidation
		userValidation.ID = id
		err := ctx.ShouldBind(&userValidation)
		if err != nil {
			response := response.ResponseError(config.MessageErr.FailedProcess, err.Error())
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		err = u.Usecase.Update(userValidation)
		if err != nil {
			go helpers.CreateLogError(uuid.NewString(), helpers.GetIP(ctx), "users", "updated users", err.Error())
			response := response.ResponseError("update failed", err.Error())
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		go helpers.CreateLogInfo(uuid.NewString(), helpers.GetIP(ctx), "users", "updated users", "updated success")
		response := response.ResponseSuccess("update success", nil)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (u *userHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User = u.Usecase.Show(id)
	if user.ID == "" {
		response := response.ResponseError("data not found", "no data with given id")
		ctx.JSON(http.StatusNotFound, response)
	} else {
		err := u.Usecase.Delete(user)
		if err != nil {
			go helpers.CreateLogError(uuid.NewString(), helpers.GetIP(ctx), "users", "deleted users", err.Error())
			response := response.ResponseError("failed to process deleted", err.Error())
			ctx.JSON(http.StatusNotFound, response)
			return
		}
		go helpers.CreateLogInfo(uuid.NewString(), helpers.GetIP(ctx), "users", "deleted users", "deleted success")
		response := response.ResultSuccess("deleted success")
		ctx.JSON(http.StatusOK, response)
	}
}
