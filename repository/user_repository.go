package repository

import (
	"go-lim/model"
	"go-lim/vo"
)

type IUserRepository interface {
	// 登录
	Login(user *model.User) (*model.User, error)
	Logout(user *model.User) (*model.User, error)
	RefreshToken(user *model.User) (*model.User, error)

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
	panic("implement me")
}

func (u UserRepository) Logout(user *model.User) (*model.User, error) {
	panic("implement me")
}

func (u UserRepository) RefreshToken(user *model.User) (*model.User, error) {
	panic("implement me")
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
