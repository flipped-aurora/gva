package boot

import (
	boot "github.com/flipped-aurora/gva/boot/gorm"
	"github.com/flipped-aurora/gva/interfaces"
	"github.com/flipped-aurora/gva/library/global"
	"github.com/gookit/color"
	"gorm.io/gorm"
	"os"
)

var Gorm = new(_gorm)

type _gorm struct{}

func (g *_gorm) Initialize(i interfaces.Gorm) {
	resolver := i.GetResolver()
	db, err := gorm.Open(i.GetGormDialector(), boot.Gorm.GenerateConfig())
	if err != nil {
		color.Error.Println("gorm链接数据库失败! err:", err)
		os.Exit(0)
	}
	err = db.Use(resolver)
	if err != nil {
		color.Error.Println("gorm链接集群失败! err:", err)
		os.Exit(0)
	}
	global.Db = db
	if global.GinVueAdminConfig.Gorm.AutoMigrate {
		if err = i.AutoMigrate(); err != nil {
			color.Error.Println("gorm注册表失败! err:", err)
			os.Exit(0)
		}
		color.Info.Println("gorm注册表成功!")
	}
	sql, _err := db.DB()
	if _err != nil {
		color.Error.Println("gorm设置数据库最大连接数失败! err:", err)
		return
	}
	sql.SetMaxIdleConns(global.GinVueAdminConfig.Gorm.GetMaxIdleConnes())
	sql.SetMaxOpenConns(global.GinVueAdminConfig.Gorm.GetMaxOpenConnes())
}
