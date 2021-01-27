package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
	"os"
)

// 系统配置，对应yml
// viper内置了mapstructure, yml文件用"-"区分单词, 转为驼峰方便

// 全局配置变量
var Conf = new(config)

type config struct {
	System    *SystemConfig    `mapstructure:"system" json:"system"`
	Logs      *LogsConfig      `mapstructure:"logs" json:"logs"`
	Mysql     *MysqlConfig     `mapstructure:"mysql" json:"mysql"`
	Casbin    *CasbinConfig    `mapstructure:"casbin" json:"casbin"`
	Jwt       *JwtConfig       `mapstructure:"jwt" json:"jwt"`
	RateLimit *RateLimitConfig `mapstructure:"rate-limit" json:"rateLimit"`
}

// 设置读取配置信息
func InitConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("读取应用目录失败:%s \n", err))
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "./")
	// 读取配置信息
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败:%s \n", err))
	}
	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("初始化配置文件失败:%s \n", err))
	}
}

type SystemConfig struct {
	Mode          string `mapstructure:"mode" json:"mode"`
	UrlPathPrefix string `mapstructure:"url-path-prefix" json:"urlPathPrefix"`
	Port          int    `mapstructure:"port" json:"port"`
	InitData      bool   `mapstructure:"init-data" json:"initData"`
}

type LogsConfig struct {
	Level      zapcore.Level `mapstructure:"level" json:"level"`
	Path       string        `mapstructure:"path" json:"path"`
	MaxSize    int           `mapstructure:"max-size" json:"maxSize"`
	MaxBackups int           `mapstructure:"max-backups" json:"maxBackups"`
	MaxAge     int           `mapstructure:"max-age" json:"maxAge"`
	Compress   bool          `mapstructure:"compress" json:"compress"`
}

type MysqlConfig struct {
	Username    string `mapstructure:"username" json:"username"`
	Password    string `mapstructure:"password" json:"password"`
	Database    string `mapstructure:"database" json:"database"`
	Host        string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	Query       string `mapstructure:"query" json:"query"`
	LogMode     bool   `mapstructure:"log-mode" json:"logMode"`
	TablePrefix string `mapstructure:"table-prefix" json:"tablePrefix"`
	Charset     string `mapstructure:"charset" json:"charset"`
	Collation   string `mapstructure:"collation" json:"collation"`
}

type CasbinConfig struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath"`
}

type JwtConfig struct {
	Realm      string `mapstructure:"realm" json:"realm"`
	Key        string `mapstructure:"key" json:"key"`
	Timeout    int    `mapstructure:"timeout" json:"timeout"`
	MaxRefresh int    `mapstructure:"max-refresh" json:"maxRefresh"`
}

type RateLimitConfig struct {
	FillInterval int64 `mapstructure:"fill-interval" json:"fillInterval"`
	Capacity     int64 `mapstructure:"capacity" json:"capacity"`
}
