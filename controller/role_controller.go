package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-lim/common"
	"go-lim/model"
	"go-lim/repository"
	"go-lim/response"
	"go-lim/vo"
)

type IRoleController interface {
	GetRoles(c *gin.Context)   // 获取角色列表
	CreateRole(c *gin.Context) // 创建角色
	UpdateRoleById(c *gin.Context)
	UpdateRoleMenusById(c *gin.Context)
	UpdateRoleApisById(c *gin.Context)
	BatchDeleteRoleByIds(c *gin.Context)
}

type RoleController struct {
	RoleRepository repository.IRoleRepository
}

func NewRoleController() IRoleController {
	roleRepository := repository.NewRoleRepository()
	roleController := RoleController{RoleRepository: roleRepository}
	return roleController
}

// 获取角色列表
func (rc RoleController) GetRoles(c *gin.Context) {
	var req vo.RoleListRequest
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

	// 查询角色列表
	roles, total, err := rc.RoleRepository.GetRoles(&req)
	if err != nil {
		response.Fail(c, nil, "查询角色列表失败: "+err.Error())
		return
	}
	response.Success(c, gin.H{"roles": roles, "total": total}, "查询角色列表成功")
}

// 创建角色
func (rc RoleController) CreateRole(c *gin.Context) {
	var req vo.CreateRoleRequest
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

	// 获取当前用户最高角色等级
	uc := repository.NewUserRepository()
	sort, ctxUser, err := uc.GetCurrentUserMinRoleSort(c)
	if err != nil {
		response.Fail(c, nil, "获取当前用户最高角色等级失败: "+err.Error())
		return
	}

	if req.Sort == 0 {
		req.Sort = 999
	}
	// 用户不能创建比自己等级高或相同等级的角色
	if sort >= req.Sort {
		response.Fail(c, nil, "不能创建比自己等级高或相同等级的角色")
		return
	}

	role := model.Role{
		Name:    req.Name,
		Keyword: req.Keyword,
		Desc:    req.Desc,
		Status:  req.Status,
		Sort:    req.Sort,
		Creator: ctxUser.Username,
	}

	// 创建角色
	err = rc.RoleRepository.CreateRole(&role)
	if err != nil {
		response.Fail(c, nil, "创建角色失败: "+err.Error())
		return
	}
	response.Success(c, nil, "创建角色成功")

}

func (rc RoleController) UpdateRoleById(c *gin.Context) {
	panic("implement me")
}

func (rc RoleController) UpdateRoleMenusById(c *gin.Context) {
	panic("implement me")
}

func (rc RoleController) UpdateRoleApisById(c *gin.Context) {
	panic("implement me")
}

func (rc RoleController) BatchDeleteRoleByIds(c *gin.Context) {
	panic("implement me")
}
