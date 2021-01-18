package routes

import (
	"github.com/gin-gonic/gin"
	"go-lim/controller"
	"go-lim/middleware"
)

func InitMenuRoutes(r *gin.RouterGroup) gin.IRoutes {
	menuController := controller.NewMenuController()
	router := r.Group("/menu")
	router.Use(middleware.AuthMiddleware())
	{
		router.GET("/tree", menuController.GetMenuTree)
		router.GET("/all/:roleId", menuController.GetAllMenuByRoleId)
		router.GET("/list", menuController.GetMenus)
		router.POST("/create", menuController.CreateMenu)
		router.PATCH("/update/:menuId", menuController.UpdateMenuById)
		router.DELETE("/delete/batch", menuController.BatchDeleteMenuByIds)
	}

	return r
}
