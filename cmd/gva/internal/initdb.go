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
package internal

import (
	"github.com/flipped-aurora/tool/cmd/gva/internal/boot"
	"github.com/flipped-aurora/tool/cmd/gva/internal/global"
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
		path, _ := cmd.Flags().GetString("path")
		global.Viper = boot.Viper(path)

		boot.Mysql.CheckDatabase()
		boot.Mysql.CheckUtf8mb4()
		boot.Mysql.Info()
		boot.Mysql.Init()
		if global.Viper.GetString("system.db-type") == "mysql" {
			boot.Mysql.AutoMigrateTables()
			boot.Mysql.InitData()
		}
	},
}

func init() {
	rootCmd.AddCommand(initdbCmd)
	initdbCmd.Flags().StringP("path", "p", "./config.yaml", "自定配置文件路径(绝对路径)")
	initdbCmd.Flags().StringP("type", "t", "mysql", "可选参数为mysql,postgresql,sqlite,sqlserver")
}
