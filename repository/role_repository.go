package repository

import (
	"go-lim/model"
	"go-lim/vo"
)

type IRoleRepository interface {
	GetRoles(role vo.RoleListRequest) (*[]model.Role, error)
	CreateRole(role vo.CreateRoleRequest) error
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

func (r RoleRepository) GetRoles(role vo.RoleListRequest) (*[]model.Role, error) {
	panic("implement me")
}

func (r RoleRepository) CreateRole(role vo.CreateRoleRequest) error {
	panic("implement me")
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
