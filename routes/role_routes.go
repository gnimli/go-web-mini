package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go-web-mini/controller"
	"go-web-mini/middleware"
)

func InitRoleRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	roleController := controller.NewRoleController()
	router := r.Group("/role")
	// 开启jwt认证中间件
	router.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	router.Use(middleware.CasbinMiddleware())
	{
		router.GET("/list", roleController.GetRoles)
		router.POST("/create", roleController.CreateRole)
		router.PATCH("/update/:roleId", roleController.UpdateRoleById)
		router.GET("/menus/get/:roleId", roleController.GetRoleMenusById)
		router.PATCH("/menus/update/:roleId", roleController.UpdateRoleMenusById)
		router.GET("/apis/get/:roleId", roleController.GetRoleApisById)
		router.PATCH("/apis/update/:roleId", roleController.UpdateRoleApisById)
		router.DELETE("/delete/batch", roleController.BatchDeleteRoleByIds)
	}
	return r
}
