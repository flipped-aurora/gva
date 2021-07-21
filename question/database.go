package question

import "github.com/AlecAivazis/survey/v2"

var (
	Database = []*survey.Question{
		{
			Name: "database",
			Prompt: &survey.Select{
				Message: "您的配置文件所配置的数据库不存在!",
				Options: []string{"创建数据库", "自行创建", "退出程序"},
				Default: "创建数据库",
			},
		},
	}
	DatabaseType = []*survey.Question{
		{
			Name: "database",
			Prompt: &survey.Select{
				Message: "您的配置文件所配置的数据库不存在!",
				Options: []string{"mysql", "postgres", "sqlite", "sqlserver"},
				Default: "postgres",
			},
		},
	}
)
