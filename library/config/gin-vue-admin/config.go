package config

import config "github.com/flipped-aurora/gva/library/config/gin-vue-admin/gorm"

type Config struct {
	JWT       JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap       Zap       `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis     Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email     Email     `mapstructure:"email" json:"email" yaml:"email"`
	Casbin    Casbin    `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System    System    `mapstructure:"system" json:"system" yaml:"system"`
	Captcha   Captcha   `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Uploader  Uploader  `mapstructure:"uploader" json:"uploader" yaml:"uploader"`
	AutoCode  AutoCode  `mapstructure:"auto-code" json:"autoCode" yaml:"auto-code"`
	Snowflake Snowflake `mapstructure:"snowflake" json:"snowflake" yaml:"snowflake"`

	Gorm config.Gorm `mapstructure:"gorm" json:"gorm" yaml:"gorm"`

	Local   Local   `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu   Qiniu   `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	Aliyun  Aliyun  `mapstructure:"aliyun" json:"aliyun" yaml:"aliyun"`
	Tencent Tencent `mapstructure:"tencent" json:"tencent" yaml:"tencent"`
}
