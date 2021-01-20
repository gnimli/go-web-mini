package dto

import "go-lim/model"

type UserInfoDto struct {
	Username     string        `json:"username"`
	Mobile       string        `json:"mobile"`
	Avatar       string        `json:"avatar"`
	Nickname     string        `json:"nickname"`
	Introduction string        `json:"introduction"`
	Roles        []*model.Role `json:"roles"`
}

func ToUserInfoDto(user model.User) UserInfoDto {
	return UserInfoDto{
		Username:     user.Username,
		Mobile:       user.Mobile,
		Avatar:       user.Avatar,
		Nickname:     user.Nickname,
		Introduction: user.Introduction,
		Roles:        user.Roles,
	}
}
