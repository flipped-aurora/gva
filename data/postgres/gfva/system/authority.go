//+build postgres

package data

import (
	"github.com/flipped-aurora/gva/library/global"
	model "github.com/flipped-aurora/gva/model/gfva/system"
	"time"

	"github.com/gookit/color"

	"gorm.io/gorm"
)

var Authority = new(authority)

type authority struct{}

// Init authorities 表数据初始化
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *authority) Init() error {
	authorities := []model.Authority{
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "888", AuthorityName: I18nHash["OrdinaryUser"], ParentId: "0", DefaultRouter: "dashboard"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "8881", AuthorityName: I18nHash["NormalUserSubRole"], ParentId: "888", DefaultRouter: "dashboard"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "9528", AuthorityName: I18nHash["TestRole"], ParentId: "0", DefaultRouter: "dashboard"},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ? ", []string{"888", "9528"}).Find(&[]model.Authority{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> authorities 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorities).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

// TableName 定义表名
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *authority) TableName() string {
	return "authorities"
}
