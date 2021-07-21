package boot

import (
	boot "github.com/flipped-aurora/gva/boot/gorm"
	"github.com/flipped-aurora/gva/interfaces"
	"github.com/flipped-aurora/gva/library/global"
	system "github.com/flipped-aurora/gva/model/gin-vue-admin/system"
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
		err = global.Db.AutoMigrate(
			new(system.Api),
			new(system.Menu),
			new(system.User),
			new(system.Casbin),
			new(system.Authority),
			new(system.Dictionary),
			new(system.JwtBlacklist),
			new(system.MenuParameter),
			new(system.OperationRecord),
			new(system.DictionaryDetail),
		)
		if err != nil {
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
