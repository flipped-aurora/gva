package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var UserAuthority = new(userAuthority)

type userAuthority struct{}

func (u *userAuthority) TableName() string {
	var entity system.UseAuthority
	return entity.TableName()
}

func (u *userAuthority) Initialize() error {
	entities := []system.UseAuthority{
		{UserId: 1, AuthorityId: "888"},
		{UserId: 1, AuthorityId: "8881"},
		{UserId: 1, AuthorityId: "9528"},
		{UserId: 2, AuthorityId: "888"},
	}
	if err := global.Db.Create(&entities).Error; err != nil {
		return errors.Wrap(err, u.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (u *userAuthority) CheckDataExist() bool {
	if errors.Is(global.Db.Where("user_id = ? AND authority_id = ?", 2, "888").First(&system.UseAuthority{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
