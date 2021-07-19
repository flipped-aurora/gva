package model

import "github.com/flipped-aurora/gva/library/global"

type DictionaryDetail struct {
	global.Model
	Sort         int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序标记"`
	Value        int    `json:"value" form:"value" gorm:"column:value;comment:字典值"`
	Label        string `json:"label" form:"label" gorm:"column:label;comment:展示值"`
	Status       *bool  `json:"status" form:"status" gorm:"column:status;comment:启用状态"`
	DictionaryID int    `json:"sysDictionaryID" gorm:"column:dictionary_id;comment:关联标记"`
}
func (d *DictionaryDetail) TableName() string {
	return "system_dictionary_details"
}