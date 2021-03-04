package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	Db    *gorm.DB
	Viper *viper.Viper
)
