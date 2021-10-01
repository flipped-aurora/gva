package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gva/library/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"reflect"
)

var AuthoritiesMenus = new(authoritiesMenus)

type authoritiesMenus struct{}

func (a *authoritiesMenus) TableName() string {
	var entity system.Authority
	types := reflect.TypeOf(entity)
	if s, o := types.FieldByName("Menus"); o {
		m1 := schema.ParseTagSetting(s.Tag.Get("gorm"), ";")
		return m1["MANY2MANY"]
	}
	return ""
}

func (a *authoritiesMenus) Initialize() error {
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
	if err := global.Db.Create(&entities).Error; err != nil {
		return errors.Wrap(err, a.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (a *authoritiesMenus) CheckDataExist() bool {
	if errors.Is(global.Db.Where("menu_id = ? AND authority_id = ?", 17, "9528").First(&system.AuthoritiesMenus{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
