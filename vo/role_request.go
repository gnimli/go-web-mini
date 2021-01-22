package vo

// 新增角色结构体
type CreateRoleRequest struct {
	Name    string `json:"name" form:"name" validate:"required"`
	Keyword string `json:"keyword" form:"keyword" validate:"required"`
	Desc    string `json:"desc" form:"desc"`
	Status  uint   `json:"status" form:"status"`
	Sort    uint   `json:"sort" form:"sort" validate:"required"`
	Creator string `json:"creator" form:"creator"`
}

// 获取用户角色结构体
type RoleListRequest struct {
	Name     string `json:"name" form:"name"`
	Keyword  string `json:"keyword" form:"keyword"`
	Status   uint   `json:"status" form:"status"`
	PageNum  uint   `json:"pageNum" form:"pageNum"`
	PageSize uint   `json:"pageSize" form:"pageSize"`
}

// 更新角色的菜单权限
type UpdateRoleMenusRequest struct {
	MenuIds []int `json:"menuIds"`
}

// 更新角色的API权限
type UpdateRoleApisRequest struct {
	ApiIds []int `json:"apiIds"`
}
