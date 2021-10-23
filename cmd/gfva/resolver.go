package gfva

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/example"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	data "github.com/flipped-aurora/gva/data/gf-vue-admin/system"
	"github.com/flipped-aurora/gva/interfaces"
)

type resolver struct{}

func (r *resolver) AutoMigrate() error {
	return global.Db.AutoMigrate(
		new(system.Api),
		new(system.Menu),
		new(system.User),
		new(system.Casbin),
		new(system.Authority),
		new(system.Dictionary),
		new(system.JwtBlacklist),
		new(system.MenuParameter),
		new(system.OperationRecord),
		new(system.DictionaryDetail),

		new(example.File),
		new(example.Customer),
	)
}

func (r *resolver) DataInitialize() {
	_ = interfaces.DataInitialize(
		data.Api,
		data.User,
		data.Menu,
		data.Casbin,
		data.Authority,
		data.Dictionary,
		data.AuthorityMenu,
		data.UserAuthority,
		data.DictionaryDetail,
		data.AuthoritiesMenus,
		data.AuthoritiesResources,
	)
}
