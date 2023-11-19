package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/dto"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/service"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/config"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/helpers"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/models"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/response"
)

// UserMenuController is a contract what this controller can do
type UserMenuController interface {
	Index(ctx *gin.Context)
	Create(ctx *gin.Context)
	Show(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type userMenuController struct {
	userMenuService service.UserMenuService
	jwtService      service.JWTService
}

// NewUserMenuController create a new instances of UserMenuController
func NewUserMenuController(userMenuServ service.UserMenuService, jwtServ service.JWTService) UserMenuController {
	return &userMenuController{
		userMenuService: userMenuServ,
		jwtService:      jwtServ,
	}
}

func (s *userMenuController) Index(ctx *gin.Context) {
	pagination := helpers.GeneratePaginationRequest(ctx)
	res := s.userMenuService.Pagination(ctx, pagination)
	if !res.Status {
		response := response.ResponseError("failed to get data userMenu", res.Message)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.ResponseSuccess("list of userMenu", res.Data)
	ctx.JSON(http.StatusOK, response)
}

func (s *userMenuController) Create(ctx *gin.Context) {
	var req []dto.UserMenuCreateValidation
	// req.ID = uuid.NewString()
	err := ctx.ShouldBind(&req)
	if err != nil {
		response := response.ResponseError(config.MessageErr.FailedProcess, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = s.userMenuService.Create(req)
	if err != nil {
		go helpers.CreateLogError(uuid.NewString(), helpers.GetIP(ctx), "user menus", "created user menus", err.Error())
		response := response.ResponseError("failed to process created", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	go helpers.CreateLogInfo(uuid.NewString(), helpers.GetIP(ctx), "user menus", "created user menus", "created success")
	response := response.ResultSuccess("created success")
	ctx.JSON(http.StatusCreated, response)
}

func (s *userMenuController) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	var userMenu models.UserMenus = s.userMenuService.Show(id)
	if userMenu.ID == "" {
		res := response.ResponseError("Data not found", "No data with given id")
		ctx.JSON(http.StatusNotFound, res)
	} else {
		response := response.ResponseSuccess("detail userMenu", userMenu)
		ctx.JSON(http.StatusOK, response)
	}
}

func (s *userMenuController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var userMenu models.UserMenus = s.userMenuService.Show(id)
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
		err = s.userMenuService.Update(id, req)
		if err != nil {
			go helpers.CreateLogError(uuid.NewString(), helpers.GetIP(ctx), "user menus", "updated user menus", err.Error())
			response := response.ResponseError("failed to process updated", err.Error())
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		go helpers.CreateLogInfo(uuid.NewString(), helpers.GetIP(ctx), "user menus", "updated user menus", "updated success")
		response := response.ResultSuccess("updated success")
		ctx.JSON(http.StatusCreated, response)
	}
}

func (s *userMenuController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var userMenu models.UserMenus = s.userMenuService.Show(id)
	if userMenu.ID == "" {
		response := response.ResponseError("data not found", "no data with given id")
		ctx.JSON(http.StatusNotFound, response)
	} else {
		err := s.userMenuService.Delete(userMenu)
		if err != nil {
			go helpers.CreateLogError(uuid.NewString(), helpers.GetIP(ctx), "user menus", "deleted user menus", err.Error())
			response := response.ResponseError("failed to process deleted", err.Error())
			ctx.JSON(http.StatusNotFound, response)
			return
		}
		go helpers.CreateLogInfo(uuid.NewString(), helpers.GetIP(ctx), "user menus", "deleted user menus", "deleted success")
		response := response.ResultSuccess("deleted success")
		ctx.JSON(http.StatusOK, response)
	}

}
