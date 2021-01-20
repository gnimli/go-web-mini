package repository

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-lim/common"
	"go-lim/model"
	"go-lim/util"
	"go-lim/vo"
)

type IUserRepository interface {
	Login(user *model.User) (*model.User, error) // 登录
	GetCurrentUser(c *gin.Context) model.User    // 获取当前用户信息
	GetUserById(id uint) (model.User, error)     // 获取单个用户
	GetUsers() (*[]model.User, error)
	ChangePwd(pwd *vo.ChangePwdRequest) error
	CreateUser(user *vo.CreateUserRequest) (model.User, error)
	UpdateUserById(id string, user *vo.CreateUserRequest) (model.User, error)
	BatchDeleteUserByIds(ids []string) error
}

type UserRepository struct {
}

// UserRepository构造函数
func NewUserRepository() IUserRepository {
	return UserRepository{}
}

// 登录
func (ur UserRepository) Login(user *model.User) (*model.User, error) {
	// 根据用户名查询用户
	var firstUser model.User
	err := common.DB.Where("username = ?", user.Username).Preload("Roles").First(&firstUser).Error
	//fmt.Println("firstUser---")
	//fmt.Printf("%+v", firstUser)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	// 校验密码
	err = util.ComparePasswd(firstUser.Password, user.Password)
	if err != nil {
		return &firstUser, errors.New("密码错误")
	}
	return &firstUser, nil
}

// 获取当前用户信息
func (ur UserRepository) GetCurrentUser(c *gin.Context) model.User {
	var newUser model.User
	ctxUser, exist := c.Get("user")
	if !exist {
		return newUser
	}
	u, _ := ctxUser.(model.User)
	user, _ := ur.GetUserById(u.ID)
	return user
}

// 获取单个用户
func (ur UserRepository) GetUserById(id uint) (model.User, error) {
	var user model.User
	err := common.DB.Where("id = ?", id).Preload("Roles").First(&user).Error
	return user, err
}

func (ur UserRepository) GetUsers() (*[]model.User, error) {
	panic("implement me")
}

func (ur UserRepository) ChangePwd(pwd *vo.ChangePwdRequest) error {
	panic("implement me")
}

func (ur UserRepository) CreateUser(user *vo.CreateUserRequest) (model.User, error) {
	panic("implement me")
}

func (ur UserRepository) UpdateUserById(id string, user *vo.CreateUserRequest) (model.User, error) {
	panic("implement me")
}

func (ur UserRepository) BatchDeleteUserByIds(ids []string) error {
	panic("implement me")
}
