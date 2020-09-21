package global

import (
	"gva/config"

	"github.com/gogf/gf/i18n/gi18n"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const ConfigPath = "./config.yaml"

var (
	Db     *gorm.DB
	Log    *zap.Logger
	I18n   *gi18n.Manager
	Viper  *viper.Viper
	Config config.Server
)
