package data

import (
	"github.com/flipped-aurora/gva/global"
	model "github.com/flipped-aurora/gva/model/gfva/system"
	"time"

	"github.com/gookit/color"
	"gorm.io/gorm"
)

var (
	_true      *bool
	_false     *bool
	Dictionary = new(dictionary)
)

func init() {
	_true = new(bool)
	*_true = true
	_false = new(bool)
	*_false = false
}

type dictionary struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: dictionaries 表数据初始化
func (d *dictionary) Init() error {
	dictionaries := []model.Dictionary{
		{Model: global.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: global.I18n.T("Sex"), Type: "sex", Status: _true, Desc: global.I18n.T("SexDictionary")},
		{Model: global.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: global.I18n.T("DBTypeInt"), Type: "int", Status: _true, Desc: global.I18n.T("DBTypeInt")},
		{Model: global.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: global.I18n.T("DBTypeDateTime"), Type: "time.Time", Status: _true, Desc: global.I18n.T("DBTypeDateTime")},
		{Model: global.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: global.I18n.T("DBTypeFloat"), Type: "float64", Status: _true, Desc: global.I18n.T("DBTypeFloat")},
		{Model: global.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: global.I18n.T("DBTypeString"), Type: "string", Status: _true, Desc: global.I18n.T("DBTypeString")},
		{Model: global.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: global.I18n.T("DBTypeBool"), Type: "bool", Status: _true, Desc: global.I18n.T("DBTypeBool")},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 6}).Find(&[]model.Dictionary{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> dictionaries 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&dictionaries).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 定义表名
func (d *dictionary) TableName() string {
	return "dictionaries"
}
