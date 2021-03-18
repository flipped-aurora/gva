package boot

import (
	"flag"
	"fmt"
	"github.com/flipped-aurora/gva/constant"
	"github.com/flipped-aurora/gva/global"
	"github.com/spf13/viper"
	"os"
)

var Viper = new(_viper)

type _viper struct{}

func (v *_viper) Initialize(path ...string) {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(constant.ConfigEnv); configEnv == "" {
				config = constant.ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", constant.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("您正在使用GVA_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Initialize()传递的值,config的路径为%v\n", config)
	}

	_v := viper.New()
	_v.SetConfigFile(config)
	err := _v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	global.Viper = _v
}
