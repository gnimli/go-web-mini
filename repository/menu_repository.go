package repository

import (
	"go-web-base/model"
	"go-web-base/vo"
)

type IMenuRepository interface {
	// 查询当前用户菜单树
	GetMenuTree(roleId uint) (*[]model.Menu, error)
	// 查询指定角色的菜单树
	GetAllMenuByRoleId(roleId uint) (*[]model.Menu, error)
	// 查询所有菜单
	GetMenus() (*[]model.Menu, error)
	// 创建菜单
	CreateMenu(menu vo.CreateMenuRequest) error
	UpdateMenuById(menuId uint, menu vo.CreateMenuRequest) error
	BatchDeleteMenuByIds(ids []uint) error
}

type MenuRepository struct {
}

func NewMenuRepository() IMenuRepository {
	return MenuRepository{}
}

func (m MenuRepository) GetMenuTree(roleId uint) (*[]model.Menu, error) {
	panic("implement me")
}

func (m MenuRepository) GetAllMenuByRoleId(roleId uint) (*[]model.Menu, error) {
	panic("implement me")
}

func (m MenuRepository) GetMenus() (*[]model.Menu, error) {
	panic("implement me")
}

func (m MenuRepository) CreateMenu(menu vo.CreateMenuRequest) error {
	panic("implement me")
}

func (m MenuRepository) UpdateMenuById(menuId uint, menu vo.CreateMenuRequest) error {
	panic("implement me")
}

func (m MenuRepository) BatchDeleteMenuByIds(ids []uint) error {
	panic("implement me")
}
