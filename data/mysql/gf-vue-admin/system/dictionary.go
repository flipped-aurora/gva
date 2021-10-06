package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var (
	status     = new(bool)
	Dictionary = new(dictionary)
)

type dictionary struct{}

func (d *dictionary) TableName() string {
	var entity system.Dictionary
	return entity.TableName()
}

func (d *dictionary) Initialize() error {
	*status = true
	entities := []system.Dictionary{
		{Name: "性别", Type: "sex", Status: status, Desc: "性别字典"},
		{Name: "数据库int类型", Type: "int", Status: status, Desc: "int类型对应的数据库类型"},
		{Name: "数据库时间日期类型", Type: "time.Time", Status: status, Desc: "数据库时间日期类型"},
		{Name: "数据库浮点型", Type: "float64", Status: status, Desc: "数据库浮点型"},
		{Name: "数据库字符串", Type: "string", Status: status, Desc: "数据库字符串"},
		{Name: "数据库bool类型", Type: "bool", Status: status, Desc: "数据库bool类型"},
	}
	if err := global.Db.Create(&entities).Error; err != nil { // 创建 model.Dictionary 初始化数据
		return err
	}
	return nil
}

func (d *dictionary) CheckDataExist() bool {
	if errors.Is(global.Db.Where("type = ?", "bool").First(&system.Dictionary{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
