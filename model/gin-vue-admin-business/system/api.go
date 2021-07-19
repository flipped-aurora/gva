package model

import (
	"errors"
	"github.com/flipped-aurora/gva/library/global"
	_errors "github.com/pkg/errors"
	"gorm.io/gorm"
)

type Api struct {
	global.Model
	Path        string `json:"path" gorm:"comment:api路径"`             // api路径
	Method      string `json:"method" gorm:"default:POST;comment:方法"` // api请求方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组"`          // api组
	Description string `json:"description" gorm:"comment:api中文描述"`    // api中文描述
}

func (a *Api) BeforeCreate(tx *gorm.DB) error {
	if errors.Is(tx.Where("path = ? AND method = ?", a.Path, a.Method).First(&Api{}).Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return _errors.New("存在相同api")
}

func (a *Api) AfterDelete(tx *gorm.DB) error {
	return tx.Model(&Casbin{}).Where("v1 = ? AND v2 = ?", a.Path, a.Method).Delete(&Casbin{}).Error
}

func (a *Api) TableName() string {
	return "system_apis"
}
