package model

import (
	"github.com/flipped-aurora/gva/library/global"
	model "github.com/flipped-aurora/gva/model/gin-vue-admin-business/system"
	_errors "github.com/pkg/errors"
	"gorm.io/gorm"
)

var (
	status     = new(bool)
	Dictionary = new(dictionary)
)

type dictionary struct{}

func (d *dictionary) TableName() string {
	var entity model.Dictionary
	return entity.TableName()
}

func (d *dictionary) Initialize() error {
	*status = true
	entities := []model.Dictionary{
		{Model: global.Model{ID: 1}, Name: "性别", Type: "sex", Status: status, Desc: "性别字典"},
		{Model: global.Model{ID: 2}, Name: "数据库int类型", Type: "int", Status: status, Desc: "int类型对应的数据库类型"},
		{Model: global.Model{ID: 3}, Name: "数据库时间日期类型", Type: "time.Time", Status: status, Desc: "数据库时间日期类型"},
		{Model: global.Model{ID: 4}, Name: "数据库浮点型", Type: "float64", Status: status, Desc: "数据库浮点型"},
		{Model: global.Model{ID: 5}, Name: "数据库字符串", Type: "string", Status: status, Desc: "数据库字符串"},
		{Model: global.Model{ID: 6}, Name: "数据库bool类型", Type: "bool", Status: status, Desc: "数据库bool类型"},
	}
	if err := global.Db.Create(&entities).Error; err != nil { // 创建 model.Dictionary 初始化数据
		return err
	}
	return nil
}

func (d *dictionary) CheckDataExist() bool {
	if _errors.Is(global.Db.Where("id = ?", 6).First(&model.Dictionary{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
