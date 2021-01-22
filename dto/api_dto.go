package dto

import "go-lim/model"

// 权限接口信息响应, 字段含义见models
type ApiGroupByCategoryResponse struct {
	//Title    string                  `json:"title"`    // 标题
	Category string       `json:"category"` // 分组名称
	Children []*model.Api `json:"children"` // 前端以树形图结构展示, 这里用children表示
}
