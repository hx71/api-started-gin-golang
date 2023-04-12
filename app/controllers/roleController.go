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

// RoleController is a contract what this controller can do
type RoleController interface {
	Index(ctx *gin.Context)
	Create(ctx *gin.Context)
	Show(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type roleController struct {
	roleService service.RoleService
	jwtService  service.JWTService
}

// NewRoleController create a new instances of RoleController
func NewRoleController(roleServ service.RoleService, jwtServ service.JWTService) RoleController {
	return &roleController{
		roleService: roleServ,
		jwtService:  jwtServ,
	}
}

func (s *roleController) Index(ctx *gin.Context) {
	pagination := helpers.GeneratePaginationRequest(ctx)
	res := s.roleService.Pagination(ctx, pagination)
	if !res.Status {
		response := response.ResponseError("failed to get data role", res.Message)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.ResponseSuccess("list of role", res.Data)
	ctx.JSON(http.StatusOK, response)
}

func (s *roleController) Create(ctx *gin.Context) {
	var req dto.RoleCreateValidation
	req.ID = uuid.NewString()
	err := ctx.ShouldBind(&req)
	if err != nil {
		response := response.ResponseError("failed to process request", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user := s.roleService.Create(req)
	response := response.ResponseSuccess("created succeess", user)
	ctx.JSON(http.StatusCreated, response)
}

func (s *roleController) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	var role models.Roles = s.roleService.Show(id)
	if role.ID == "" {
		res := response.ResponseError("Data not found", "No data with given id")
		ctx.JSON(http.StatusNotFound, res)
	} else {
		response := response.ResponseSuccess("detail user", role)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *roleController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var role models.Roles = c.roleService.Show(id)
	if role.ID == "" {
		res := response.ResponseError("data not found", "no data with given id")
		ctx.JSON(http.StatusNotFound, res)
	} else {
		var req dto.RoleCreateValidation
		req.ID = id
		err := ctx.ShouldBind(&req)
		if err != nil {
			response := response.ResponseError("failed to process request", err.Error())
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		err = c.roleService.Update(req)
		if err != nil {
			response := response.ResponseError("update failed", err.Error())
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		response := response.ResponseSuccess("update success", nil)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *roleController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response := response.ResponseError("data not found", "no data with given id")
		ctx.JSON(http.StatusNotFound, response)
	} else {
		role := c.roleService.Delete(id)
		response := response.ResponseSuccess("deleted success", role)
		ctx.JSON(http.StatusOK, response)
	}
}
