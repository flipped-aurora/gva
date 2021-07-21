package model

import "github.com/flipped-aurora/gva/library/global"

type JwtBlacklist struct {
	global.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}

func (j *JwtBlacklist) TableName() string {
	return "system_jwt_blacklist"
}
