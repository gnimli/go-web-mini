package vo

// 创建接口结构体
type CreateMenuRequest struct {
	Name       string `json:"name" form:"name" validate:"required,min=1,max=50"`
	Title      string `json:"title" form:"title" validate:"required,min=1,max=50"`
	Icon       string `json:"icon" form:"icon" validate:"min=0,max=50"`
	Path       string `json:"path" form:"path" validate:"required,min=1,max=100"`
	Redirect   string `json:"redirect" form:"redirect" validate:"min=0,max=100"`
	Component  string `json:"component" form:"component" validate:"required,min=1,max=100"`
	Sort       uint   `json:"sort" form:"sort" validate:"gte=1,lte=999"`
	Status     uint   `json:"status" form:"status" validate:"oneof=1 2"`
	Hidden     uint   `json:"hidden" form:"hidden" validate:"oneof=1 2"`
	NoCache    uint   `json:"noCache" form:"noCache" validate:"oneof=1 2"`
	AlwaysShow uint   `json:"alwaysShow" form:"alwaysShow" validate:"oneof=1 2"`
	Breadcrumb uint   `json:"breadcrumb" form:"breadcrumb" validate:"oneof=1 2"`
	ActiveMenu string `json:"activeMenu" form:"activeMenu" validate:"min=0,max=100"`
	ParentId   uint   `json:"parentId" form:"parentId"`
}

// 更新接口结构体
type UpdateMenuRequest struct {
	Name       string `json:"name" form:"name" validate:"required,min=1,max=50"`
	Title      string `json:"title" form:"title" validate:"required,min=1,max=50"`
	Icon       string `json:"icon" form:"icon" validate:"min=0,max=50"`
	Path       string `json:"path" form:"path" validate:"required,min=1,max=100"`
	Redirect   string `json:"redirect" form:"redirect" validate:"min=0,max=100"`
	Component  string `json:"component" form:"component" validate:"min=0,max=100"`
	Sort       uint   `json:"sort" form:"sort" validate:"gte=1,lte=999"`
	Status     uint   `json:"status" form:"status" validate:"oneof=1 2"`
	Hidden     uint   `json:"hidden" form:"hidden" validate:"oneof=1 2"`
	NoCache    uint   `json:"noCache" form:"noCache" validate:"oneof=1 2"`
	AlwaysShow uint   `json:"alwaysShow" form:"alwaysShow" validate:"oneof=1 2"`
	Breadcrumb uint   `json:"breadcrumb" form:"breadcrumb" validate:"oneof=1 2"`
	ActiveMenu string `json:"activeMenu" form:"activeMenu" validate:"min=0,max=100"`
	ParentId   uint   `json:"parentId" form:"parentId"`
}

// 删除接口结构体
type DeleteMenuRequest struct {
	MenuIds []uint `json:"menuIds" form:"menuIds"`
}
