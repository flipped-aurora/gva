package system

import (
	"fmt"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gva/library/global"
	"github.com/pkg/errors"
)

var AuthorityMenu = new(authorityMenu)

type authorityMenu struct{}

func (a *authorityMenu) TableName() string {
	var entity system.AuthorityMenu
	return entity.TableName()
}

func (a *authorityMenu) Initialize() error {

	sql := fmt.Sprintf("CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `%v` AS select `menus`.`id` AS `id`, `menus`.`created_at` AS `created_at`, `menus`.`updated_at` AS `updated_at`, `menus`.`deleted_at` AS `deleted_at`, `menus`.`menu_level` AS `menu_level`, `menus`.`parent_id` AS `parent_id`, `menus`.`path` AS `path`, `menus`.`name` AS `name`, `menus`.`hidden` AS `hidden`, `menus`.`component` AS `component`, `menus`.`title` AS `title`, `menus`.`icon` AS `icon`, `menus`.`sort` AS `sort`, `authorities_menus`.`authority_id` AS `authority_id`, `authorities_menus`.`menu_id` AS `menu_id`, `menus`.`keep_alive` AS `keep_alive`, `menus`.`default_menu` AS `default_menu`from (`authorities_menus` join `menus` on ((`authorities_menus`.`menu_id` = `menus`.`id`)));", a.TableName())
	if err := global.Db.Exec(sql).Error; err != nil {
		return errors.Wrap(err, a.TableName()+"视图创建失败!")
	}
	return nil
}

func (a *authorityMenu) CheckDataExist() bool {
	err := errors.New(fmt.Sprintf("Error 1146: Table '%v.%v' doesn't exist", global.GinVueAdminConfig.Gorm.GetDbName(), a.TableName()))
	if errors.As(global.Db.Find(&[]system.AuthorityMenu{}).Error, &err) {
		return true
	}
	return false
}
