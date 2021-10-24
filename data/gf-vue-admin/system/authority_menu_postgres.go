//go:build postgres
// +build postgres

package system

import (
	"fmt"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
	"strings"
)

var AuthorityMenu = new(authorityMenu)

type authorityMenu struct{}

func (a *authorityMenu) TableName() string {
	var entity system.AuthorityMenu
	return entity.TableName()
}

func (a *authorityMenu) Initialize() error {
	var (
		menus          system.Menu
		_authorityMenu system.AuthoritiesMenus
	)
	sql := `
	CREATE VIEW @table_name as
	select @menus.id                       as id,
		   @menus.path                     as path,
		   @menus.name                     as name,
		   @menus.icon                     as icon,
		   @menus.sort                     as sort,
		   @menus.title                    as title,
		   @menus.hidden                   as hidden,
		   @menus.parent_id                as parent_id,
		   @menus.component                as component,
		   @menus.keep_alive               as keep_alive,
		   @menus.created_at               as created_at,
		   @menus.updated_at               as updated_at,
		   @menus.deleted_at               as deleted_at,
		   @menus.menu_level               as menu_level,
		   @menus.default_menu             as default_menu,
		   @authorities_menus.menu_id      as menu_id,
		   @authorities_menus.authority_id as authority_id
	from (@authorities_menus join @menus on ((@authorities_menus.menu_id = @menus.id)));`
	sql = strings.Replace(sql, "@table_name", "system_authority_menu", -1)
	sql = strings.Replace(sql, "@menus", menus.TableName(), -1)
	sql = strings.Replace(sql, "@authorities_menus", _authorityMenu.TableName(), -1)
	if err := global.Db.Exec(sql).Error; err != nil {
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
