package repository

import (
	"fmt"
	"go-lim/common"
	"go-lim/model"
	"go-lim/vo"
	"strings"
)

type IRoleRepository interface {
	GetRoles(req *vo.RoleListRequest) ([]model.Role, int64, error) // 获取角色列表
	GetRolesByIds(roleIds []uint) ([]*model.Role, error)           //根据角色ID查询角色
	CreateRole(role *model.Role) error
	UpdateRoleById(roleId uint, role vo.CreateRoleRequest) error
	UpdateRoleMenusById(roleId uint, menuIds vo.UpdateRoleMenusRequest) error
	UpdateRoleApisById(roleId uint, apiIds vo.UpdateRoleApisRequest) error
	BatchDeleteRoleByIds(roleIds []uint) error
}

type RoleRepository struct {
}

func NewRoleRepository() IRoleRepository {
	return RoleRepository{}
}

// 获取角色列表
func (r RoleRepository) GetRoles(req *vo.RoleListRequest) ([]model.Role, int64, error) {
	var list []model.Role
	db := common.DB.Model(&model.Role{}).Order("created_at DESC")

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	keyword := strings.TrimSpace(req.Keyword)
	if keyword != "" {
		db = db.Where("keyword LIKE ?", fmt.Sprintf("%%%s%%", keyword))
	}
	status := req.Status
	if status != 0 {
		db = db.Where("status = ?", status)
	}
	// 当pageNum > 0 且 pageSize > 0 才分页
	//记录总条数
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	pageNum := int(req.PageNum)
	pageSize := int(req.PageSize)
	if pageNum > 0 && pageSize > 0 {
		err = db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}
	return list, total, err
}

//根据角色ID查询角色
func (r RoleRepository) GetRolesByIds(roleIds []uint) ([]*model.Role, error) {
	var list []*model.Role
	err := common.DB.Where("id IN (?)", roleIds).Find(&list).Error
	return list, err
}

func (r RoleRepository) CreateRole(role *model.Role) error {
	err := common.DB.Debug().Create(role).Error
	return err
}

func (r RoleRepository) UpdateRoleById(roleId uint, role vo.CreateRoleRequest) error {
	panic("implement me")
}

func (r RoleRepository) UpdateRoleMenusById(roleId uint, menuIds vo.UpdateRoleMenusRequest) error {
	panic("implement me")
}

func (r RoleRepository) UpdateRoleApisById(roleId uint, apiIds vo.UpdateRoleApisRequest) error {
	panic("implement me")
}

func (r RoleRepository) BatchDeleteRoleByIds(roleIds []uint) error {
	panic("implement me")
}
