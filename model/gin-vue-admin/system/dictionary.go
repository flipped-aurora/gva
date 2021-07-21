package model

import (
	"errors"
	"github.com/flipped-aurora/gva/library/global"
	_errors "github.com/pkg/errors"
	"gorm.io/gorm"
)

type Dictionary struct {
	global.Model
	Name              string             `json:"name" form:"name" gorm:"column:name;comment:字典名（中）"`
	Type              string             `json:"type" form:"type" gorm:"column:type;comment:字典名（英）"`
	Desc              string             `json:"desc" form:"desc" gorm:"column:desc;comment:描述"`
	Status            *bool              `json:"status" form:"status" gorm:"column:status;comment:状态"`
	DictionaryDetails []DictionaryDetail `json:"sysDictionaryDetails" form:"sysDictionaryID" gorm:"foreignKey:DictionaryID;references:ID"`
}

func (d *Dictionary) TableName() string {
	return "system_dictionaries"
}

func (d *Dictionary) BeforeCreate(tx *gorm.DB) error {
	if !errors.Is(tx.First(&Dictionary{}, "type = ? ", d.Type).Error, gorm.ErrRecordNotFound) {
		return _errors.Wrap(gorm.ErrRecordNotFound, "已存在数据,不允许创建!")
	}
	return nil
}
