package common

import (
	"errors"
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
		err := DB.Debug().Create(&newRoles).Error
		if err != nil {
			Log.Errorf("写入系统角色失败：%v", err)
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
		err := DB.Debug().Create(&newMenus).Error
		if err != nil {
			Log.Errorf("写入系统菜单失败：%v", err)
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
		err := DB.Debug().Create(&newUsers).Error
		if err != nil {
			Log.Errorf("写入用户失败：%v", err)
		}
	}

	// 4.写入api

}
