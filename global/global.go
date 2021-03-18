package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
)
var (
	Db         *gorm.DB
	Viper      *viper.Viper
	Redis      *redis.Client
	GormViper  *viper.Viper
)


type Model struct {
	ID        uint           `orm:"id" json:"ID" gorm:"primarykey"`
	CreatedAt time.Time      `orm:"created_at" json:"CreatedAt"`
	UpdatedAt time.Time      `orm:"updated_at" json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `orm:"deleted_at" json:"-" gorm:"index"`
}
