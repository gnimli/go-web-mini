package routes

import (
	"github.com/gin-gonic/gin"
	"go-lim/controller"
)

// 注册基础路由
func InitBaseRoutes(r *gin.RouterGroup) gin.IRoutes {
	baseController := controller.NewBaseController()
	router := r.Group("/base")
	{
		// 登录登出/刷新token无需鉴权
		router.POST("/login", baseController.Login)
		router.POST("/logout", baseController.Logout)
		router.POST("/refreshToken", baseController.RefreshToken)
	}
	return r
}
