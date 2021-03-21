package data

import (
	"github.com/flipped-aurora/gva/data"
	extra "github.com/flipped-aurora/gva/data/gf/extra"
	system "github.com/flipped-aurora/gva/data/gf/system"
	workflow "github.com/flipped-aurora/gva/data/gf/workflow"
	"github.com/flipped-aurora/gva/global"
	"github.com/flipped-aurora/gva/interfaces"
	"github.com/gogf/gf/i18n/gi18n"
)

func GinVueAdmin(options ...data.Options) error {
	if len(options) != 0 {
		for _, option := range options {
			if option.Viper != nil {
				global.Viper = option.Viper
				global.I18n = gi18n.New()
				if language := global.Viper.GetString("language"); language != "" {
					global.I18n.SetLanguage(language)
				}
			}
			if option.Gorm != nil {
				global.Db = option.Gorm
			}
		}
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
