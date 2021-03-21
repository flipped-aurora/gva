package data

import (
	extra "github.com/flipped-aurora/gva/data/gf/extra"
	system "github.com/flipped-aurora/gva/data/gf/system"
	workflow "github.com/flipped-aurora/gva/data/gf/workflow"
	"github.com/flipped-aurora/gva/global"
	"github.com/flipped-aurora/gva/interfaces"
	"gorm.io/gorm"
)

func GinVueAdmin(db ...*gorm.DB) error {
	if len(db) > 0 {
		global.Db = db[0]
	}
	return interfaces.InitDb(
		system.Api,
		system.Menu,
		system.Admin,
		system.Casbin,
		system.Authority,
		system.Dictionary,
		system.AuthorityMenu,
		system.AuthoritiesMenus,
		system.ResourcesAuthorities,
		system.DictionaryDetail,

		workflow.Edge,
		workflow.Node,
		workflow.EndPoint,
		workflow.StartPoint,
		workflow.Process,

		extra.File,
	)
}
