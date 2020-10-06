package core

import (
	"os"

	"github.com/flipped-aurora/gva-ctl/global"

	"github.com/gookit/color"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

func Viper(path string) {
	v := viper.New()
	if path != "" {
		v.SetConfigFile(path)
	} else {
		v.SetConfigFile(global.ConfigPath)
	}
	err := v.ReadInConfig()
	if err != nil {
		color.Warn.Printf("读取config.yaml文件失败,err:%s \n", err)
		os.Exit(0)
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		color.Warn.Printf("更新%v文件成功:", e.Name)
		if err := v.Unmarshal(&global.Config); err != nil {
			color.Error.Println(err)
		}
	})
	if err := v.Unmarshal(&global.Config); err != nil {
		color.Error.Println(err)
	}
	global.Viper = v
}
