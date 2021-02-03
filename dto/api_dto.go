package dto

import "go-web-mini/model"

type ApiGroupByCategoryResponse struct {
	Category string       `json:"category"`
	Children []*model.Api `json:"children"`
}
