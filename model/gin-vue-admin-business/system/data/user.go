// +build postgres

package data

import (
	"errors"
	"github.com/flipped-aurora/gva/library/fmt"
	"github.com/flipped-aurora/gva/library/global"
	model "github.com/flipped-aurora/gva/model/gin-vue-admin-business/system"
	_errors "github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

var User = new(user)

type user struct{}

// DataInit  system_users 表数据初始化
// Author [SliverHorn](https://github.com/SliverHorn)
func (u *user) DataInit() error {
	entities := []model.User{
		{Model: global.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Uuid: uuid.NewV4().String(), Username: "admin", Password: "123456", Nickname: "超级管理员", Avatar: "http://qmplusimg.henrongyi.top/gva_header.jpg", AuthorityId: "888"},
		{Model: global.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Uuid: uuid.NewV4().String(), Username: "a303176530", Password: "123456", Nickname: "示例用户", Avatar: "http://qmplusimg.henrongyi.top/1572075907logo.png", AuthorityId: "9528"},
	}

	for i := 0; i < len(entities); i++ { // 加密密码
		if err := entities[i].EncryptedPassword(); err != nil {
			return _errors.Wrap(err, "密码加密失败!")
		}
	}

	var entity model.User
	if !errors.Is(global.Db.Where("id = ?", 2).First(&entity).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		fmt.Printf(fmt.InitDataExist, "postgres", entity.TableName())
		return nil
	}

	if err := global.Db.Create(&entities).Error; err != nil { // 创建 model.User 初始化数据
		return err
	}
	fmt.Printf(fmt.InitDataSuccess, "postgres", entity.TableName())

	return nil
}
