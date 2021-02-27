package common

import (
	"errors"
	"github.com/thoas/go-funk"
	"go-web-mini/config"
	"go-web-mini/model"
	"go-web-mini/util"
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
			Desc:    new(string),
			Sort:    1,
			Status:  1,
			Creator: "系统",
		},
		{
			Model:   gorm.Model{ID: 2},
			Name:    "普通用户",
			Keyword: "user",
			Desc:    new(string),
			Sort:    3,
			Status:  1,
			Creator: "系统",
		},
		{
			Model:   gorm.Model{ID: 3},
			Name:    "访客",
			Keyword: "guest",
			Desc:    new(string),
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
	var uint0 uint = 0
	var uint1 uint = 1
	componentStr := "component"
	systemUserStr := "/system/user"
	userStr := "user"
	peoplesStr := "peoples"
	treeTableStr := "tree-table"
	treeStr := "tree"
	exampleStr := "example"
	logOperationStr := "/log/operation-log"
	documentationStr := "documentation"
	var uint6 uint = 6
	menus := []model.Menu{
		{
			Model:     gorm.Model{ID: 1},
			Name:      "System",
			Title:     "系统管理",
			Icon:      &componentStr,
			Path:      "/system",
			Component: "Layout",
			Redirect:  &systemUserStr,
			Sort:      10,
			ParentId:  &uint0,
			Roles:     roles[:1],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 2},
			Name:      "User",
			Title:     "用户管理",
			Icon:      &userStr,
			Path:      "user",
			Component: "/system/user/index",
			Sort:      11,
			ParentId:  &uint1,
			Roles:     roles[:1],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 3},
			Name:      "Role",
			Title:     "角色管理",
			Icon:      &peoplesStr,
			Path:      "role",
			Component: "/system/role/index",
			Sort:      12,
			ParentId:  &uint1,
			Roles:     roles[:1],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 4},
			Name:      "Menu",
			Title:     "菜单管理",
			Icon:      &treeTableStr,
			Path:      "menu",
			Component: "/system/menu/index",
			Sort:      13,
			ParentId:  &uint1,
			Roles:     roles[:1],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 5},
			Name:      "Api",
			Title:     "接口管理",
			Icon:      &treeStr,
			Path:      "api",
			Component: "/system/api/index",
			Sort:      14,
			ParentId:  &uint1,
			Roles:     roles[:1],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 6},
			Name:      "Log",
			Title:     "日志管理",
			Icon:      &exampleStr,
			Path:      "/log",
			Component: "Layout",
			Redirect:  &logOperationStr,
			Sort:      20,
			ParentId:  &uint0,
			Roles:     roles[:2],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 7},
			Name:      "OperationLog",
			Title:     "操作日志",
			Icon:      &documentationStr,
			Path:      "operation-log",
			Component: "/log/operation-log/index",
			Sort:      21,
			ParentId:  &uint6,
			Roles:     roles[:2],
			Creator:   "系统",
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
			Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
			Nickname:     new(string),
			Introduction: new(string),
			Status:       1,
			Creator:      "系统",
			Roles:        roles[:1],
		},
		{
			Model:        gorm.Model{ID: 2},
			Username:     "faker",
			Password:     util.GenPasswd("123456"),
			Mobile:       "19999999999",
			Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
			Nickname:     new(string),
			Introduction: new(string),
			Status:       1,
			Creator:      "系统",
			Roles:        roles[:2],
		},
		{
			Model:        gorm.Model{ID: 3},
			Username:     "nike",
			Password:     util.GenPasswd("123456"),
			Mobile:       "13333333333",
			Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
			Nickname:     new(string),
			Introduction: new(string),
			Status:       1,
			Creator:      "系统",
			Roles:        roles[1:2],
		},
		{
			Model:        gorm.Model{ID: 4},
			Username:     "bob",
			Password:     util.GenPasswd("123456"),
			Mobile:       "15555555555",
			Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
			Nickname:     new(string),
			Introduction: new(string),
			Status:       1,
			Creator:      "系统",
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
			Method:   "POST",
			Path:     "/base/login",
			Category: "base",
			Desc:     "用户登录",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/base/logout",
			Category: "base",
			Desc:     "用户登出",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/base/refreshToken",
			Category: "base",
			Desc:     "刷新JWT令牌",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/user/info",
			Category: "user",
			Desc:     "获取当前登录用户信息",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/user/list",
			Category: "user",
			Desc:     "获取用户列表",
			Creator:  "系统",
		},
		{
			Method:   "PUT",
			Path:     "/user/changePwd",
			Category: "user",
			Desc:     "更新用户登录密码",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/user/create",
			Category: "user",
			Desc:     "创建用户",
			Creator:  "系统",
		},
		{
			Method:   "PATCH",
			Path:     "/user/update/:userId",
			Category: "user",
			Desc:     "更新用户",
			Creator:  "系统",
		},
		{
			Method:   "DELETE",
			Path:     "/user/delete/batch",
			Category: "user",
			Desc:     "批量删除用户",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/role/list",
			Category: "role",
			Desc:     "获取角色列表",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/role/create",
			Category: "role",
			Desc:     "创建角色",
			Creator:  "系统",
		},
		{
			Method:   "PATCH",
			Path:     "/role/update/:roleId",
			Category: "role",
			Desc:     "更新角色",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/role/menus/get/:roleId",
			Category: "role",
			Desc:     "获取角色的权限菜单",
			Creator:  "系统",
		},
		{
			Method:   "PATCH",
			Path:     "/role/menus/update/:roleId",
			Category: "role",
			Desc:     "更新角色的权限菜单",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/role/apis/get/:roleId",
			Category: "role",
			Desc:     "获取角色的权限接口",
			Creator:  "系统",
		},
		{
			Method:   "PATCH",
			Path:     "/role/apis/update/:roleId",
			Category: "role",
			Desc:     "更新角色的权限接口",
			Creator:  "系统",
		},
		{
			Method:   "DELETE",
			Path:     "/role/delete/batch",
			Category: "role",
			Desc:     "批量删除角色",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/menu/list",
			Category: "menu",
			Desc:     "获取菜单列表",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/menu/tree",
			Category: "menu",
			Desc:     "获取菜单树",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/menu/create",
			Category: "menu",
			Desc:     "创建菜单",
			Creator:  "系统",
		},
		{
			Method:   "PATCH",
			Path:     "/menu/update/:menuId",
			Category: "menu",
			Desc:     "更新菜单",
			Creator:  "系统",
		},
		{
			Method:   "DELETE",
			Path:     "/menu/delete/batch",
			Category: "menu",
			Desc:     "批量删除菜单",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/menu/access/list/:userId",
			Category: "menu",
			Desc:     "获取用户的可访问菜单列表",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/menu/access/tree/:userId",
			Category: "menu",
			Desc:     "获取用户的可访问菜单树",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/api/list",
			Category: "api",
			Desc:     "获取接口列表",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/api/tree",
			Category: "api",
			Desc:     "获取接口树",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/api/create",
			Category: "api",
			Desc:     "创建接口",
			Creator:  "系统",
		},
		{
			Method:   "PATCH",
			Path:     "/api/update/:roleId",
			Category: "api",
			Desc:     "更新接口",
			Creator:  "系统",
		},
		{
			Method:   "DELETE",
			Path:     "/api/delete/batch",
			Category: "api",
			Desc:     "批量删除接口",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/log/operation/list",
			Category: "log",
			Desc:     "获取操作日志列表",
			Creator:  "系统",
		},
		{
			Method:   "DELETE",
			Path:     "/log/operation/delete/batch",
			Category: "log",
			Desc:     "批量删除操作日志",
			Creator:  "系统",
		},
	}
	newApi := make([]model.Api, 0)
	newRoleCasbin := make([]model.RoleCasbin, 0)
	for i, api := range apis {
		api.ID = uint(i + 1)
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
				"/menu/access/tree/:userId",
			}

			if funk.ContainsString(basePaths, api.Path) {
				newRoleCasbin = append(newRoleCasbin, model.RoleCasbin{
					Keyword: roles[1].Keyword,
					Path:    api.Path,
					Method:  api.Method,
				})
				newRoleCasbin = append(newRoleCasbin, model.RoleCasbin{
					Keyword: roles[2].Keyword,
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
