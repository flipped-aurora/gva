//+build postgres

package data

import (
	"github.com/flipped-aurora/gva/library/global"
	model "github.com/flipped-aurora/gva/model/gfva/system"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var ResourcesAuthorities = new(resources)

type resources struct{}

// Init data_authorities 表数据初始化
// Author: [SliverHorn](https://github.com/SliverHorn)
func (d *resources) Init() error {
	infos := []model.DataAuthorities{
		{AuthorityId: "888", DataAuthority: "888"},
		{AuthorityId: "888", DataAuthority: "8881"},
		{AuthorityId: "888", DataAuthority: "9528"},
		{AuthorityId: "9528", DataAuthority: "8881"},
		{AuthorityId: "9528", DataAuthority: "9528"},
	}
	return global.Db.Table("data_authorities").Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ('888', '9528')").Find(&[]model.DataAuthorities{}).RowsAffected == 5 {
			color.Danger.Printf("\n[postgres] --> %v 表初始数据已存在!", d.TableName())
			return nil
		}
		if err := tx.Create(&infos).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Printf("\n[postgres] --> %v 表初始数据成功!", d.TableName())
		return nil
	})
}

// TableName 自定义表名
// Author: [SliverHorn](https://github.com/SliverHorn)
func (d *resources) TableName() string {
	return "data_authorities"
}
