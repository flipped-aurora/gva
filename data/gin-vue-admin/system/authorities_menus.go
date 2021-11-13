package system

import (
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var AuthoritiesMenus = new(authoritiesMenus)

type authoritiesMenus struct{}

func (a *authoritiesMenus) TableName() string {
	return "sys_authority_menus"
}

func (a *authoritiesMenus) Initialize() error {
	entities := []AuthorityMenus{
		{BaseMenuId: 1, AuthorityId: "888"},
		{BaseMenuId: 2, AuthorityId: "888"},
		{BaseMenuId: 3, AuthorityId: "888"},
		{BaseMenuId: 4, AuthorityId: "888"},
		{BaseMenuId: 5, AuthorityId: "888"},
		{BaseMenuId: 6, AuthorityId: "888"},
		{BaseMenuId: 7, AuthorityId: "888"},
		{BaseMenuId: 8, AuthorityId: "888"},
		{BaseMenuId: 9, AuthorityId: "888"},
		{BaseMenuId: 10, AuthorityId: "888"},
		{BaseMenuId: 11, AuthorityId: "888"},
		{BaseMenuId: 12, AuthorityId: "888"},
		{BaseMenuId: 13, AuthorityId: "888"},
		{BaseMenuId: 14, AuthorityId: "888"},
		{BaseMenuId: 15, AuthorityId: "888"},
		{BaseMenuId: 16, AuthorityId: "888"},
		{BaseMenuId: 17, AuthorityId: "888"},
		{BaseMenuId: 18, AuthorityId: "888"},
		{BaseMenuId: 19, AuthorityId: "888"},
		{BaseMenuId: 20, AuthorityId: "888"},
		{BaseMenuId: 22, AuthorityId: "888"},
		{BaseMenuId: 23, AuthorityId: "888"},
		{BaseMenuId: 24, AuthorityId: "888"},
		{BaseMenuId: 25, AuthorityId: "888"},

		{BaseMenuId: 1, AuthorityId: "8881"},
		{BaseMenuId: 2, AuthorityId: "8881"},
		{BaseMenuId: 8, AuthorityId: "8881"},

		{BaseMenuId: 1, AuthorityId: "9528"},
		{BaseMenuId: 2, AuthorityId: "9528"},
		{BaseMenuId: 3, AuthorityId: "9528"},
		{BaseMenuId: 4, AuthorityId: "9528"},
		{BaseMenuId: 5, AuthorityId: "9528"},
		{BaseMenuId: 6, AuthorityId: "9528"},
		{BaseMenuId: 7, AuthorityId: "9528"},
		{BaseMenuId: 8, AuthorityId: "9528"},
		{BaseMenuId: 9, AuthorityId: "9528"},
		{BaseMenuId: 10, AuthorityId: "9528"},
		{BaseMenuId: 11, AuthorityId: "9528"},
		{BaseMenuId: 12, AuthorityId: "9528"},
		{BaseMenuId: 14, AuthorityId: "9528"},
		{BaseMenuId: 15, AuthorityId: "9528"},
		{BaseMenuId: 16, AuthorityId: "9528"},
		{BaseMenuId: 17, AuthorityId: "9528"},
	}
	if err := global.Db.Create(&entities).Error; err != nil {
		return errors.Wrap(err, a.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (a *authoritiesMenus) CheckDataExist() bool {
	if errors.Is(global.Db.Where("menu_id = ? AND authority_id = ?", 17, "9528").First(&AuthorityMenus{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}

type AuthorityMenus struct {
	AuthorityId string `gorm:"column:sys_authority_authority_id"`
	BaseMenuId  uint   `gorm:"column:sys_base_menu_id"`
}

func (a *AuthorityMenus) TableName() string {
	return "sys_authority_menus"
}
