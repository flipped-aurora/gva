package data

import (
	"github.com/flipped-aurora/gva/global"
	model "github.com/flipped-aurora/gva/model/gin/system"
	"github.com/gookit/color"
	"gorm.io/gorm"
	"time"
)

var Authority = new(authority)

type authority struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_authorities 表数据初始化
func (a *authority) Init() error {
	var authorities = []model.SysAuthority{
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "888", AuthorityName: global.I18n.T("OrdinaryUser"), ParentId: "0", DefaultRouter: "dashboard"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "8881", AuthorityName: global.I18n.T("NormalUserSubRole"), ParentId: "888", DefaultRouter: "dashboard"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "9528", AuthorityName: global.I18n.T("TestRole"), ParentId: "0", DefaultRouter: "dashboard"},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ? ", []string{"888", "9528"}).Find(&[]model.SysAuthority{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_authorities 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorities).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] -->  表初始数据成功!")
		return nil
	})
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 定义表名
func (a *authority) TableName() string {
	return "sys_authorities"
}
