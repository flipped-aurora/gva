//go:build postgres
// +build postgres

package system

import (
	"fmt"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
)

var AuthorityMenu = new(authorityMenu)

type authorityMenu struct{}

func (a *authorityMenu) TableName() string {
	var entity system.AuthorityMenu
	return entity.TableName()
}

func (a *authorityMenu) Initialize() error {
	var (
		menus             system.Menu
		authorities_menus system.AuthoritiesMenus
	)
	if err := global.Db.Exec(`
	CREATE VIEW @table_name as
	select @menus.id                       id,
		   @menus.path                     path,
		   @menus.name                     name,
		   @menus.icon                     icon,
		   @menus.sort                     sort,
		   @menus.title                    title,
		   @menus.hidden                   hidden,
		   @menus.parent_id                parent_id,
		   @menus.component                component,
		   @menus.keep_alive               keep_alive,
		   @menus.created_at               created_at,
		   @menus.updated_at               updated_at,
		   @menus.deleted_at               deleted_at,
		   @menus.menu_level               menu_level,
		   @menus.default_menu             default_menu,
		   @authorities_menus.menu_id      menu_id,
		   @authorities_menus.authority_id authority_id
	from (@authorities_menus join @menus on ((@authorities_menus.menu_id = @menus.id)));
	`, map[string]interface{}{"table_name": a.TableName(), "menus": menus.TableName(), "authorities_menus": authorities_menus.TableName()}).Error; err != nil {
		return errors.Wrap(err, a.TableName()+"视图创建失败!")
	}
	return nil
}

func (a *authorityMenu) CheckDataExist() bool {
	err1 := global.Db.Find(&[]system.AuthorityMenu{}).Error
	err2 := errors.New(fmt.Sprintf("Error 1146: Table '%v.%v' doesn't exist", global.Config.Gorm.Dsn.GetDefaultDbName(), a.TableName()))
	if errors.As(err1, &err2) {
		return false
	}
	return true
}
