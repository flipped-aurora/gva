package global

import (
	"github.com/flipped-aurora/gva-ctl/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const ConfigPath = "./config.yaml"

var (
	Db     *gorm.DB
	Viper  *viper.Viper
	Config config.Server
)
