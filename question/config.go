package question

import "github.com/AlecAivazis/survey/v2"

var GfVueAdminConfigurationFile = []*survey.Question{
	{
		Name: "GfVueAdminConfigurationFile",
		Prompt: &survey.Select{
			Message: "请选择配置文件!",
			Options: []string{"config/config.mysql.yaml", "config/config.postgres.yaml", "config/config.sqlite.yaml", "config/config.sqlserver.yaml", "config/config.docker-compose.yaml"},
			Default: "config/config.postgres.yaml",
		},
	},
}