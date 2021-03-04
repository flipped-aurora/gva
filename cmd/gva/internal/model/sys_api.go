package model

import (
	"github.com/flipped-aurora/tool/cmd/gva/internal/global"
)

type SysApi struct {
	global.Model
	Path        string `json:"path" gorm:"comment:api路径"`
	Description string `json:"description" gorm:"comment:api中文描述"`
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组"`
	Method      string `json:"method" gorm:"default:POST" gorm:"comment:方法"`
}
