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
		response := response.ResponseError(config.MessageErr.FailedProcess, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = s.roleService.Create(req)
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

func (s *roleController) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	var role models.Roles = s.roleService.Show(id)
	if role.ID == "" {
		res := response.ResponseError("Data not found", "No data with given id")
		ctx.JSON(http.StatusNotFound, res)
	} else {
		response := response.ResponseSuccess("detail role", role)
		ctx.JSON(http.StatusOK, response)
	}
}

func (s *roleController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var role models.Roles = s.roleService.Show(id)
	if role.ID == "" {
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

		err = s.roleService.Update(req)
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

func (s *roleController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var role models.Roles = s.roleService.Show(id)
	if role.ID == "" {
		response := response.ResponseError("data not found", "no data with given id")
		ctx.JSON(http.StatusNotFound, response)
	} else {
		err := s.roleService.Delete(role)
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
