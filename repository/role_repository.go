package repository

import (
	"errors"
	"fmt"
	"go-web-mini/common"
	"go-web-mini/model"
	"go-web-mini/vo"
	"strings"
)

type IRoleRepository interface {
	GetRoles(req *vo.RoleListRequest) ([]model.Role, int64, error) // 获取角色列表
	GetRolesByIds(roleIds []uint) ([]*model.Role, error)           // 根据角色ID获取角色
	CreateRole(role *model.Role) error                             // 创建角色
	UpdateRoleById(roleId uint, role *model.Role) error            // 更新角色
	UpdateRoleMenusById(roleId uint, menuIds vo.UpdateRoleMenusRequest) error
	UpdateRoleApis(roleKeyword string, reqRolePolicies [][]string) error // 更新角色的权限接口（先全部删除再新增）
	BatchDeleteRoleByIds(roleIds []uint) error                           // 删除角色
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

//根据角色ID获取角色
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

// 更新角色
func (r RoleRepository) UpdateRoleById(roleId uint, role *model.Role) error {
	err := common.DB.Model(&model.Role{}).Where("id = ?", roleId).Updates(role).Error
	return err
}

func (r RoleRepository) UpdateRoleMenusById(roleId uint, menuIds vo.UpdateRoleMenusRequest) error {
	panic("implement me")
}

// 更新角色的权限接口（先全部删除再新增）
func (r RoleRepository) UpdateRoleApis(roleKeyword string, reqRolePolicies [][]string) error {
	// 先获取path中的角色ID对应角色已有的police(需要先删除的)
	rmPolicies := common.CasbinEnforcer.GetFilteredPolicy(0, roleKeyword)
	if len(rmPolicies) > 0 {
		isRemoved, _ := common.CasbinEnforcer.RemovePolicies(rmPolicies)
		if !isRemoved {
			return errors.New("更新角色的权限接口失败")
		}
	}
	isAdded, _ := common.CasbinEnforcer.AddPolicies(reqRolePolicies)
	if !isAdded {
		return errors.New("更新角色的权限接口失败")
	}
	err := common.CasbinEnforcer.LoadPolicy()
	if err != nil {
		return errors.New("更新角色的权限接口成功，角色的权限接口策略加载失败")
	} else {
		return err
	}
}

// 删除角色
func (r RoleRepository) BatchDeleteRoleByIds(roleIds []uint) error {
	err := common.DB.Where("id IN (?)", roleIds).Delete(&model.Role{}).Error
	return err
}
