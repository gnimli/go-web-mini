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
	Method   string `json:"method" binding:"required"`
	Path     string `json:"path" binding:"required"`
	Category string `json:"category" binding:"required"`
	Desc     string `json:"desc"`
	//Title    string `json:"title"`
	Creator  string `json:"creator"`
	RoleIds  []uint `json:"roleIds"` // 绑定可以访问该接口的角色
}

