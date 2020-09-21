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
	"fmt"
	"gva/data"
	"gva/data/mysql"
	"gva/data/postgresql"
	"gva/global"

	"github.com/spf13/cobra"
)

// initdbCmd represents the initdb command
var initdbCmd = &cobra.Command{
	Use:   "initdb",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
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
			fmt.Println("功能开发中")
		case "sqlserver":
			fmt.Println("功能开发中")
		default:
			fmt.Println("功能开发中")
		}
		frame, _ := cmd.Flags().GetString("frame")
		if frame == "gf" {
			fmt.Println("功能开发中")
			return
		} else {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(initdbCmd)
	initdbCmd.Flags().StringP("frame", "f", "gin", "可选参数为gin,gf")
	initdbCmd.Flags().StringP("type", "t", "mysql", "可选参数为mysql,postgresql,sqlite,sqlserver")
}
