package data

import (
	"fmt"
	"github.com/flipped-aurora/gva/global"
	"github.com/gogf/gf/i18n/gi18n"
)

func init() {
	global.I18n = gi18n.New()
	fmt.Println(global.Viper.GetString("system.language"))
	fmt.Println(global.Viper.GetString("language"))
	if language := global.Viper.GetString("system.language"); language != "" {
		global.I18n.SetLanguage(language)
	}
	if language := global.Viper.GetString("language"); language != "" {
		global.I18n.SetLanguage(language)
	}
}
