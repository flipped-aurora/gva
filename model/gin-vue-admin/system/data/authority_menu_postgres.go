//+build postgres

package model

import (
	"github.com/flipped-aurora/gva/library/global"
	model "github.com/flipped-aurora/gva/model/gin-vue-admin/system"
)

var AuthorityMenu = new(authorityMenu)

type authorityMenu struct{}

func (a *authorityMenu) TableName() string {
	var entity model.AuthorityMenu
	return entity.TableName()
}

func (a *authorityMenu) Initialize() error {
	if err := global.Db.Exec("CREATE VIEW system_authority_menu as select system_menus.id id, system_menus.created_at created_at, system_menus.updated_at updated_at,system_menus.deleted_at deleted_at, system_menus.menu_level menu_level, system_menus.parent_id parent_id, system_menus.path path, system_menus.name as name, system_menus.hidden hidden, system_menus.component component, system_menus.title title,system_menus.icon icon, system_menus.sort sort, system_authorities_menus.authority_id authority_id, system_authorities_menus.menu_id menu_id, system_menus.keep_alive keep_alive, system_menus.default_menu default_menu from (system_authorities_menus join system_menus on ((system_authorities_menus.menu_id = system_menus.id)));").Error; err != nil {
		return err
	}
	return nil
}

func (a *authorityMenu) CheckDataExist() bool {
	if global.Db.Find(&[]model.AuthorityMenu{}).RowsAffected == 0 {
		return false
	}
	return true
}
