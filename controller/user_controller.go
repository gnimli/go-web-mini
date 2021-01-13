package controller

import (
	"github.com/gin-gonic/gin"
	"go-lim/repository"
)

type IUserController interface {
	GetUserInfo(c *gin.Context)
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

func (u UserController) GetUserInfo(c *gin.Context) {

}

func (u UserController) GetUsers(c *gin.Context) {

}

func (u UserController) ChangePwd(c *gin.Context) {

}

func (u UserController) CreateUser(c *gin.Context) {

}

func (u UserController) UpdateUserById(c *gin.Context) {

}

func (u UserController) BatchDeleteUserByIds(c *gin.Context) {

}
