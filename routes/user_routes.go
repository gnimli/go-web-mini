package routes

import (
	"github.com/gin-gonic/gin"
	"go-lim/controller"
	"go-lim/middleware"
)

// 注册用户路由
func InitUserRoutes(r *gin.RouterGroup) gin.IRoutes {
	userController := controller.NewUserController()
	router := r.Group("/user")
	// 开启jwt认证中间件
	router.Use(middleware.AuthMiddleware())
	{
		router.POST("/info", userController.GetUserInfo)
		router.GET("/list", userController.GetUsers)
		router.PUT("/changePwd", userController.ChangePwd)
		router.POST("/create", userController.CreateUser)
		router.PATCH("/update/:userId", userController.UpdateUserById)
		router.DELETE("/delete/batch", userController.BatchDeleteUserByIds)
	}
	return r
}
