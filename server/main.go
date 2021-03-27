package main

import (
	"github.com/flipped-aurora/gva/boot"
	gfData "github.com/flipped-aurora/gva/data/gf"
	ginData "github.com/flipped-aurora/gva/data/gin"
	"github.com/flipped-aurora/gva/global"
	gfModel "github.com/flipped-aurora/gva/model/gf"
	ginModel "github.com/flipped-aurora/gva/model/gin"
	"github.com/gogf/gf/i18n/gi18n"

	"github.com/gin-gonic/gin"
	"github.com/zserge/lorca"
	"log"
	"net/http"
	"runtime"
)

////go:embed www
//var fs embed.FS

func main() {
	args := make([]string, 0, 1)
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	ui, err := lorca.New("", "", 1920, 1080, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = ui.Close()
	}()

	go func() {
		engine := gin.Default()
		group := engine.Group("")
		group.POST("initdb", func(context *gin.Context) {
			if err = context.ShouldBind(&global.Config); err != nil {
				context.JSON(http.StatusOK, gin.H{"code": 7, "err": err, "message": "接收数据失败!"})
				return
			}
			global.I18n = gi18n.New()
			global.I18n.SetLanguage(global.Config.Language)
			switch global.Config.FrameType {
			case "gin-vue-admin":
				boot.Mysql.Initialize()
				ginModel.GinVueAdminAutoMigrateTables(global.Db)
				_ = ginData.GinVueAdmin()
			case "gf-vue-admin":
				boot.Mysql.Initialize()
				gfModel.GfVueAdminAutoMigrateTables(global.Db)
				_ = gfData.GfVueAdmin()
			default:
				boot.Mysql.Initialize()
				ginModel.GinVueAdminAutoMigrateTables(global.Db)
				_ = ginData.GinVueAdmin()
			}
		})
		_ = engine.Run(":7777")
	}()
	select {}
}
