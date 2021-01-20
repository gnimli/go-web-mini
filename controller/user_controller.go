package controller

import (
	"github.com/gin-gonic/gin"
	"go-lim/dto"
	"go-lim/repository"
	"go-lim/response"
)

type IUserController interface {
	GetUserInfo(c *gin.Context) // 获取当前用户信息
	GetUsers(c *gin.Context)
	ChangePwd(c *gin.Context)
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

// 获取当前用户信息
func (uc UserController) GetUserInfo(c *gin.Context) {
	user := uc.UserRepository.GetCurrentUser(c)
	userInfoDto := dto.ToUserInfoDto(user)
	response.Success(c, gin.H{
		"userInfo": userInfoDto,
	}, "获取当前用户信息成功")
}

func (uc UserController) GetUsers(c *gin.Context) {

}

func (uc UserController) ChangePwd(c *gin.Context) {

}

func (uc UserController) CreateUser(c *gin.Context) {

}

func (uc UserController) UpdateUserById(c *gin.Context) {

}

func (uc UserController) BatchDeleteUserByIds(c *gin.Context) {

}
