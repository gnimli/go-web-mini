package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-mini/repository"
)

type IApiController interface {
	GetApis(c *gin.Context)
	GetAllApiGroupByCategoryByRoleId(c *gin.Context)
	CreateApi(c *gin.Context)
	UpdateApiById(c *gin.Context)
	BatchDeleteApiByIds(c *gin.Context)
}

type ApiController struct {
	ApiRepository repository.IApiRepository
}

func NewApiController() IApiController {
	apiRepository := repository.NewApiRepository()
	apiController := ApiController{ApiRepository: apiRepository}
	return apiController
}

func (a ApiController) GetApis(c *gin.Context) {
	panic("implement me")
}

func (a ApiController) GetAllApiGroupByCategoryByRoleId(c *gin.Context) {
	panic("implement me")
}

func (a ApiController) CreateApi(c *gin.Context) {
	panic("implement me")
}

func (a ApiController) UpdateApiById(c *gin.Context) {
	panic("implement me")
}

func (a ApiController) BatchDeleteApiByIds(c *gin.Context) {
	panic("implement me")
}
