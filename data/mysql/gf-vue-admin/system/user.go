//go:build mysql
// +build mysql

package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	model "github.com/flipped-aurora/gva/model/gfva/system"
	"github.com/gookit/color"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var User = new(user)

type user struct{}

// Init users 表数据初始化
// Author: [SliverHorn](https://github.com/SliverHorn)
func (u *user) Init() error {
	entities := []system.User{
		{Model: global.Model{ID: 1}, Uuid: uuid.NewV4().String(), Username: "admin", Password: "123456", Nickname: "超级管理员", Avatar: "https://qmplusimg.henrongyi.top/gva_header.jpg", AuthorityId: "888"},
		{Model: global.Model{ID: 2}, Uuid: uuid.NewV4().String(), Username: "a303176530", Password: "123456", Nickname: "QMPlusUser", Avatar: "https://qmplusimg.henrongyi.top/1572075907logo.png", AuthorityId: "9528"},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		for i := range entities {
			if err := entities[i].EncryptedPassword(); err != nil {
				return errors.Wrap(err, "加密密码失败!")
			}
		}
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.Admin{}).RowsAffected == 2 {
			color.Danger.Printf("\n[Mysql] --> %v 表的初始数据已存在!", u.TableName())
			return nil
		}
		if err := tx.Create(&entities).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

// TableName 定义表名
// Author: [SliverHorn](https://github.com/SliverHorn)
func (u *user) TableName() string {
	var entity system.User
	return entity.TableName()
}
