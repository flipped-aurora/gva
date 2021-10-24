//go:build mysql
// +build mysql

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
	sql := "CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `@table_name` AS select `@menus`.id AS id, `@menus`.path AS path, `@menus`.icon AS icon, `@menus`.name AS name, `@menus`.sort AS sort, `@menus`.title AS title, `@menus`.hidden AS hidden, `@menus`.component AS component, `@menus`.parent_id AS parent_id, `@menus`.created_at AS created_at, `@menus`.updated_at AS updated_at, `@menus`.deleted_at AS deleted_at, `@menus`.keep_alive AS keep_alive, `@menus`.menu_level AS menu_level, `@menus`.default_menu AS default_menu, `@authorities_menus`.menu_id AS menu_id, `@authorities_menus`.authority_id AS authority_id from (`@authorities_menus` join `@menus` on ((`@authorities_menus`.menu_id = `@menus`.id)));"
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
