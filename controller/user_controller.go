package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thoas/go-funk"
	"go-lim/common"
	"go-lim/dto"
	"go-lim/model"
	"go-lim/repository"
	"go-lim/response"
	"go-lim/util"
	"go-lim/vo"
)

type IUserController interface {
	GetUserInfo(c *gin.Context)    // 获取当前登录用户信息
	GetUsers(c *gin.Context)       // 获取用户列表
	ChangePwd(c *gin.Context)      // 修改密码
	CreateUser(c *gin.Context)     // 创建用户
	UpdateUserById(c *gin.Context) // 更新用户
	BatchDeleteUserByIds(c *gin.Context)
}

type UserController struct {
	UserRepository repository.IUserRepository
}

// 构造函数
func NewUserController() IUserController {
	userRepository := repository.NewUserRepository()
	userController := UserController{UserRepository: userRepository}
	return userController
}

// 获取当前登录用户信息
func (uc UserController) GetUserInfo(c *gin.Context) {
	user := uc.UserRepository.GetCurrentUser(c)
	userInfoDto := dto.ToUserInfoDto(user)
	response.Success(c, gin.H{
		"userInfo": userInfoDto,
	}, "获取当前用户信息成功")
}

// 获取用户列表
func (uc UserController) GetUsers(c *gin.Context) {
	var req vo.UserListRequest
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

	// 查询
	users, total, err := uc.UserRepository.GetUsers(&req)
	if err != nil {
		response.Fail(c, nil, "查询用户列表失败: "+err.Error())
		return
	}
	response.Success(c, gin.H{"users": dto.ToUsersDto(users), "total": total}, "查询用户列表成功")
}

// 修改密码
func (uc UserController) ChangePwd(c *gin.Context) {
	var req vo.ChangePwdRequest
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

	// 获取当前用户
	user := uc.UserRepository.GetCurrentUser(c)
	// 获取用户的真实正确密码
	correctPasswd := user.Password
	// 判断前端请求的密码是否等于真实密码
	err := util.ComparePasswd(correctPasswd, req.OldPassword)
	if err != nil {
		response.Fail(c, nil, "原密码有误")
		return
	}
	// 修改密码
	err = uc.UserRepository.ChangePwd(user.Username, util.GenPasswd(req.NewPassword))
	if err != nil {
		response.Fail(c, nil, "修改密码失败: "+err.Error())
		return
	}
	response.Success(c, nil, "修改密码成功")
}

// 创建用户
func (uc UserController) CreateUser(c *gin.Context) {
	var req vo.CreateUserRequest
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
	// 获取当前用户
	ctxUser := uc.UserRepository.GetCurrentUser(c)
	// 获取当前用户的所有角色
	currentRoles := ctxUser.Roles
	// 获取当前用户角色的排序，和前端传来的角色排序做比较
	var currentRoleSorts []int
	for _, role := range currentRoles {
		currentRoleSorts = append(currentRoleSorts, int(role.Sort))
	}
	// 当前用户角色排序最小值（最高等级角色）
	currentRoleSortMin := funk.MinInt(currentRoleSorts).(int)

	// 获取前端传来的用户角色id
	reqRoleIds := req.RoleIds
	// 根据角色id查询角色
	rr := repository.NewRoleRepository()
	roles, err := rr.GetRolesByIds(reqRoleIds)
	if err != nil {
		response.Fail(c, nil, "根据角色ID查询角色信息失败: "+err.Error())
		return
	}
	var reqRoleSorts []int
	for _, role := range roles {
		reqRoleSorts = append(reqRoleSorts, int(role.Sort))
	}
	// 前端传来用户角色排序最小值（最高等级角色）
	reqRoleSortMin := funk.MinInt(reqRoleSorts).(int)

	// 当前用户的角色排序最小值 需要小于 前端传来的角色排序最小值（用户不能创建比自己等级高的或者相同等级的用户）
	if currentRoleSortMin >= reqRoleSortMin {
		response.Fail(c, nil, "用户不能创建比自己等级高的或者相同等级的用户")
		return
	}

	// 创建用户
	user := model.User{
		Username:     req.Username,
		Password:     util.GenPasswd(req.Password),
		Mobile:       req.Mobile,
		Avatar:       req.Avatar,
		Nickname:     req.Nickname,
		Introduction: req.Introduction,
		Status:       req.Status,
		Creator:      ctxUser.Username,
		Roles:        roles,
	}

	err = uc.UserRepository.CreateUser(&user)
	if err != nil {
		response.Fail(c, nil, "创建用户失败: "+err.Error())
		return
	}
	response.Success(c, nil, "创建用户成功")

}

// 更新用户
func (uc UserController) UpdateUserById(c *gin.Context) {

}

func (uc UserController) BatchDeleteUserByIds(c *gin.Context) {

}
