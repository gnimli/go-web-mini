package repository

import (
	"go-lim/model"
	"go-lim/vo"
)

type IUserRepository interface {
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
