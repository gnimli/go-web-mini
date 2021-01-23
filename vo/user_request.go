package vo

// 用户登录结构体
type RegisterAndLoginRequest struct {
	Username string `form:"username" json:"username" binding:"required,min=3,max=20"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=255"`
}

// 创建用户结构体
type CreateUserRequest struct {
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password"`
	Mobile       string `json:"mobile" validate:"required"`
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
	Status   uint   `json:"status" form:"status"`
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

// 修改密码结构体
type ChangePwdRequest struct {
	OldPassword string `json:"oldPassword" form:"oldPassword" validate:"required,min=6,max=255"`
	NewPassword string `json:"newPassword" form:"newPassword" validate:"required,min=6,max=255,nefield=OldPassword"`
}
