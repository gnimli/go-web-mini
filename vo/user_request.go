package vo

// 创建用户结构体
type CreateUserRequest struct {
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password"`
	Mobile       string `json:"mobile" binding:"required"`
	Avatar       string `json:"avatar"`
	Nickname     string `json:"nickname"`
	Introduction string `json:"introduction"`
	Status       uint   `json:"status"`
	Creator      string `json:"creator"`
}

// 获取用户列表结构体
type UserListRequest struct {
	Username string `json:"username" form:"username"`
	Mobile   string `json:"mobile" form:"mobile"`
	Nickname string `json:"nickname" form:"nickname"`
	Status   *uint  `json:"status" form:"status"`
	PageNum  uint   `json:"pageNum" form:"pageNum"`
	PageSize uint   `json:"pageSize" form:"pageSize"`
}

// 修改密码结构体
type ChangePwdRequest struct {
	OldPassword string `json:"oldPassword" form:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" form:"newPassword" binding:"required"`
}
