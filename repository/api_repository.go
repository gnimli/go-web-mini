package repository

import (
	"go-lim/dto"
	"go-lim/model"
	"go-lim/vo"
)

type IApiRepository interface {
	GetApis(api *vo.ApiListRequest) ([]*model.Api, error)
	GetAllApiGroupByCategoryByRoleId() ([]*dto.ApiGroupByCategoryResponse ,[]uint, error)
	CreateApi(api vo.CreateApiRequest) error
	UpdateApiById(apiId uint, api vo.CreateApiRequest) error
	BatchDeleteApiByIds(apiIds []uint) error
}

type ApiRepository struct {

}

func NewApiRepository() IApiRepository{
	return ApiRepository{}
}

func (a ApiRepository) GetApis(api *vo.ApiListRequest) ([]*model.Api, error) {
	panic("implement me")
}

func (a ApiRepository) GetAllApiGroupByCategoryByRoleId() ([]*dto.ApiGroupByCategoryResponse, []uint, error) {
	panic("implement me")
}

func (a ApiRepository) CreateApi(api vo.CreateApiRequest) error {
	panic("implement me")
}

func (a ApiRepository) UpdateApiById(apiId uint, api vo.CreateApiRequest) error {
	panic("implement me")
}

func (a ApiRepository) BatchDeleteApiByIds(apiIds []uint) error {
	panic("implement me")
}
