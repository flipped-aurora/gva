package global

import (
	gfva "github.com/flipped-aurora/gva/library/config/gfva"
	business "github.com/flipped-aurora/gva/library/config/gin-vue-admin"
	gva "github.com/flipped-aurora/gva/library/config/gva"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
)

var (
	Db                *gorm.DB
	Viper             *viper.Viper
	GvaConfig         gva.Config
	GFVAConfig        gfva.Config
	GinVueAdminConfig business.Config
)

type Model struct {
	ID uint `orm:"id" json:"ID" gorm:"primaryKey;comment:主键ID"` // 自增ID

	CreatedAt time.Time `orm:"created_at" json:"CreatedAt" gorm:"column:created_at;comment:创建时间"` // 创建时间

	UpdatedAt time.Time `orm:"updated_at" json:"UpdatedAt" gorm:"column:updated_at;comment:更新时间"` // 更新时间

	DeletedAt gorm.DeletedAt `orm:"deleted_at" json:"-" gorm:"index;column:deleted_at;comment:删除时间"` // 删除时间

	gorm.Model
}
