package middleware

import (
	"github.com/gin-gonic/gin"
	"go-lim/config"
	"go-lim/model"
	"go-lim/repository"
	"strings"
	"time"
)

// 操作日志channel
var OperationLogChan = make(chan *model.OperationLog, 50)

func OperationLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行耗时
		timeCost := endTime.Sub(startTime).Milliseconds()

		// 获取当前登录用户
		var username string
		ctxUser, exists := c.Get("user")
		if !exists {
			username = "未登录"
		}
		user, ok := ctxUser.(model.User)
		if !ok {
			username = "未登录"
		}
		username = user.Username

		// 获取访问路径
		path := strings.TrimPrefix(c.Request.URL.Path, "/"+config.Conf.System.UrlPathPrefix)

		// 获取接口描述
		apiRepository := repository.NewApiRepository()
		apiDesc, _ := apiRepository.GetApiDescByPath(path)

		operationLog := model.OperationLog{
			Username:   username,
			Ip:         c.ClientIP(),
			IpLocation: "",
			Method:     c.Request.Method,
			Path:       path,
			Desc:       apiDesc,
			Status:     0,
			StartTime:  startTime,
			TimeCost:   timeCost,
			//UserAgent:  c.Request.UserAgent(),
		}

		// 最好是将日志发送到rabbitmq或者kafka中
		// 这里是发送到channel中，开启5个goroutine处理
		OperationLogChan <- &operationLog
	}
}
