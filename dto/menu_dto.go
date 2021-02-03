package dto

import (
	"go-web-mini/model"
	"gorm.io/gorm"
)

// 返回给前端的菜单树
type MenuTreeDto struct {
	gorm.Model
	Name       string        `json:"name"`
	Title      string        `json:"title"`
	Icon       string        `json:"icon"`
	Path       string        `json:"path"`
	Redirect   string        `json:"redirect"`
	Component  string        `json:"component"`
	Permission string        `json:"permission"`
	Sort       uint          `json:"sort"`
	Status     uint          `json:"status"`
	Visible    uint          `json:"visible"`
	Breadcrumb uint          `json:"breadcrumb"`
	ParentId   uint          `json:"parentId"`
	Creator    string        `json:"creator"`
	Children   []*model.Menu `json:"children"` // 子菜单集合
}

func ToMenuTreeDto([]*model.Menu) MenuTreeDto {
	return MenuTreeDto{}
}
