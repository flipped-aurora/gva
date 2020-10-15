package datas

import (
	"time"

	"github.com/flipped-aurora/gva/data/model"

	"gorm.io/gorm"
)

func InitExaFileUploadAndDownload(db *gorm.DB) (err error) {
	tx := db.Begin() // 开始事务
	insert := []model.ExaFileUploadAndDownload{
		{gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "10.png", "http://qmplusimg.henrongyi.top/gvalogo.png", "png", "158787308910.png"},
		{gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "logo.png", "http://qmplusimg.henrongyi.top/1576554439myAvatar.png", "png", "1587973709logo.png"},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}
