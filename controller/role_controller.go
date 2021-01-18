package controller

import (
	"github.com/gin-gonic/gin"
	"go-lim/repository"
)

type IRoleController interface {
	GetRoles(c *gin.Context)
	CreateRole(c *gin.Context)
	UpdateRoleById(c *gin.Context)
	UpdateRoleMenusById(c *gin.Context)
	UpdateRoleApisById(c *gin.Context)
	BatchDeleteRoleByIds(c *gin.Context)
}

type RoleController struct {
	RoleRepository repository.IRoleRepository
}

func NewRoleController() IRoleController {
	roleRepository := repository.NewRoleRepository()
	roleController := RoleController{RoleRepository: roleRepository}
	return roleController
}

func (r RoleController) GetRoles(c *gin.Context) {
	panic("implement me")
}

func (r RoleController) CreateRole(c *gin.Context) {
	panic("implement me")
}

func (r RoleController) UpdateRoleById(c *gin.Context) {
	panic("implement me")
}

func (r RoleController) UpdateRoleMenusById(c *gin.Context) {
	panic("implement me")
}

func (r RoleController) UpdateRoleApisById(c *gin.Context) {
	panic("implement me")
}

func (r RoleController) BatchDeleteRoleByIds(c *gin.Context) {
	panic("implement me")
}
