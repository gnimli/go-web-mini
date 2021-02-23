package vo

// 新增角色结构体
type CreateRoleRequest struct {
	Name    string `json:"name" form:"name" validate:"required,min=1,max=20"`
	Keyword string `json:"keyword" form:"keyword" validate:"required,min=1,max=20"`
	Desc    string `json:"desc" form:"desc" validate:"min=0,max=100"`
	Status  uint   `json:"status" form:"status" validate:"oneof=1 2"`
	Sort    uint   `json:"sort" form:"sort" validate:"gte=1,lte=999"`
}

// 获取用户角色结构体
type RoleListRequest struct {
	Name     string `json:"name" form:"name"`
	Keyword  string `json:"keyword" form:"keyword"`
	Status   uint   `json:"status" form:"status"`
	PageNum  uint   `json:"pageNum" form:"pageNum"`
	PageSize uint   `json:"pageSize" form:"pageSize"`
}

// 批量删除角色结构体
type DeleteRoleRequest struct {
	RoleIds []uint `json:"roleIds" form:"roleIds"`
}

// 更新角色的权限菜单
type UpdateRoleMenusRequest struct {
	MenuIds []uint `json:"menuIds" form:"menuIds"`
}

// 更新角色的权限接口
type UpdateRoleApisRequest struct {
	ApiIds []uint `json:"apiIds" form:"apiIds"`
}
