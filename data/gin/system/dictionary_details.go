package data

import (
	"github.com/flipped-aurora/gva/global"
	model "github.com/flipped-aurora/gva/model/gva/system"
	"github.com/gookit/color"
	"gorm.io/gorm"
	"time"
)

var DictionaryDetail = new(dictionaryDetail)

type dictionaryDetail struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: dictionary_details 表数据初始化
func (d *dictionaryDetail) Init() error {
	details := []model.SysDictionaryDetail{
		{Model: global.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "smallint", Value: 1, Status: status, Sort: 1, SysDictionaryID: 2},
		{Model: global.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "mediumint", Value: 2, Status: status, Sort: 2, SysDictionaryID: 2},
		{Model: global.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "int", Value: 3, Status: status, Sort: 3, SysDictionaryID: 2},
		{Model: global.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "bigint", Value: 4, Status: status, Sort: 4, SysDictionaryID: 2},
		{Model: global.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "date", Status: status, SysDictionaryID: 3},
		{Model: global.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "time", Value: 1, Status: status, Sort: 1, SysDictionaryID: 3},
		{Model: global.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "year", Value: 2, Status: status, Sort: 2, SysDictionaryID: 3},
		{Model: global.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "datetime", Value: 3, Status: status, Sort: 3, SysDictionaryID: 3},
		{Model: global.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "timestamp", Value: 5, Status: status, Sort: 5, SysDictionaryID: 3},
		{Model: global.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "float", Status: status, SysDictionaryID: 4},
		{Model: global.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "double", Value: 1, Status: status, Sort: 1, SysDictionaryID: 4},
		{Model: global.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "decimal", Value: 2, Status: status, Sort: 2, SysDictionaryID: 4},
		{Model: global.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "char", Status: status, SysDictionaryID: 5},
		{Model: global.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "varchar", Value: 1, Status: status, Sort: 1, SysDictionaryID: 5},
		{Model: global.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "tinyblob", Value: 2, Status: status, Sort: 2, SysDictionaryID: 5},
		{Model: global.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "tinytext", Value: 3, Status: status, Sort: 3, SysDictionaryID: 5},
		{Model: global.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "text", Value: 4, Status: status, Sort: 4, SysDictionaryID: 5},
		{Model: global.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "blob", Value: 5, Status: status, Sort: 5, SysDictionaryID: 5},
		{Model: global.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "mediumblob", Value: 6, Status: status, Sort: 6, SysDictionaryID: 5},
		{Model: global.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "mediumtext", Value: 7, Status: status, Sort: 7, SysDictionaryID: 5},
		{Model: global.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "longblob", Value: 8, Status: status, Sort: 8, SysDictionaryID: 5},
		{Model: global.Model{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "longtext", Value: 9, Status: status, Sort: 9, SysDictionaryID: 5},
		{Model: global.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Label: "tinyint", Status: status, SysDictionaryID: 6},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 23}).Find(&[]model.SysDictionaryDetail{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_dictionary_details 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&details).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 定义表名
func (d *dictionaryDetail) TableName() string {
	return "sys_dictionary_details"
}
