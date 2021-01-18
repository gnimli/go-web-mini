package vo

type CreateMenuRequest struct {
	Name       string `json:"name" binding:"required"`
	Title      string `json:"title" `
	Icon       string `json:"icon" `
	Path       string `json:"path" `
	Redirect   string `json:"redirect"`
	Component  string `json:"component"`
	Permission string `json:"permission"`
	Sort       uint   `json:"sort"`
	Status     uint   `json:"status"`
	Visible    uint   `json:"visible"`
	Breadcrumb uint   `json:"breadcrumb"`
	ParentId   uint   `json:"parentId"`
	Creator    string `json:"creator"`
}
