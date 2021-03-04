package model

import "github.com/flipped-aurora/tool/cmd/gva/internal/global"

type JwtBlacklist struct {
	global.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}
