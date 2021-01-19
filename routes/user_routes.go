package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go-lim/controller"
)

// 注册用户路由
func InitUserRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	userController := controller.NewUserController()
	router := r.Group("/user")
	// 开启jwt认证中间件
	router.Use(authMiddleware.MiddlewareFunc())
	//
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
