package datas

import (
	"github.com/flipped-aurora/gva/data/model"
	"gorm.io/gorm"
)

var DataAuthorityId = []model.SysDataAuthorityId{
	{"888", "888"},
	{"888", "8881"},
	{"888", "9528"},
	{"9528", "8881"},
	{"9528", "9528"},
}

func InitSysDataAuthorityId(db *gorm.DB) (err error) {
	return db.Transaction(func(tx *gorm.DB) error {
		if tx.Create(&DataAuthorityId).Error != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}
