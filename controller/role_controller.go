package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-lim/common"
	"go-lim/model"
	"go-lim/repository"
	"go-lim/response"
	"go-lim/vo"
	"strconv"
)

type IRoleController interface {
	GetRoles(c *gin.Context)       // 获取角色列表
	CreateRole(c *gin.Context)     // 创建角色
	UpdateRoleById(c *gin.Context) // 修改角色
	UpdateRoleMenusById(c *gin.Context)
	UpdateRoleApisById(c *gin.Context)   // 更新角色的权限接口
	BatchDeleteRoleByIds(c *gin.Context) // 删除角色
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

// 修改角色
func (rc RoleController) UpdateRoleById(c *gin.Context) {
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
	// 获取path中的roleId
	roleId, _ := strconv.Atoi(c.Param("roleId"))
	if roleId <= 0 {
		response.Fail(c, nil, "角色ID不正确")
		return
	}

	if req.Sort == 0 {
		req.Sort = 999
	}

	// 当前用户角色排序最小值（最高等级角色）以及当前用户
	uc := repository.NewUserRepository()
	minSort, _, err := uc.GetCurrentUserMinRoleSort(c)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	// 不能修改比自己角色等级高或相等的角色
	// 根据path中的角色ID查询该角色信息
	roles, err := rc.RoleRepository.GetRolesByIds([]uint{uint(roleId)})
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	if len(roles) == 0 {
		response.Fail(c, nil, "未查询到角色信息")
		return
	}
	if minSort >= roles[0].Sort {
		response.Fail(c, nil, "不能修改比自己角色等级高或相等的角色")
		return
	}

	// 不能把角色等级修改得比当前用户的等级高
	if minSort >= req.Sort {
		response.Fail(c, nil, "不能把角色等级修改得比当前用户的等级高或相同")
		return
	}

	role := model.Role{
		Name:    req.Name,
		Keyword: req.Keyword,
		Desc:    req.Desc,
		Status:  req.Status,
		Sort:    req.Sort,
	}

	// 修改角色
	err = rc.RoleRepository.UpdateRoleById(uint(roleId), &role)
	if err != nil {
		response.Fail(c, nil, "修改角色失败: "+err.Error())
		return
	}

	// 如果修改成功，且修改了角色的keyword, 则更新casbin中policy
	if req.Keyword != roles[0].Keyword {
		// 获取policy
		rolePolicies := common.CasbinEnforcer.GetFilteredPolicy(0, roles[0].Keyword)
		rolePoliciesCopy := make([][]string, 0)
		// 替换keyword
		for _, policy := range rolePolicies {
			policyCopy := make([]string, len(policy))
			copy(policyCopy, policy)
			rolePoliciesCopy = append(rolePoliciesCopy, policyCopy)
			policy[0] = req.Keyword
		}

		//gormadapter未实现UpdatePolicies方法，等gorm更新---
		//isUpdated, _ := common.CasbinEnforcer.UpdatePolicies(rolePoliciesCopy, rolePolicies)
		//if !isUpdated {
		//	response.Fail(c, nil, "修改角色成功，但角色关键字关联的权限接口更新失败！")
		//	return
		//}

		// 这里需要先新增再删除（先删除再增加会出错）
		isAdded, _ := common.CasbinEnforcer.AddPolicies(rolePolicies)
		if !isAdded {
			response.Fail(c, nil, "修改角色成功，但角色关键字关联的权限接口更新失败")
			return
		}
		isRemoved, _ := common.CasbinEnforcer.RemovePolicies(rolePoliciesCopy)
		if !isRemoved {
			response.Fail(c, nil, "修改角色成功，但角色关键字关联的权限接口更新失败")
			return
		}
		err := common.CasbinEnforcer.LoadPolicy()
		if err != nil {
			response.Fail(c, nil, "修改角色成功，但角色关键字关联的权限接口加载失败")
			return
		}

	}

	// 修改角色成功可以更新用户信息缓存，但是这里没有更新，casbin中间件会拦截非法操作

	response.Success(c, nil, "修改角色成功")
}

func (rc RoleController) UpdateRoleMenusById(c *gin.Context) {
	panic("implement me")
}

// 更新角色的权限接口
func (rc RoleController) UpdateRoleApisById(c *gin.Context) {
	var req vo.UpdateRoleApisRequest
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

	// 获取path中的roleId
	roleId, _ := strconv.Atoi(c.Param("roleId"))
	if roleId <= 0 {
		response.Fail(c, nil, "角色ID不正确")
		return
	}
	// 根据path中的角色ID查询该角色信息
	roles, err := rc.RoleRepository.GetRolesByIds([]uint{uint(roleId)})
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	if len(roles) == 0 {
		response.Fail(c, nil, "未查询到角色信息")
		return
	}

	// 当前用户角色排序最小值（最高等级角色）以及当前用户
	ur := repository.NewUserRepository()
	minSort, _, err := ur.GetCurrentUserMinRoleSort(c)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	// 不能修改比自己角色等级高或相等角色的权限接口
	if minSort >= roles[0].Sort {
		response.Fail(c, nil, "不能修改比自己角色等级高或相等角色的权限接口")
		return
	}

	// 不能把角色的权限接口设置的比当前用户所拥有的权限接口多
	// 获取当前用户所拥有的权限接口

	// 前端传来最新的ApiID集合（这里采用先全部删除之前的，再全部添加最新的方法）

}

// 删除角色
func (rc RoleController) BatchDeleteRoleByIds(c *gin.Context) {
	var req vo.DeleteRoleRequest
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

	// 获取当前用户最高等级角色
	ur := repository.NewUserRepository()
	minSort, _, err := ur.GetCurrentUserMinRoleSort(c)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	// 前端传来需要删除的角色ID
	roleIds := req.RoleIds
	// 查询角色信息
	roles, err := rc.RoleRepository.GetRolesByIds(roleIds)
	if err != nil {
		response.Fail(c, nil, "查询角色信息失败: "+err.Error())
		return
	}
	if len(roles) == 0 {
		response.Fail(c, nil, "未查询到角色信息")
		return
	}

	// 不能删除比自己角色等级高或相等的角色
	for _, role := range roles {
		if minSort >= role.Sort {
			response.Fail(c, nil, "不能删除比自己角色等级高或相等的角色")
			return
		}
	}

	// 删除角色
	err = rc.RoleRepository.BatchDeleteRoleByIds(roleIds)
	if err != nil {
		response.Fail(c, nil, "删除角色失败")
		return
	}

	// 删除角色成功可以更新用户信息缓存，但是这里没有更新，casbin中间件会拦截非法操作

	response.Success(c, nil, "删除角色成功")

}
