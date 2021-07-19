// Package gfva
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
package gfva

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/flipped-aurora/gva/answer"
	"github.com/flipped-aurora/gva/boot"
	"github.com/flipped-aurora/gva/cmd/gfva/internal"
	"github.com/flipped-aurora/gva/library/constant"
	"github.com/flipped-aurora/gva/question"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// initdbCmd represents the initdb command
var initdbCmd = &cobra.Command{
	Use:   "initdb",
	Short: "flipped-aurora/gf-vue-admin/server初始化数据",
	Long:  `flipped-aurora/gf-vue-admin/server初始化数据适配数据库情况: 1. mysql完美适配, 2. postgresql未适配, 3. sqlite未适配, 4. sqlserver未适配`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		input := answer.Frame{}
		if err := survey.Ask(question.Database, &input); err != nil {
			color.Warn.Printf("[cobra] --> 获取用户输入失败! error:%v\n", err)
		}
		switch input.Frame {
		case "gf-vue-admin":
		case "gin-vue-admin":
		case "gin-vue-admin-business":
			boot.Viper.Initialize(path)
			boot.Zap.Initialize()
			internal.DbResolver.Database()
			internal.DbResolver.DataInitialize()
		default:

		}
		return
	},
}

func init() {
	rootCmd.AddCommand(initdbCmd)
	initdbCmd.Flags().StringP("path", "p", constant.ConfigPostgresFile, "自定配置文件路径(绝对路径)")
	initdbCmd.Flags().StringP("frame", "f", "gf", "可选参数为gin,gf")
	initdbCmd.Flags().StringP("type", "t", "mysql", "可选参数为mysql,postgresql,sqlite,sqlserver")
}
