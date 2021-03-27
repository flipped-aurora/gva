package data

import (
	"github.com/flipped-aurora/gva/global"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var ResourcesAuthorities = new(resourcesAuthorities)

type resourcesAuthorities struct{}

type ResourcesAuthority struct {
	AuthorityId   string `gorm:"column:sys_authority_authority_id"`
	DataAuthority string `gorm:"column:data_authority_id_authority_id"`
}

var infos = []ResourcesAuthority{
	{AuthorityId: "888", DataAuthority: "888"},
	{AuthorityId: "888", DataAuthority: "8881"},
	{AuthorityId: "888", DataAuthority: "9528"},
	{AuthorityId: "9528", DataAuthority: "8881"},
	{AuthorityId: "9528", DataAuthority: "9528"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_data_authority_id 表数据初始化
func (d *resourcesAuthorities) Init() error {
	return global.Db.Table("sys_data_authority_id").Transaction(func(tx *gorm.DB) error {
		if tx.Where("sys_authority_authority_id IN ('888', '9528') ").Find(&[]ResourcesAuthority{}).RowsAffected == 5 {
			color.Danger.Println("\n[Mysql] --> sys_data_authority_id 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&infos).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 定义表名
func (d resourcesAuthorities) TableName() string {
	return "sys_data_authority_id"
}
