package common

import (
	"errors"
	"github.com/thoas/go-funk"
	"go-lim/config"
	"go-lim/model"
	"go-lim/util"
	"gorm.io/gorm"
)

// 初始化mysql数据
func InitData() {
	// 是否初始化数据
	if !config.Conf.System.InitData {
		return
	}

	// 1.写入角色数据
	newRoles := make([]*model.Role, 0)
	roles := []*model.Role{
		{
			Model:   gorm.Model{ID: 1},
			Name:    "管理员",
			Keyword: "admin",
			Desc:    "管理员",
			Sort:    0,
			Status:  1,
			Creator: "系统",
		},
		{
			Model:   gorm.Model{ID: 2},
			Name:    "普通用户",
			Keyword: "user",
			Desc:    "有管理权限的用户",
			Sort:    3,
			Status:  1,
			Creator: "系统",
		},
		{
			Model:   gorm.Model{ID: 3},
			Name:    "访客",
			Keyword: "guest",
			Desc:    "没有管理权限的用户",
			Sort:    5,
			Status:  1,
			Creator: "系统",
		},
	}

	for _, role := range roles {
		err := DB.First(&role, role.ID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newRoles = append(newRoles, role)
		}
	}

	if len(newRoles) > 0 {
		err := DB.Create(&newRoles).Error
		if err != nil {
			Log.Errorf("写入系统角色数据失败：%v", err)
		}
	}

	// 2写入菜单
	newMenus := make([]model.Menu, 0)
	menus := []model.Menu{
		{
			Model:     gorm.Model{ID: 1},
			Path:      "/",
			Component: "Layout",
			Redirect:  "/dashboard",
			Sort:      0,
			ParentId:  0,
			Roles:     roles[:],
		},
		{
			Model:     gorm.Model{ID: 2},
			Name:      "Dashboard",
			Title:     "首页",
			Icon:      "dashboard",
			Path:      "dashboard",
			Component: "/dashboard/index",
			Sort:      1,
			ParentId:  1,
			Roles:     roles[:],
		},
		{
			Model:     gorm.Model{ID: 3},
			Name:      "System",
			Title:     "系统管理",
			Icon:      "system",
			Path:      "/system",
			Component: "Layout",
			Redirect:  "/system/user",
			Sort:      10,
			ParentId:  0,
			Roles:     roles[:2],
		},
		{
			Model:     gorm.Model{ID: 4},
			Name:      "User",
			Title:     "用户管理",
			Icon:      "user",
			Path:      "user",
			Component: "/system/user/index",
			Sort:      11,
			ParentId:  3,
			Roles:     roles[:2],
		},
		{
			Model:     gorm.Model{ID: 5},
			Name:      "Role",
			Title:     "角色管理",
			Icon:      "role",
			Path:      "role",
			Component: "/system/role/index",
			Sort:      12,
			ParentId:  3,
			Roles:     roles[:2],
		},
		{
			Model:     gorm.Model{ID: 6},
			Name:      "Menu",
			Title:     "菜单管理",
			Icon:      "menu",
			Path:      "menu",
			Component: "/system/menu/index",
			Sort:      13,
			ParentId:  3,
			Roles:     roles[:2],
		},
		{
			Model:     gorm.Model{ID: 7},
			Name:      "Api",
			Title:     "接口管理",
			Icon:      "api",
			Path:      "api",
			Component: "/system/api/index",
			Sort:      14,
			ParentId:  3,
			Roles:     roles[:2],
		},
	}
	for _, menu := range menus {
		err := DB.First(&menu, menu.ID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newMenus = append(newMenus, menu)
		}
	}
	if len(newMenus) > 0 {
		err := DB.Create(&newMenus).Error
		if err != nil {
			Log.Errorf("写入系统菜单数据失败：%v", err)
		}
	}

	// 3.写入用户
	newUsers := make([]model.User, 0)
	users := []model.User{
		{
			Model:        gorm.Model{ID: 1},
			Username:     "admin",
			Password:     util.GenPasswd("123456"),
			Mobile:       "18888888888",
			Avatar:       "",
			Nickname:     "管理员",
			Introduction: "我是系统的管理员",
			Status:       1,
			Creator:      "",
			Roles:        roles[:1],
		},
		{
			Model:        gorm.Model{ID: 2},
			Username:     "lim",
			Password:     util.GenPasswd("123456"),
			Mobile:       "19999999999",
			Avatar:       "",
			Nickname:     "明哥",
			Introduction: "哈哈哈哈哈",
			Status:       1,
			Creator:      "",
			Roles:        roles[:2],
		},
		{
			Model:        gorm.Model{ID: 3},
			Username:     "nike",
			Password:     util.GenPasswd("123456"),
			Mobile:       "13333333333",
			Avatar:       "",
			Nickname:     "little nike",
			Introduction: "haha",
			Status:       1,
			Creator:      "",
			Roles:        roles[2:3],
		},
	}

	for _, user := range users {
		err := DB.First(&user, user.ID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newUsers = append(newUsers, user)
		}
	}

	if len(newUsers) > 0 {
		err := DB.Create(&newUsers).Error
		if err != nil {
			Log.Errorf("写入用户数据失败：%v", err)
		}
	}

	// 4.写入api
	apis := []model.Api{
		{
			Model:    gorm.Model{ID: 1},
			Method:   "POST",
			Path:     "/base/login",
			Category: "base",
			Desc:     "用户登录",
		},
		{
			Model:    gorm.Model{ID: 2},
			Method:   "POST",
			Path:     "/base/logout",
			Category: "base",
			Desc:     "用户登出",
		},
		{
			Model:    gorm.Model{ID: 3},
			Method:   "POST",
			Path:     "/base/refreshToken",
			Category: "base",
			Desc:     "刷新JWT令牌",
		},
		{
			Model:    gorm.Model{ID: 4},
			Method:   "POST",
			Path:     "/user/info",
			Category: "user",
			Desc:     "获取当前登录用户信息",
		},
		{
			Model:    gorm.Model{ID: 5},
			Method:   "GET",
			Path:     "/user/list",
			Category: "user",
			Desc:     "获取用户列表",
		},
		{
			Model:    gorm.Model{ID: 6},
			Method:   "PUT",
			Path:     "/user/changePwd",
			Category: "user",
			Desc:     "修改用户登录密码",
		},
		{
			Model:    gorm.Model{ID: 7},
			Method:   "POST",
			Path:     "/user/create",
			Category: "user",
			Desc:     "创建用户",
		},
		{
			Model:    gorm.Model{ID: 8},
			Method:   "PATCH",
			Path:     "/user/update/:userId",
			Category: "user",
			Desc:     "更新用户",
		},
		{
			Model:    gorm.Model{ID: 9},
			Method:   "DELETE",
			Path:     "/user/delete/batch",
			Category: "user",
			Desc:     "批量删除用户",
		},
		{
			Model:    gorm.Model{ID: 10},
			Method:   "GET",
			Path:     "/menu/tree",
			Category: "menu",
			Desc:     "获取权限菜单",
		},
		{
			Model:    gorm.Model{ID: 11},
			Method:   "GET",
			Path:     "/menu/list",
			Category: "menu",
			Desc:     "获取菜单列表",
		},
		{
			Model:    gorm.Model{ID: 12},
			Method:   "POST",
			Path:     "/menu/create",
			Category: "menu",
			Desc:     "创建菜单",
		},
		{
			Model:    gorm.Model{ID: 13},
			Method:   "PATCH",
			Path:     "/menu/update/:menuId",
			Category: "menu",
			Desc:     "更新菜单",
		},
		{
			Model:    gorm.Model{ID: 14},
			Method:   "DELETE",
			Path:     "/menu/delete/batch",
			Category: "menu",
			Desc:     "批量删除菜单",
		},
		{
			Model:    gorm.Model{ID: 15},
			Method:   "GET",
			Path:     "/role/list",
			Category: "role",
			Desc:     "获取角色列表",
		},
		{
			Model:    gorm.Model{ID: 16},
			Method:   "POST",
			Path:     "/role/create",
			Category: "role",
			Desc:     "创建角色",
		},
		{
			Model:    gorm.Model{ID: 17},
			Method:   "PATCH",
			Path:     "/role/update/:roleId",
			Category: "role",
			Desc:     "更新角色",
		},
		{
			Model:    gorm.Model{ID: 18},
			Method:   "DELETE",
			Path:     "/role/delete/batch",
			Category: "role",
			Desc:     "批量删除角色",
		},
		{
			Model:    gorm.Model{ID: 19},
			Method:   "GET",
			Path:     "/api/list",
			Category: "api",
			Desc:     "获取接口列表",
		},
		{
			Model:    gorm.Model{ID: 20},
			Method:   "POST",
			Path:     "/api/create",
			Category: "api",
			Desc:     "创建接口",
		},
		{
			Model:    gorm.Model{ID: 21},
			Method:   "PATCH",
			Path:     "/api/update/:roleId",
			Category: "api",
			Desc:     "更新接口",
		},
		{
			Model:    gorm.Model{ID: 22},
			Method:   "DELETE",
			Path:     "/api/delete/batch",
			Category: "api",
			Desc:     "批量删除接口",
		},
		{
			Model:    gorm.Model{ID: 23},
			Method:   "GET",
			Path:     "/menu/all/:roleId",
			Category: "menu",
			Desc:     "查询指定角色的菜单树",
		},
		{
			Model:    gorm.Model{ID: 24},
			Method:   "GET",
			Path:     "/api/all/category/:roleId",
			Category: "api",
			Desc:     "查询指定角色的接口(以分类分组)",
		},
		{
			Model:    gorm.Model{ID: 25},
			Method:   "PATCH",
			Path:     "/role/menus/update/:roleId",
			Category: "role",
			Desc:     "更新角色的权限菜单",
		},
		{
			Model:    gorm.Model{ID: 26},
			Method:   "PATCH",
			Path:     "/role/apis/update/:roleId",
			Category: "role",
			Desc:     "更新角色的权限接口",
		},
	}
	newApi := make([]model.Api, 0)
	newRoleCasbin := make([]model.RoleCasbin, 0)
	for _, api := range apis {
		err := DB.First(&api, api.ID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newApi = append(newApi, api)

			// 管理员拥有所有API权限
			newRoleCasbin = append(newRoleCasbin, model.RoleCasbin{
				Keyword: roles[0].Keyword,
				Path:    api.Path,
				Method:  api.Method,
			})

			// 非管理员拥有基础权限
			basePaths := []string{
				"/base/login",
				"/base/logout",
				"/base/refreshToken",
				"/user/info",
				"/menu/tree",
			}

			if funk.ContainsString(basePaths, api.Path) {
				newRoleCasbin = append(newRoleCasbin, model.RoleCasbin{
					Keyword: roles[1].Keyword,
					Path:    api.Path,
					Method:  api.Method,
				})
			}
		}
	}

	if len(newApi) > 0 {
		if err := DB.Create(&newApi).Error; err != nil {
			Log.Errorf("写入api数据失败：%v", err)
		}
	}

	if len(newRoleCasbin) > 0 {
		rules := make([][]string, 0)
		for _, c := range newRoleCasbin {
			rules = append(rules, []string{
				c.Keyword, c.Path, c.Method,
			})
		}
		isAdd, err := CasbinEnforcer.AddPolicies(rules)
		if !isAdd {
			Log.Errorf("写入casbin数据失败：%v", err)
		}
	}
}
