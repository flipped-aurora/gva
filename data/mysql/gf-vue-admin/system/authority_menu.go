package system

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	_errors "github.com/pkg/errors"
)

var AuthorityMenu = new(authorityMenu)

type authorityMenu struct{}

func (a *authorityMenu) TableName() string {
	var entity system.AuthorityMenu
	return entity.TableName()
}

func (a *authorityMenu) Initialize() error {
	sql := fmt.Sprintf("CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `%s` AS select `system_menus`.`id` AS `id`, `system_menus`.`created_at` AS `created_at`, `system_menus`.`updated_at` AS `updated_at`, `system_menus`.`deleted_at` AS `deleted_at`, `system_menus`.`menu_level` AS `menu_level`, `system_menus`.`parent_id`  AS `parent_id`, `system_menus`.`path` AS `path`, `system_menus`.`name` AS `name`, `system_menus`.`hidden`     AS `hidden`, `system_menus`.`component`  AS `component`, `system_menus`.`title` AS `title`, `system_menus`.`icon` AS `icon`, `system_menus`.`sort` AS `sort`, `system_authorities_menus`.`authority_id` AS `authority_id`, `system_authorities_menus`.`menu_id` AS `menu_id`, `system_menus`.`keep_alive` AS `keep_alive`, `system_menus`.`default_menu` AS `default_menu` from (`system_authorities_menus` join `system_menus` on ((`system_authorities_menus`.`menu_id` = `system_menus`.`id`)));", a.TableName())
	if err := global.Db.Exec(sql).Error; err != nil {
		return _errors.Wrap(err, a.TableName()+"视图创建失败!")
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
