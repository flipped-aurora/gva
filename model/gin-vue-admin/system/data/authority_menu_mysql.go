//+build mysql

package model

import (
	"errors"
	"fmt"
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
	if err := global.Db.Exec("CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `authority_menu` AS select `system_menus`.`id` AS `id`, `system_menus`.`created_at` AS `created_at`, `system_menus`.`updated_at` AS `updated_at`, `system_menus`.`deleted_at` AS `deleted_at`, `system_menus`.`sort` AS `sort`, `system_menus`.`icon` AS `icon`, `system_menus`.`path` AS `path`, `system_menus`.`name` AS `name`, `system_menus`.`title` AS `title`, `system_menus`.`hidden` AS `hidden`, `system_menus`.`component` AS `component`, `system_menus`.`parent_id` AS `parent_id`, `system_menus`.`menu_level` AS `menu_level`, `system_menus`.`keep_alive` AS `keep_alive`, `system_menus`.`default_menu` AS `default_menu`, `system_authorities_menus`.`authority_id` AS `authority_id`, `system_authorities_menus`.`menu_id` AS `menu_id` from (`system_authorities_menus` join `system_menus` on ((`system_authorities_menus`.`menu_id` = `system_menus`.`id`)))").Error; err != nil {
		return err
	}
	return nil
}

func (a *authorityMenu) CheckDataExist() bool {
	err := global.Db.Unscoped().Find(&[]model.AuthorityMenu{}).Error
	_err := errors.New(fmt.Sprintf("Error 1146: Table '%v.%v' doesn't exist",global.GinVueAdminConfig.Gorm.GetDbName(), a.TableName()))
	if errors.As(err, &_err) {
		return false
	}
	return true
}
