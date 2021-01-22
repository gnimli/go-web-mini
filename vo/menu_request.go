package vo

type CreateMenuRequest struct {
	Name       string `json:"name" form:"name" validate:"required"`
	Title      string `json:"title" form:"title"`
	Icon       string `json:"icon" form:"icon"`
	Path       string `json:"path" form:"path"`
	Redirect   string `json:"redirect" form:"redirect"`
	Component  string `json:"component" form:"component"`
	Permission string `json:"permission" form:"permission"`
	Sort       uint   `json:"sort" form:"sort"`
	Status     uint   `json:"status" form:"status"`
	Visible    uint   `json:"visible" form:"visible"`
	Breadcrumb uint   `json:"breadcrumb" form:"breadcrumb"`
	ParentId   uint   `json:"parentId" form:"parentId"`
	Creator    string `json:"creator" form:"creator"`
}
