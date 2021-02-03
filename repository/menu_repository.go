package repository

import (
	"go-web-mini/common"
	"go-web-mini/dto"
	"go-web-mini/model"
)

type IMenuRepository interface {
	GetMenus() ([]*model.Menu, error)                   // 获取菜单列表
	GetMenuTree() (dto.MenuTreeDto, error)              // 获取菜单树
	CreateMenu(menu *model.Menu) error                  // 创建菜单
	UpdateMenuById(menuId uint, menu *model.Menu) error // 更新菜单
	BatchDeleteMenuByIds(menuIds []uint) error          // 批量删除菜单
}

type MenuRepository struct {
}

func NewMenuRepository() IMenuRepository {
	return MenuRepository{}
}

// 获取菜单列表
func (m MenuRepository) GetMenus() ([]*model.Menu, error) {
	var menus []*model.Menu
	err := common.DB.Find(&menus).Error
	return menus, err
}

// 获取菜单树
func (m MenuRepository) GetMenuTree() (dto.MenuTreeDto, error) {
	var menus []*model.Menu
	err := common.DB.Find(&menus).Error
	return dto.ToMenuTreeDto(menus), err
}

// 创建菜单
func (m MenuRepository) CreateMenu(menu *model.Menu) error {
	err := common.DB.Create(menu).Error
	return err
}

// 更新菜单
func (m MenuRepository) UpdateMenuById(menuId uint, menu *model.Menu) error {
	err := common.DB.Model(menu).Where("id = ?", menuId).Updates(menu).Error
	return err
}

// 批量删除菜单
func (m MenuRepository) BatchDeleteMenuByIds(menuIds []uint) error {
	err := common.DB.Where("id IN (?)", menuIds).Delete(&model.Menu{}).Error
	return err
}
