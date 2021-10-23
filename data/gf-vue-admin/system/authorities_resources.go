package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"reflect"
)

var AuthoritiesResources = new(authoritiesResources)

type authoritiesResources struct{}

func (a *authoritiesResources) TableName() string {
	var entity system.Authority
	types := reflect.TypeOf(entity)
	if s, o := types.FieldByName("Resources"); o {
		m1 := schema.ParseTagSetting(s.Tag.Get("gorm"), ";")
		return m1["MANY2MANY"]
	}
	return ""
}

func (a *authoritiesResources) Initialize() error {
	entities := []system.AuthoritiesResources{
		{AuthorityId: "888", ResourcesId: "888"},
		{AuthorityId: "888", ResourcesId: "8881"},
		{AuthorityId: "888", ResourcesId: "9528"},
		{AuthorityId: "9528", ResourcesId: "8881"},
		{AuthorityId: "9528", ResourcesId: "9528"},
	}
	if err := global.Db.Create(&entities).Error; err != nil {
		return errors.Wrap(err, a.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (a *authoritiesResources) CheckDataExist() bool {
	if errors.Is(global.Db.Where("authority_id = ? AND resources_id = ?", "9528", "9528").First(&system.AuthoritiesResources{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
