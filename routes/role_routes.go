package routes

import (
	"github.com/gin-gonic/gin"
	"go-lim/controller"
	"go-lim/middleware"
)

func InitRoleRoutes(r *gin.RouterGroup) gin.IRoutes {
	roleController := controller.NewRoleController()
	router := r.Group("/role")
	// 开启jwt认证中间件
	router.Use(middleware.AuthMiddleware())
	{
		router.GET("/list", roleController.GetRoles)
		router.POST("/create", roleController.CreateRole)
		router.PATCH("/update/:roleId", roleController.UpdateRoleById)
		router.PATCH("/menus/update/:roleId", roleController.UpdateRoleMenusById)
		router.PATCH("/apis/update/:roleId", roleController.UpdateRoleApisById)
		router.DELETE("/delete/batch", roleController.BatchDeleteRoleByIds)
	}
	return r
}
