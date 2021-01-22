package repository

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-lim/common"
	"go-lim/model"
	"go-lim/util"
	"go-lim/vo"
	"strings"
)

type IUserRepository interface {
	Login(user *model.User) (*model.User, error) // 登录
	GetCurrentUser(c *gin.Context) model.User    // 获取当前登录用户信息
	GetUserById(id uint) (model.User, error)     // 获取单个用户
	GetUsers(req *vo.UserListRequest) ([]model.User, int64, error)
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

// 获取当前登录用户信息
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

func (ur UserRepository) GetUsers(req *vo.UserListRequest) ([]model.User, int64, error) {
	var list []model.User
	db := common.DB.Model(&model.User{}).Order("created_at DESC")

	username := strings.TrimSpace(req.Username)
	if username != "" {
		db = db.Where("username LIKE ?", fmt.Sprintf("%s%%", username))
	}
	nickname := strings.TrimSpace(req.Nickname)
	if nickname != "" {
		db = db.Where("nickname LIKE ?", fmt.Sprintf("%s%%", nickname))
	}
	mobile := strings.TrimSpace(req.Mobile)
	if mobile != "" {
		db = db.Where("mobile LIKE ?", fmt.Sprintf("%s%%", mobile))
	}
	status := req.Status
	if status == 0 || status == 1 {
		db = db.Where("username LIKE ?", fmt.Sprintf("%s%%", username))
	}
	// 当pageNum > 0 且 pageSize > 0 才分页
	//记录总条数
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	pageNum := req.PageNum
	pageSize := req.PageSize
	if pageNum > 0 && pageSize > 0 {
		err = db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}
	return list, total, err
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
