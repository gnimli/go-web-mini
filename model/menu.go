package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name       string  `gorm:"type:varchar(50);comment:'菜单名称(英文名, 可用于国际化)'" json:"name"`
	Title      string  `gorm:"type:varchar(50);comment:'菜单标题(无法国际化时使用)'" json:"title"`
	Icon       string  `gorm:"type:varchar(50);comment:'菜单图标'" json:"icon"`
	Path       string  `gorm:"type:varchar(50);comment:'菜单访问路径'" json:"path"`
	Redirect   string  `gorm:"type:varchar(50);comment:'重定向路径'" json:"redirect"`
	Component  string  `gorm:"type:varchar(50);comment:'前端组件路径'" json:"component"`
	Permission string  `gorm:"type:varchar(50);comment:'权限标识'" json:"permission"`
	Sort       uint    `gorm:"type:int(3) unsigned;comment:'菜单顺序(同级菜单, 从0开始, 越小显示越靠前)'" json:"sort"`
	Status     uint    `gorm:"type:tinyint(1);default:1;comment:'菜单状态(正常/禁用, 默认正常)'" json:"status"` // 由于设置了默认值, 这里使用ptr, 可避免赋值失败
	Visible    uint    `gorm:"type:tinyint(1);default:1;comment:'菜单可见性(可见/隐藏, 默认可见)'" json:"visible"`
	Breadcrumb uint    `gorm:"type:tinyint(1);default:1;comment:'面包屑可见性(可见/隐藏, 默认可见)'" json:"breadcrumb"`
	ParentId   uint    `gorm:"default:0;comment:'父菜单编号(编号为0时表示根菜单)'" json:"parentId"`
	Creator    string  `gorm:"type:varchar(20);comment:'创建人'" json:"creator"`
	Children   []Menu  `gorm:"-" json:"children"`                  // 子菜单集合
	Roles      []*Role `gorm:"many2many:role_menus;" json:"roles"` // 角色菜单多对多关系
}
