package routes

import (
	"github.com/gin-gonic/gin"
	"go-lim/controller"
	"go-lim/middleware"
)

func InitApiRoutes(r *gin.RouterGroup) gin.IRoutes {
	apiController := controller.NewApiController()
	router := r.Group("/api")
	router.Use(middleware.AuthMiddleware())
	{
		router.GET("/list", apiController.GetApis)
		router.GET("/all/category/:roleId", apiController.GetAllApiGroupByCategoryByRoleId)
		router.POST("/create", apiController.CreateApi)
		router.PATCH("/update/:apiId", apiController.UpdateApiById)
		router.DELETE("/delete/batch", apiController.BatchDeleteApiByIds)
	}

	return r
}