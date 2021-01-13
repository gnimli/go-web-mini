package repository

import (
	"go-lim/model"
)

type IBaseRepository interface {
	Login(user model.User) (*model.User, error)
	Logout(user model.User) (*model.User, error)
	RefreshToken(user model.User) (*model.User, error)
}

type BaseRepository struct {
}

// BaseRepository构造函数
func NewBaseRepository() IBaseRepository {
	return BaseRepository{}
}

func (b BaseRepository) Login(user model.User) (*model.User, error) {
	return nil, nil
}

func (b BaseRepository) Logout(user model.User) (*model.User, error) {
	return nil, nil
}

func (b BaseRepository) RefreshToken(user model.User) (*model.User, error) {
	return nil, nil
}
