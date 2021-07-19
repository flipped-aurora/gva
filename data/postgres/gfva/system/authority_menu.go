//+build postgres

package data

import (
	"github.com/flipped-aurora/gva/library/global"
	"github.com/gookit/color"
)

var AuthorityMenu = new(authorityMenu)

type authorityMenu struct{}

// Init authority_menu 视图 创建
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *authorityMenu) Init() error {
	if global.Db.Raw("select * from authority_menu").RowsAffected > 0 {
		color.Danger.Println("\n[postgres] --> authority_menu 视图已存在!")
		return nil
	}
	if err := global.Db.Exec(`
CREATE VIEW authority_menu as
select menus.id                       id,
       menus.created_at               created_at,
       menus.updated_at               updated_at,
       menus.deleted_at               deleted_at,
       menus.menu_level               menu_level,
       menus.parent_id                parent_id,
       menus.path                     path,
       menus.name                     menus_name,
       menus.hidden                   hidden,
       menus.component                component,
       menus.title                    title,
       menus.icon                     icon,
       menus.sort                     sort,
       authorities_menus.authority_id authority_id,
       authorities_menus.menu_id      menu_id,
       menus.keep_alive               keep_alive,
       menus.default_menu 			  default_menu
from (authorities_menus
         join menus on ((authorities_menus.menu_id = menus.id)));
`).Error; err != nil {
		return err
	}
	return nil
}

// TableName 定义表名
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *authorityMenu) TableName() string {
	return "authority_menu 视图"
}
