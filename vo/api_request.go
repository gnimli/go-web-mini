package vo

// 获取接口列表结构体
type ApiListRequest struct {
	Method   string `json:"method" form:"method"`
	Path     string `json:"path" form:"path"`
	Category string `json:"category" form:"category"`
	Creator  string `json:"creator" form:"creator"`
	PageNum  uint   `json:"pageNum" form:"pageNum"`
	PageSize uint   `json:"pageSize" form:"pageSize"`
}

// 创建接口结构体
type CreateApiRequest struct {
	Method   string `json:"method" form:"method" validate:"required"`
	Path     string `json:"path" form:"path" validate:"required"`
	Category string `json:"category" form:"category" validate:"required"`
	Desc     string `json:"desc" form:"desc"`
}

// 更新接口结构体
type UpdateApiRequest struct {
	Method   string `json:"method" form:"method"`
	Path     string `json:"path" form:"path"`
	Category string `json:"category" form:"category"`
	Desc     string `json:"desc" form:"desc"`
}

// 批量删除接口结构体
type DeleteApiRequest struct {
	ApiIds []uint `json:"apiIds" form:"apiIds"`
}
