/*
Copyright © 2020 SliverHorn

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/flipped-aurora/gva/core"
	"github.com/flipped-aurora/gva/data"
	"github.com/flipped-aurora/gva/data/mysql"
	"github.com/flipped-aurora/gva/data/postgresql"
	"github.com/flipped-aurora/gva/global"
	"github.com/gookit/color"

	"github.com/spf13/cobra"
)

// initdbCmd represents the initdb command
var initdbCmd = &cobra.Command{
	Use:   "initdb",
	Short: "gin-vue-admin初始化数据",
	Long: `gin-vue-admin初始化数据适配数据库情况: 
1. mysql完美适配,
2. postgresql不能保证完美适配,
3. sqlite未适配,
4. sqlserver未适配`,
	Run: func(cmd *cobra.Command, args []string) {
		if path, _ := cmd.Flags().GetString("path"); path != "" {
			core.Viper(path)
		}
		db := data.Gorm()
		if global.Config.System.DbType == "" {
			dbType, _ := cmd.Flags().GetString("type")
			global.Config.System.DbType = dbType
		}
		switch global.Config.System.DbType {
		case "mysql":
			mysql.InitMysqlTables(db)
			mysql.InitMysqlData(db)
		case "postgresql":
			postgresql.InitPostgresqlTables(db)
			postgresql.InitPostgresqlData(db)
		case "sqlite":
			color.Info.Println("功能开发中")
		case "sqlserver":
			color.Info.Println("功能开发中")
		default:
			color.Info.Println("功能开发中")
		}
		frame, _ := cmd.Flags().GetString("frame")
		if frame == "gf" {
			color.Info.Println("功能开发中")
			return
		} else {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(initdbCmd)
	initdbCmd.Flags().StringP("path", "p", "./config.yaml", "自定配置文件路径(绝对路径)")
	initdbCmd.Flags().StringP("frame", "f", "gin", "可选参数为gin,gf")
	initdbCmd.Flags().StringP("type", "t", "mysql", "可选参数为mysql,postgresql,sqlite,sqlserver")
}
