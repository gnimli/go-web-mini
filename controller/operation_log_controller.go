package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-web-mini/common"
	"go-web-mini/repository"
	"go-web-mini/response"
	"go-web-mini/vo"
)

type IOperationLogController interface {
	GetOperationLogs(c *gin.Context)             // 获取操作日志列表
	BatchDeleteOperationLogByIds(c *gin.Context) //批量删除操作日志
}

type OperationLogController struct {
	operationLogRepository repository.IOperationLogRepository
}

func NewOperationLogController() IOperationLogController {
	operationLogRepository := repository.NewOperationLogRepository()
	operationLogController := OperationLogController{operationLogRepository: operationLogRepository}
	return operationLogController
}

// 获取操作日志列表
func (oc OperationLogController) GetOperationLogs(c *gin.Context) {
	var req vo.OperationLogListRequest
	// 绑定参数
	if err := c.ShouldBind(&req); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := common.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(common.Trans)
		response.Fail(c, nil, errStr)
		return
	}
	// 获取
	logs, total, err := oc.operationLogRepository.GetOperationLogs(&req)
	if err != nil {
		response.Fail(c, nil, "获取操作日志列表失败: "+err.Error())
		return
	}
	response.Success(c, gin.H{"logs": logs, "total": total}, "获取操作日志列表成功")
}

// 批量删除操作日志
func (oc OperationLogController) BatchDeleteOperationLogByIds(c *gin.Context) {
	var req vo.DeleteOperationLogRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := common.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(common.Trans)
		response.Fail(c, nil, errStr)
		return
	}

	// 删除接口
	err := oc.operationLogRepository.BatchDeleteOperationLogByIds(req.OperationLogIds)
	if err != nil {
		response.Fail(c, nil, "删除日志失败: "+err.Error())
		return
	}

	response.Success(c, nil, "删除日志成功")
}
