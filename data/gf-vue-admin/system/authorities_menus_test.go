package system

import (
	"database/sql"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"testing"
)

var (
	menus          system.Menu
	_authorityMenu system.AuthoritiesMenus
)

func init() {
	dsn := "root:gdkid,,..@tcp(127.0.0.1:13307)/gf-vue-admin?charset=utf8mb4&parseTime=True&loc=Local"
	global.Db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func TestAuthoritiesMenus_TableName(t *testing.T) {
	t.Log(AuthoritiesMenus.TableName())
}

func TestAuthorityMenu(t *testing.T) {

	err := global.Db.Debug().Session(&gorm.Session{DryRun: true}).Exec("CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW @table_name AS select @menus.id AS id, @menus.path AS path, @menus.icon AS icon, @menus.name AS name, @menus.sort AS sort, @menus.title AS title, @menus.hidden AS hidden, @menus.component AS component, @menus.parent_id AS parent_id, @menus.created_at AS created_at, @menus.updated_at AS updated_at, @menus.deleted_at AS deleted_at, @menus.keep_alive AS keep_alive, @menus.menu_level AS menu_level, @menus.default_menu AS default_menu, @authorities_menus.menu_id AS menu_id, @authorities_menus.authority_id AS authority_id from (@authorities_menus join @menus on ((@authorities_menus.menu_id = @menus.id)));",
		sql.Named("table_name", "system_authority_menu"),
		sql.Named("menus", menus.TableName()),
		sql.Named("authorities_menus", _authorityMenu.TableName()),
		//map[string]interface{}{"table_name": "system_authority_menu", "menus": menus.TableName(), "authorities_menus": _authorityMenu.TableName()},
	).Error
	if err != nil {
		t.Error(err)
	}

}

func TestName(t *testing.T) {
	sql := "CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `@table_name` AS select `@menus`.id AS id, `@menus`.path AS path, `@menus`.icon AS icon, `@menus`.name AS name, `@menus`.sort AS sort, `@menus`.title AS title, `@menus`.hidden AS hidden, `@menus`.component AS component, `@menus`.parent_id AS parent_id, `@menus`.created_at AS created_at, `@menus`.updated_at AS updated_at, `@menus`.deleted_at AS deleted_at, `@menus`.keep_alive AS keep_alive, `@menus`.menu_level AS menu_level, `@menus`.default_menu AS default_menu, `@authorities_menus`.menu_id AS menu_id, `@authorities_menus`.authority_id AS authority_id from (`@authorities_menus` join `@menus` on ((`@authorities_menus`.menu_id = `@menus`.id)));"
	sql = strings.Replace(sql, "@table_name", "system_authority_menu", -1)
	sql = strings.Replace(sql, "@menus", menus.TableName(), -1)
	sql = strings.Replace(sql, "@authorities_menus", _authorityMenu.TableName(), -1)

	err := global.Db.Debug().Exec(sql).Error
	if err != nil {
		t.Error(err)
		return
	}
}
