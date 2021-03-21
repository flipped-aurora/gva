package data

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Options struct {
	Gorm  *gorm.DB
	Viper *viper.Viper
}
