package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 注册基础路由
func InitBaseRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	router := r.Group("/base")
	{
		// 登录登出刷新token无需鉴权
		router.POST("/login", authMiddleware.LoginHandler)
		router.POST("/logout", authMiddleware.LogoutHandler)
		router.POST("/refreshToken", authMiddleware.RefreshHandler)
	}
	return r
}
