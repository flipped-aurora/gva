package model

import (
	"github.com/flipped-aurora/gva/library/global"
	model "github.com/flipped-aurora/gva/model/gin-vue-admin/system"
	_errors "github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

var Authority = new(authority)

type authority struct{}

func (a *authority) TableName() string {
	var entity model.Authority
	return entity.TableName()
}

func (a *authority) Initialize() error {
	entities := []model.Authority{
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "888", AuthorityName: "超级管理员", ParentId: "0", DefaultRouter: "dashboard"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "8881", AuthorityName: "超级管理员子角色", ParentId: "888", DefaultRouter: "dashboard"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "9528", AuthorityName: "测试角色", ParentId: "0", DefaultRouter: "dashboard"},
	}
	if err := global.Db.Create(&entities).Error; err != nil { // 创建 model.Authority 初始化数据
		return _errors.Wrap(err, a.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (a *authority) CheckDataExist() bool {
	if _errors.Is(global.Db.Where("authority_id = ? AND authority_name = ?", "9528", "测试角色").First(&model.Authority{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
