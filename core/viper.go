package core

import (
	"fmt"
	"gva/global"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

func Viper() {
	v := viper.New()
	v.SetConfigFile(global.ConfigPath)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}
	global.Viper = v
}
