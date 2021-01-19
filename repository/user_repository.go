package repository

import (
	"errors"
	"go-lim/common"
	"go-lim/model"
	"go-lim/util"
	"go-lim/vo"
)

type IUserRepository interface {
	// 登录
	Login(user *model.User) (*model.User, error)

	// 用户
	GetUserInfo()
	GetUsers(user *vo.UserListRequest) (*[]model.User, error)
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
func (u UserRepository) Login(user *model.User) (*model.User, error) {
	// 根据用户名查询用户
	var firstUser model.User
	err := common.DB.Debug().Where("username = ?", user.Username).Preload("Roles").First(&firstUser).Error
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

func (u UserRepository) GetUserInfo() {
	panic("implement me")
}

func (u UserRepository) GetUsers(user *vo.UserListRequest) (*[]model.User, error) {
	panic("implement me")
}

func (u UserRepository) ChangePwd(pwd *vo.ChangePwdRequest) error {
	panic("implement me")
}

func (u UserRepository) CreateUser(user *vo.CreateUserRequest) (model.User, error) {
	panic("implement me")
}

func (u UserRepository) UpdateUserById(id string, user *vo.CreateUserRequest) (model.User, error) {
	panic("implement me")
}

func (u UserRepository) BatchDeleteUserByIds(ids []string) error {
	panic("implement me")
}
