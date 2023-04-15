package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/dto"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/service"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/helpers"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/models"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/response"
)

// MenuController is a contract what this controller can do
type MenuController interface {
	Index(ctx *gin.Context)
	Create(ctx *gin.Context)
	Show(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type menuController struct {
	menuService service.MenuService
	jwtService  service.JWTService
}

// NewMenuController create a new instances of MenuController
func NewMenuController(menuServ service.MenuService, jwtServ service.JWTService) MenuController {
	return &menuController{
		menuService: menuServ,
		jwtService:  jwtServ,
	}
}

func (s *menuController) Index(ctx *gin.Context) {
	pagination := helpers.GeneratePaginationRequest(ctx)
	res := s.menuService.Pagination(ctx, pagination)
	if !res.Status {
		response := response.ResponseError("failed to get data role", res.Message)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.ResponseSuccess("list of role", res.Data)
	ctx.JSON(http.StatusOK, response)
}

func (s *menuController) Create(ctx *gin.Context) {
	var req dto.MenuCreateValidation
	req.ID = uuid.NewString()
	err := ctx.ShouldBind(&req)
	if err != nil {
		response := response.ResponseError("failed to process request", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = s.menuService.Create(req)
	if err != nil {
		response := response.ResponseError("failed to process created", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.ResultSuccess("created success")
	ctx.JSON(http.StatusCreated, response)
}

func (s *menuController) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	var role models.Menus = s.menuService.Show(id)
	if role.ID == "" {
		res := response.ResponseError("Data not found", "No data with given id")
		ctx.JSON(http.StatusNotFound, res)
	} else {
		response := response.ResponseSuccess("detail user", role)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *menuController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var role models.Menus = c.menuService.Show(id)
	if role.ID == "" {
		res := response.ResponseError("data not found", "no data with given id")
		ctx.JSON(http.StatusNotFound, res)
	} else {
		var req dto.MenuCreateValidation
		req.ID = id
		err := ctx.ShouldBind(&req)
		if err != nil {
			response := response.ResponseError("failed to process request", err.Error())
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		err = c.menuService.Update(req)
		if err != nil {
			response := response.ResponseError("failed to process deleted", err.Error())
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		response := response.ResultSuccess("updated success")
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *menuController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response := response.ResponseError("data not found", "no data with given id")
		ctx.JSON(http.StatusNotFound, response)
	} else {
		err := c.menuService.Delete(id)
		if err != nil {
			response := response.ResponseError("failed to process deleted", err.Error())
			ctx.JSON(http.StatusNotFound, response)
			return
		}
		response := response.ResultSuccess("deleted success")
		ctx.JSON(http.StatusOK, response)
	}
}
