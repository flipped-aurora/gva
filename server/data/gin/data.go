package data

import (
	extra "github.com/flipped-aurora/gva/data/gin/extra"
	system "github.com/flipped-aurora/gva/data/gin/system"
	workflow "github.com/flipped-aurora/gva/data/gin/workflow"
	"github.com/flipped-aurora/gva/interfaces"
)

func GinVueAdmin() error {
	return interfaces.InitDb(
		system.Api,
		system.Admin,
		system.Casbin,
		system.Authority,
		system.Dictionary,
		system.AuthorityMenu,
		system.ResourcesAuthorities,
		system.DictionaryDetail,
		system.AuthoritiesMenus,

		workflow.Edge,
		workflow.Node,
		workflow.EndPoint,
		workflow.StartPoint,
		workflow.Process,

		extra.File,
	)
}
