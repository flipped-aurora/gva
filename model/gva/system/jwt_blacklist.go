package model

import "github.com/flipped-aurora/gva/global"

type JwtBlacklist struct {
	global.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}
