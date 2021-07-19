package model

import (
	"github.com/flipped-aurora/gva/library/global"
	_errors "github.com/pkg/errors"
	"gorm.io/gorm"
)

var AuthoritiesResources = new(authoritiesResources)

type (
	authoritiesResources struct{}
	AuthorityResources   struct {
		AuthorityId string `gorm:"column:authority_id"`
		ResourcesId string `gorm:"column:resources_id"`
	}
)

func (a *AuthorityResources) TableName() string {
	// todo 反射 /Users/sliverhorn/Go/pkg/mod/gorm.io/gorm@v1.21.10/schema/utils.go
	return ""
}

func (a *authoritiesResources) TableName() string {
	// todo 反射 /Users/sliverhorn/Go/pkg/mod/gorm.io/gorm@v1.21.10/schema/utils.go
	return ""
}

func (a *authoritiesResources) Initialize() error {
	entities := []AuthorityResources{
		{AuthorityId: "888", ResourcesId: "888"},
		{AuthorityId: "888", ResourcesId: "8881"},
		{AuthorityId: "888", ResourcesId: "9528"},
		{AuthorityId: "9528", ResourcesId: "8881"},
		{AuthorityId: "9528", ResourcesId: "9528"},
	}
	if err := global.Db.Create(&entities).Error; err != nil { // 创建 model.Authority 初始化数据
		return _errors.Wrap(err, a.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (a *authoritiesResources) CheckDataExist() bool {
	if _errors.Is(global.Db.Where("authority_id = ? AND resources_id = ?", "9528", "9528").First(&AuthorityResources{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
