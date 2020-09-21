package core

import (
	"gva/global"

	"github.com/gogf/gf/i18n/gi18n"
)

func init() {
	manager := gi18n.New()
	manager.SetLanguage("en")
	global.I18n = manager
}
