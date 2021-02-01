package repository

import (
	"fmt"
	"go-web-mini/common"
	"go-web-mini/model"
	"go-web-mini/vo"
	"strings"
)

type IRoleRepository interface {
	GetRoles(req *vo.RoleListRequest) ([]model.Role, int64, error) // 获取角色列表
	GetRolesByIds(roleIds []uint) ([]*model.Role, error)           // 根据角色ID查询角色
	CreateRole(role *model.Role) error                             // 创建角色
	UpdateRoleById(roleId uint, role *model.Role) error            // 修改角色
	UpdateRoleMenusById(roleId uint, menuIds vo.UpdateRoleMenusRequest) error
	UpdateRoleApisById(roleId uint, apiIds vo.UpdateRoleApisRequest) error
	BatchDeleteRoleByIds(roleIds []uint) error // 删除角色
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

// 创建角色
func (r RoleRepository) CreateRole(role *model.Role) error {
	err := common.DB.Debug().Create(role).Error
	return err
}

// 修改角色
func (r RoleRepository) UpdateRoleById(roleId uint, role *model.Role) error {
	err := common.DB.Model(&model.Role{}).Where("id = ?", roleId).Updates(role).Error
	return err
}

func (r RoleRepository) UpdateRoleMenusById(roleId uint, menuIds vo.UpdateRoleMenusRequest) error {
	panic("implement me")
}

func (r RoleRepository) UpdateRoleApisById(roleId uint, apiIds vo.UpdateRoleApisRequest) error {
	panic("implement me")
}

// 删除角色
func (r RoleRepository) BatchDeleteRoleByIds(roleIds []uint) error {
	err := common.DB.Where("id IN (?)", roleIds).Delete(&model.Role{}).Error
	return err
}
