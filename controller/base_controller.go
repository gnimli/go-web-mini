package controller

import (
	"github.com/gin-gonic/gin"
	"go-lim/repository"
)

type IBaseController interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type BaseController struct {
	BaseRepository repository.IBaseRepository
}

// BaseController 构造函数
func NewBaseController() IBaseController {
	baseRepository := repository.NewBaseRepository()
	baseController := BaseController{BaseRepository: baseRepository}
	return baseController
}

func (b BaseController) Login(c *gin.Context) {

}

func (b BaseController) Logout(c *gin.Context) {

}

func (b BaseController) RefreshToken(c *gin.Context) {

}
