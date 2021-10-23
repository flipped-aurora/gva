package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var DictionaryDetail = new(dictionaryDetail)

type dictionaryDetail struct{}

func (d *dictionaryDetail) TableName() string {
	var entity system.DictionaryDetail
	return entity.TableName()
}

func (d *dictionaryDetail) Initialize() error {
	entities := []system.DictionaryDetail{
		{Model: global.Model{ID: 1}, Label: "smallint", Value: 1, Status: status, Sort: 1, DictionaryID: 2},
		{Model: global.Model{ID: 2}, Label: "mediumint", Value: 2, Status: status, Sort: 2, DictionaryID: 2},
		{Model: global.Model{ID: 3}, Label: "int", Value: 3, Status: status, Sort: 3, DictionaryID: 2},
		{Model: global.Model{ID: 4}, Label: "bigint", Value: 4, Status: status, Sort: 4, DictionaryID: 2},
		{Model: global.Model{ID: 5}, Label: "date", Status: status, DictionaryID: 3},
		{Model: global.Model{ID: 6}, Label: "time", Value: 1, Status: status, Sort: 1, DictionaryID: 3},
		{Model: global.Model{ID: 7}, Label: "year", Value: 2, Status: status, Sort: 2, DictionaryID: 3},
		{Model: global.Model{ID: 8}, Label: "datetime", Value: 3, Status: status, Sort: 3, DictionaryID: 3},
		{Model: global.Model{ID: 9}, Label: "timestamp", Value: 5, Status: status, Sort: 5, DictionaryID: 3},
		{Model: global.Model{ID: 10}, Label: "float", Status: status, DictionaryID: 4},
		{Model: global.Model{ID: 11}, Label: "double", Value: 1, Status: status, Sort: 1, DictionaryID: 4},
		{Model: global.Model{ID: 12}, Label: "decimal", Value: 2, Status: status, Sort: 2, DictionaryID: 4},
		{Model: global.Model{ID: 13}, Label: "char", Status: status, DictionaryID: 5},
		{Model: global.Model{ID: 14}, Label: "varchar", Value: 1, Status: status, Sort: 1, DictionaryID: 5},
		{Model: global.Model{ID: 15}, Label: "tinyblob", Value: 2, Status: status, Sort: 2, DictionaryID: 5},
		{Model: global.Model{ID: 16}, Label: "tinytext", Value: 3, Status: status, Sort: 3, DictionaryID: 5},
		{Model: global.Model{ID: 17}, Label: "text", Value: 4, Status: status, Sort: 4, DictionaryID: 5},
		{Model: global.Model{ID: 18}, Label: "blob", Value: 5, Status: status, Sort: 5, DictionaryID: 5},
		{Model: global.Model{ID: 19}, Label: "mediumblob", Value: 6, Status: status, Sort: 6, DictionaryID: 5},
		{Model: global.Model{ID: 20}, Label: "mediumtext", Value: 7, Status: status, Sort: 7, DictionaryID: 5},
		{Model: global.Model{ID: 21}, Label: "longblob", Value: 8, Status: status, Sort: 8, DictionaryID: 5},
		{Model: global.Model{ID: 22}, Label: "longtext", Value: 9, Status: status, Sort: 9, DictionaryID: 5},
		{Model: global.Model{ID: 23}, Label: "tinyint", Status: status, DictionaryID: 6},
	}
	if err := global.Db.Create(&entities).Error; err != nil { // 创建 model.Dictionary 初始化数据
		return err
	}
	return nil
}

func (d *dictionaryDetail) CheckDataExist() bool {
	if errors.Is(global.Db.Where("id = ?", 23).First(&system.DictionaryDetail{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
