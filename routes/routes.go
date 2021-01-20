package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-lim/common"
	"go-lim/config"
	"go-lim/middleware"
	"time"
)

// 初始化
func InitRoutes() *gin.Engine {
	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	r := gin.Default()
	// 创建不带中间件的路由:
	//r := gin.New()
	//r.Use(gin.Recovery())

	//启用限流中间件,默认每50毫秒填充一个令牌，最多填充200个
	fillInterval := time.Duration(config.Conf.RateLimit.FillInterval)
	capacity := config.Conf.RateLimit.Capacity
	r.Use(middleware.RateLimitMiddleware(time.Millisecond*fillInterval, capacity))
	//启用全局跨域中间件
	r.Use(middleware.CORSMiddleware())

	authMiddleware, err := middleware.InitAuth()
	if err != nil {
		common.Log.Panicf("初始化JWT中间件失败：%v", err)
		panic(fmt.Sprintf("初始化JWT中间件失败：%v", err))
	}

	// 路由分组
	apiGroup := r.Group("/" + config.Conf.System.UrlPathPrefix)

	// 注册路由
	InitBaseRoutes(apiGroup, authMiddleware) // 注册公共路由,无需jwt中间件
	InitUserRoutes(apiGroup, authMiddleware) // 注册用户路由
	InitRoleRoutes(apiGroup, authMiddleware) // 注册角色路由
	InitMenuRoutes(apiGroup, authMiddleware) // 注册菜单路由
	InitApiRoutes(apiGroup, authMiddleware)  // 注册接口路由

	common.Log.Info("初始化路由完成！")
	return r
}
