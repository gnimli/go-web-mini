package repository

import (
	"go-web-mini/common"
	"go-web-mini/dto"
	"go-web-mini/model"
	"go-web-mini/vo"
)

type IApiRepository interface {
	GetApis(api *vo.ApiListRequest) ([]*model.Api, error)                                 // 获取接口列表
	GetApisById(apiIds []uint) ([]*model.Api, error)                                      // 根据接口ID获取接口列表
	GetAllApiGroupByCategoryByRoleId() ([]*dto.ApiGroupByCategoryResponse, []uint, error) // 查询指定角色的接口(以分类分组)
	CreateApi(api vo.CreateApiRequest) error                                              // 创建接口
	UpdateApiById(apiId uint, api vo.CreateApiRequest) error                              // 更新接口
	BatchDeleteApiByIds(apiIds []uint) error                                              //批量删除接口
	GetApiDescByPath(path string, method string) (string, error)                          // 根据接口路径和请求方式获取接口描述
}

type ApiRepository struct {
}

func NewApiRepository() IApiRepository {
	return ApiRepository{}
}

func (a ApiRepository) GetApis(api *vo.ApiListRequest) ([]*model.Api, error) {
	panic("implement me")
}

// 根据接口ID获取接口列表
func (a ApiRepository) GetApisById(apiIds []uint) ([]*model.Api, error) {
	var apis []*model.Api
	err := common.DB.Where("id IN (?)", apiIds).Find(&apis).Error
	return apis, err
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

func (a ApiRepository) GetApiDescByPath(path string, method string) (string, error) {
	var api model.Api
	err := common.DB.Where("path = ?", path).Where("method = ?", method).First(&api).Error
	return api.Desc, err
}
