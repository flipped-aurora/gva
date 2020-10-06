package postgresql

import (
	"time"

	"github.com/flipped-aurora/gva/data/model"

	"gorm.io/gorm"
)

func InitSysDictionaryToPostgresql(db *gorm.DB) (err error) {
	status := new(bool)
	*status = true
	tx := db.Begin() // 开始事务
	insert := []model.SysDictionaryToPostgresql{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "性别", Type: "sex", Status: status, Description: "性别字典"},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库int类型", Type: "int", Status: status, Description: "int类型对应的数据库类型"},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库时间日期类型", Type: "time.Time", Status: status, Description: "数据库时间日期类型"},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库浮点型", Type: "float64", Status: status, Description: "数据库浮点型"},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库字符串", Type: "string", Status: status, Description: "数据库字符串"},
		{Model: gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库bool类型", Type: "bool", Status: status, Description: "数据库bool类型"},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}
