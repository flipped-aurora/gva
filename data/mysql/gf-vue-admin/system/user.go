//go:build mysql
// +build mysql

package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var User = new(user)

type user struct{}

func (u *user) TableName() string {
	var entity system.User
	return entity.TableName()
}

func (u *user) Initialize() error {
	entities := []system.User{
		{Uuid: uuid.NewV4().String(), Username: "admin", Password: "123456", Nickname: "超级管理员", Avatar: "https://qmplusimg.henrongyi.top/gva_header.jpg", AuthorityId: "888"},
		{Uuid: uuid.NewV4().String(), Username: "a303176530", Password: "123456", Nickname: "QMPlusUser", Avatar: "https://qmplusimg.henrongyi.top/1572075907logo.png", AuthorityId: "9528"},
	}

	for i := 0; i < len(entities); i++ {
		if err := entities[i].EncryptedPassword(); err != nil {
			return errors.Wrap(err, "用户密码加密失败!")
		}
	}

	if err := global.Db.Create(&entities).Error; err != nil {
		return errors.Wrap(err, u.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (u *user) CheckDataExist() bool {
	if errors.Is(global.Db.Where("username = ?", "a303176530").First(&system.User{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
