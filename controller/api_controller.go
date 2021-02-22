package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-web-mini/common"
	"go-web-mini/model"
	"go-web-mini/repository"
	"go-web-mini/response"
	"go-web-mini/vo"
	"strconv"
)

type IApiController interface {
	GetApis(c *gin.Context)             // 获取接口列表
	GetApiTree(c *gin.Context)          // 获取接口树(按接口Category字段分类)
	CreateApi(c *gin.Context)           // 创建接口
	UpdateApiById(c *gin.Context)       // 更新接口
	BatchDeleteApiByIds(c *gin.Context) // 批量删除接口
}

type ApiController struct {
	ApiRepository repository.IApiRepository
}

func NewApiController() IApiController {
	apiRepository := repository.NewApiRepository()
	apiController := ApiController{ApiRepository: apiRepository}
	return apiController
}

// 获取接口列表
func (ac ApiController) GetApis(c *gin.Context) {
	var req vo.ApiListRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := common.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(common.Trans)
		response.Fail(c, nil, errStr)
		return
	}
	// 获取
	apis, total, err := ac.ApiRepository.GetApis(&req)
	if err != nil {
		response.Fail(c, nil, "获取接口列表失败")
		return
	}
	response.Success(c, gin.H{
		"apis": apis, "total": total,
	}, "获取接口列表成功")
}

// 获取接口树(按接口Category字段分类)
func (ac ApiController) GetApiTree(c *gin.Context) {
	tree, err := ac.ApiRepository.GetApiTree()
	if err != nil {
		response.Fail(c, nil, "获取接口树失败")
		return
	}
	response.Success(c, gin.H{
		"apiTree": tree,
	}, "获取接口树成功")
}

// 创建接口
func (ac ApiController) CreateApi(c *gin.Context) {
	var req vo.CreateApiRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := common.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(common.Trans)
		response.Fail(c, nil, errStr)
		return
	}

	// 获取当前用户
	ur := repository.NewUserRepository()
	ctxUser, err := ur.GetCurrentUser(c)
	if err != nil {
		response.Fail(c, nil, "获取当前用户信息失败")
		return
	}

	api := model.Api{
		Method:   req.Method,
		Path:     req.Path,
		Category: req.Category,
		Desc:     req.Desc,
		Creator:  ctxUser.Username,
	}

	// 创建接口
	err = ac.ApiRepository.CreateApi(&api)
	if err != nil {
		response.Fail(c, nil, "创建接口失败: "+err.Error())
		return
	}

	response.Success(c, nil, "创建接口成功")
	return
}

// 更新接口
func (ac ApiController) UpdateApiById(c *gin.Context) {
	var req vo.UpdateApiRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := common.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(common.Trans)
		response.Fail(c, nil, errStr)
		return
	}

	// 获取路径中的apiId
	apiId, _ := strconv.Atoi(c.Param("apiId"))
	if apiId <= 0 {
		response.Fail(c, nil, "接口ID不正确")
		return
	}

	// 获取当前用户
	ur := repository.NewUserRepository()
	ctxUser, err := ur.GetCurrentUser(c)
	if err != nil {
		response.Fail(c, nil, "获取当前用户信息失败")
		return
	}

	api := model.Api{
		Method:   req.Method,
		Path:     req.Path,
		Category: req.Category,
		Desc:     req.Desc,
		Creator:  ctxUser.Username,
	}

	err = ac.ApiRepository.UpdateApiById(uint(apiId), &api)
	if err != nil {
		response.Fail(c, nil, "更新接口失败: "+err.Error())
		return
	}

	response.Success(c, nil, "更新接口成功")
}

// 批量删除接口
func (ac ApiController) BatchDeleteApiByIds(c *gin.Context) {
	var req vo.DeleteApiRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := common.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(common.Trans)
		response.Fail(c, nil, errStr)
		return
	}

	// 删除接口
	err := ac.ApiRepository.BatchDeleteApiByIds(req.ApiIds)
	if err != nil {
		response.Fail(c, nil, "删除接口失败: "+err.Error())
		return
	}

	response.Success(c, nil, "删除接口成功")
}
