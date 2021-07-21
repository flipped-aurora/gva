// Package internal
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
	"github.com/AlecAivazis/survey/v2"
	"github.com/flipped-aurora/gva/answer"
	"github.com/flipped-aurora/gva/boot"
	gorm "github.com/flipped-aurora/gva/boot/gorm"
	"github.com/flipped-aurora/gva/question"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// initdbCmd represents the initdb command
var initdbCmd = &cobra.Command{
	Use:   "initdb",
	Short: "初始化数据",
	Long:  `初始化数据适配数据库情况: 1. mysql适配中, 2. postgresql适配中, 3. sqlite适配中, 4. sqlserver不考虑不支持`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		input := answer.Frame{}
		if err := survey.Ask(question.Frame, &input); err != nil {
			color.Warn.Printf("[cobra] --> 获取用户输入失败! error:%v\n", err)
			return
		}
		switch input.Frame {
		case "gf-vue-admin":
		case "gin-vue-admin":
		case "gin-vue-admin-business":
			boot.Viper.Initialize(path)
			if err := gorm.DbResolver.LinkDatabase(); err != nil {
				color.Warn.Printf("[cobra] --> 链接数据失败! error:%v\n", err)
				return
			}
			gorm.DbResolver.DataInitialize()
		default:

		}
		return
	},
}

func init() {
	rootCmd.AddCommand(initdbCmd)
	initdbCmd.Flags().StringP("path", "p", gorm.DbResolver.GetConfigPath(), "自定配置文件路径(绝对路径)")
	initdbCmd.Flags().StringP("type", "t", "mysql", "可选参数为mysql、postgresql、sqlite, sqlserver不考虑不支持")
}
