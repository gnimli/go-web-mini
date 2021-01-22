package repository

import (
	"go-lim/common"
	"go-lim/dto"
	"go-lim/model"
	"go-lim/vo"
)

type IApiRepository interface {
	GetApis(api *vo.ApiListRequest) ([]*model.Api, error)                                 // 获取接口列表
	GetAllApiGroupByCategoryByRoleId() ([]*dto.ApiGroupByCategoryResponse, []uint, error) // 查询指定角色的接口(以分类分组)
	CreateApi(api vo.CreateApiRequest) error                                              // 创建接口
	UpdateApiById(apiId uint, api vo.CreateApiRequest) error                              // 更新接口
	BatchDeleteApiByIds(apiIds []uint) error                                              //批量删除接口
	GetApiDescByPath(path string) (string, error)
}

type ApiRepository struct {
}

func NewApiRepository() IApiRepository {
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

func (a ApiRepository) GetApiDescByPath(path string) (string, error) {
	var api model.Api
	err := common.DB.Where("path = ?", path).First(&api).Error
	return api.Desc, err
}
