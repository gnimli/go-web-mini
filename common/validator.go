package common

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ch_translations "github.com/go-playground/validator/v10/translations/zh"
)

// 全局Validate数据校验实列
var Validate *validator.Validate

// 全局翻译器
var Trans ut.Translator

// 初始化Validator数据校验
func InitValidate() {
	zh := zh.New()
	uni := ut.New(zh, zh)
	trans, _ := uni.GetTranslator("zh")
	Trans = trans
	Validate = validator.New()
	ch_translations.RegisterDefaultTranslations(Validate, Trans)
	Log.Infof("初始化validator.v10数据校验器完成")
}
