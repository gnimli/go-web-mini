package dto

import "go-web-mini/model"

type ApiTreeDto struct {
	Category string       `json:"category"`
	Children []*model.Api `json:"children"`
}
