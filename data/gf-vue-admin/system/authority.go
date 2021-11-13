package system

import (
	"fmt"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var Authority = new(authority)

type authority struct{}

func (a *authority) TableName() string {
	var entity system.Authority
	return entity.TableName()
}

func (a *authority) Initialize() error {
	entities := []system.Authority{
		{AuthorityId: "9528", AuthorityName: "测试角色", ParentId: "0", DefaultRouter: "dashboard"},
		{AuthorityId: "888", AuthorityName: "超级管理员", ParentId: "0", DefaultRouter: "dashboard"},
		{AuthorityId: "8881", AuthorityName: "超级管理员子角色", ParentId: "888", DefaultRouter: "dashboard"},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if !errors.Is(global.Db.Where(entities[2]).First(&system.Authority{}).Error, gorm.ErrRecordNotFound) {
			return errors.New(fmt.Sprintf("%v 表数据初始数据已存在!", a.TableName()))
		} // 判断是否存在数据

		if err := tx.Create(&entities).Error; err != nil {
			return errors.Wrapf(err, "%s表数据初始化失败!", a.TableName())
		} // 初始化数据
		return nil
	})
}

func (a *authority) CheckDataExist() bool {
	if errors.Is(global.Db.Where("authority_id = ?", "8881").First(&system.Authority{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
