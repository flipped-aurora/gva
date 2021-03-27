package main

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gva/boot"
	gfData "github.com/flipped-aurora/gva/data/gf"
	ginData "github.com/flipped-aurora/gva/data/gin"
	"github.com/flipped-aurora/gva/global"
	gfModel "github.com/flipped-aurora/gva/model/gf"
	ginModel "github.com/flipped-aurora/gva/model/gin"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/i18n/gi18n"
	"github.com/zserge/lorca"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)


func main() {
	args := make([]string, 0, 1)
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	dir, _ := os.Getwd()
	path := filepath.Join(dir, "dist")
	//file, _ := os.Open(path + "index.html")
	fmt.Println(path)
	ui, err := lorca.New("", path, 1920, 1080, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = ui.Close()
	}()

	go func() {
		engine := gin.Default()
		engine.StaticFS("/web", http.Dir(path))
		_ = engine.Run(":7777")
	}()

	_ = ui.Bind("submit", func(dsn string) bool {
		if err = json.Unmarshal([]byte(dsn), &global.Config); err != nil {
			fmt.Println(err)
			return false
		}
		fmt.Println(global.Config)
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
		defer func() {
			_ = ui.Close()
			os.Exit(0)
		}()
		return true
	})


	_ = ui.Load("http://127.0.0.1:7777/web/index.html")
	select {}
}
