package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-mini/repository"
)

type IMenuController interface {
	GetMenuTree(c *gin.Context)
	GetAllMenuByRoleId(c *gin.Context)
	GetMenus(c *gin.Context)
	CreateMenu(c *gin.Context)
	UpdateMenuById(c *gin.Context)
	BatchDeleteMenuByIds(c *gin.Context)
}

type MenuController struct {
	MenuRepository repository.IMenuRepository
}

func NewMenuController() IMenuController {
	menuRepository := repository.NewMenuRepository()
	menuController := MenuController{MenuRepository: menuRepository}
	return menuController
}

func (m MenuController) GetMenuTree(c *gin.Context) {
	panic("implement me")
}

func (m MenuController) GetAllMenuByRoleId(c *gin.Context) {
	panic("implement me")
}

func (m MenuController) GetMenus(c *gin.Context) {
	panic("implement me")
}

func (m MenuController) CreateMenu(c *gin.Context) {
	panic("implement me")
}

func (m MenuController) UpdateMenuById(c *gin.Context) {
	panic("implement me")
}

func (m MenuController) BatchDeleteMenuByIds(c *gin.Context) {
	panic("implement me")
}
