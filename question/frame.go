package question

import "github.com/AlecAivazis/survey/v2"

var Frame = []*survey.Question{
	{
		Name: "frame",
		Prompt: &survey.Select{
			Message: "请选择框架",
			Options: []string{"gf-vue-admin", "gin-vue-admin", "gin-vue-admin-business"},
			Default: "gf-vue-admin",
		},
	},
}
