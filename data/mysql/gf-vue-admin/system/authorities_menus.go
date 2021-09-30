package system

import (
	"fmt"
	system "github.com/flipped-aurora/gf-vue-admin/app/model"
	"github.com/flipped-aurora/gva/library/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var AuthoritiesMenus = new(authoritiesMenus)

type authoritiesMenus struct{}

// Init authorities_menus 表数据初始化
// Author SliverHorn
func (a *authoritiesMenus) Init() error {
	entities := []system.AuthoritiesMenus{
		{MenuId: 1, AuthorityId: "888"},
		{MenuId: 2, AuthorityId: "888"},
		{MenuId: 3, AuthorityId: "888"},
		{MenuId: 4, AuthorityId: "888"},
		{MenuId: 5, AuthorityId: "888"},
		{MenuId: 6, AuthorityId: "888"},
		{MenuId: 7, AuthorityId: "888"},
		{MenuId: 8, AuthorityId: "888"},
		{MenuId: 9, AuthorityId: "888"},
		{MenuId: 10, AuthorityId: "888"},
		{MenuId: 11, AuthorityId: "888"},
		{MenuId: 12, AuthorityId: "888"},
		{MenuId: 13, AuthorityId: "888"},
		{MenuId: 14, AuthorityId: "888"},
		{MenuId: 15, AuthorityId: "888"},
		{MenuId: 16, AuthorityId: "888"},
		{MenuId: 17, AuthorityId: "888"},
		{MenuId: 18, AuthorityId: "888"},
		{MenuId: 19, AuthorityId: "888"},
		{MenuId: 20, AuthorityId: "888"},
		{MenuId: 22, AuthorityId: "888"},
		{MenuId: 23, AuthorityId: "888"},
		{MenuId: 24, AuthorityId: "888"},
		{MenuId: 25, AuthorityId: "888"},

		{MenuId: 1, AuthorityId: "8881"},
		{MenuId: 2, AuthorityId: "8881"},
		{MenuId: 8, AuthorityId: "8881"},
		{MenuId: 1, AuthorityId: "9528"},
		{MenuId: 2, AuthorityId: "9528"},
		{MenuId: 3, AuthorityId: "9528"},
		{MenuId: 4, AuthorityId: "9528"},
		{MenuId: 5, AuthorityId: "9528"},
		{MenuId: 6, AuthorityId: "9528"},
		{MenuId: 7, AuthorityId: "9528"},
		{MenuId: 8, AuthorityId: "9528"},
		{MenuId: 9, AuthorityId: "9528"},
		{MenuId: 10, AuthorityId: "9528"},
		{MenuId: 11, AuthorityId: "9528"},
		{MenuId: 12, AuthorityId: "9528"},
		{MenuId: 14, AuthorityId: "9528"},
		{MenuId: 15, AuthorityId: "9528"},
		{MenuId: 16, AuthorityId: "9528"},
		{MenuId: 17, AuthorityId: "9528"},
	}
	return global.Db.Table("authorities_menus").Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ('888', '8881', '9528')").Find(&[]system.AuthoritiesMenus{}).RowsAffected == 48 {
			return errors.New(fmt.Sprintf("%v 表的初始数据已存在!", a.TableName()))
		}
		if err := tx.Create(&entities).Error; err != nil { // 遇到错误时回滚事务
			return errors.Wrap(err, "%v 表初始化数据创建失败!")
		}
		return nil
	})
}

// TableName 定义表名
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *authoritiesMenus) TableName() string {
	var entity system.AuthoritiesMenus
	return entity.TableName()
}
