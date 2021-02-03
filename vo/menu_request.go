package vo

// 创建接口结构体
type CreateMenuRequest struct {
	Name       string `json:"name" form:"name" validate:"required"`
	Title      string `json:"title" form:"title" validate:"required"`
	Icon       string `json:"icon" form:"icon"`
	Path       string `json:"path" form:"path" validate:"required"`
	Redirect   string `json:"redirect" form:"redirect"`
	Component  string `json:"component" form:"component" validate:"required"`
	Permission string `json:"permission" form:"permission"`
	Sort       uint   `json:"sort" form:"sort" validate:"gte=1,lte=999"`
	Status     uint   `json:"status" form:"status" validate:"oneof=1 2"`
	Visible    uint   `json:"visible" form:"visible" validate:"oneof=1 2"`
	Breadcrumb uint   `json:"breadcrumb" form:"breadcrumb"`
	ParentId   uint   `json:"parentId" form:"parentId"`
}

// 更新接口结构体
type UpdateMenuRequest struct {
	Name       string `json:"name" form:"name"`
	Title      string `json:"title" form:"title"`
	Icon       string `json:"icon" form:"icon"`
	Path       string `json:"path" form:"path"`
	Redirect   string `json:"redirect" form:"redirect"`
	Component  string `json:"component" form:"component"`
	Permission string `json:"permission" form:"permission"`
	Sort       uint   `json:"sort" form:"sort" validate:"gte=1,lte=999"`
	Status     uint   `json:"status" form:"status" validate:"oneof=1 2"`
	Visible    uint   `json:"visible" form:"visible" validate:"oneof=1 2"`
	Breadcrumb uint   `json:"breadcrumb" form:"breadcrumb"`
	ParentId   uint   `json:"parentId" form:"parentId"`
}

// 删除接口结构体
type DeleteMenuRequest struct {
	MenuIds []uint `json:"menuIds" form:"menuIds"`
}
