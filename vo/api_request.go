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
	//Title    string `json:"title"`
	Creator string `json:"creator" form:"creator"`
	RoleIds []uint `json:"roleIds" form:"roleIds"` // 绑定可以访问该接口的角色
}
