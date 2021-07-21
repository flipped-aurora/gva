package boot

import (
	"flag"
	"fmt"
	boot "github.com/flipped-aurora/gva/boot/gorm"
	"github.com/flipped-aurora/gva/library/constant"
	"github.com/flipped-aurora/gva/library/global"
	"github.com/gookit/color"
	"github.com/spf13/viper"
	"os"
)

var Viper = new(_viper)

type _viper struct {
	path string
}

func (v *_viper) Initialize(path ...string) {
	if len(path) == 0 {
		flag.StringVar(&v.path, "c", "", "choose config file.")
		flag.Parse()
		if v.path == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if env := os.Getenv(constant.ConfigEnv); env == "" {
				if p := boot.DbResolver.GetConfigPath(); p != "" {
					v.path = p
					fmt.Println(`您正在使用 boot.DbResolver.GetConfigPath() 方法传递的变量, config的路径为: `, v.path)
				} else {
					v.path = constant.ConfigDevelopFile
					fmt.Println(`您现在的环境是 Develop, config的路径为: `, v.path)
				}
			} else {
				v.path = constant.ConfigEnv
				color.Info.Println(`您正在使用GVA_CONFIG环境变量,config的路径为: `, v.path)
			}
		} else {
			color.Info.Println(`您正在使用命令行的 -c 参数传递的值! path:`, v.path)
		}
	} else {
		v.path = path[0]
		color.Info.Println(`您正在使用func (v *_viper) Initialize()传递的值! path:`, v.path)
	}

	_v := viper.New()
	_v.SetConfigFile(v.path)
	if err := _v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf(`读取config.yaml文件失败, err: %v`, err))
	}
	if err := _v.Unmarshal(&global.GFVAConfig); err != nil {
		fmt.Println(`Json 序列化数据失败, err :`, err)
	} else {
		global.Viper = _v
	}
	if err := _v.Unmarshal(&global.GinVueAdminConfig); err != nil {
		fmt.Println(`Json 序列化数据失败, err :`, err)
	} else {
		global.GinVueAdminConfig.Gorm.AutoMigrate = true
		global.Viper = _v
	}
}
