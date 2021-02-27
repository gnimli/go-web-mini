package vo

// 操作日志请求结构体
type OperationLogListRequest struct {
	Username string `json:"username" form:"username"`
	Ip       string `json:"ip" form:"ip"`
	Path     string `json:"path" form:"path"`
	Status   int    `json:"status" form:"status"`
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

// 批量删除操作日志结构体
type DeleteOperationLogRequest struct {
	OperationLogIds []uint `json:"operationLogIds" form:"operationLogIds"`
}
