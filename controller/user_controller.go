package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-lim/common"
	"go-lim/dto"
	"go-lim/repository"
	"go-lim/response"
	"go-lim/util"
	"go-lim/vo"
)

type IUserController interface {
	GetUserInfo(c *gin.Context) // 获取当前登录用户信息
	GetUsers(c *gin.Context)    // 获取用户列表
	ChangePwd(c *gin.Context)   // 修改密码
	CreateUser(c *gin.Context)
	UpdateUserById(c *gin.Context)
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

func (uc UserController) CreateUser(c *gin.Context) {

}

func (uc UserController) UpdateUserById(c *gin.Context) {

}

func (uc UserController) BatchDeleteUserByIds(c *gin.Context) {

}
