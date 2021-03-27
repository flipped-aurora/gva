package model

import (
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"github.com/flipped-aurora/gva/global"
	extra "github.com/flipped-aurora/gva/model/gf/extra"
	system "github.com/flipped-aurora/gva/model/gf/system"
	workflow "github.com/flipped-aurora/gva/model/gf/workflow"
	"github.com/gookit/color"
	"gorm.io/gorm"
	"os"
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: gorm 同步模型 生成mysql表
func GfVueAdminAutoMigrateTables(db *gorm.DB) {
	err := db.AutoMigrate(
		new(system.Api),
		new(system.Admin),
		new(system.Menu),
		new(gormAdapter.CasbinRule),
		new(system.JwtBlacklist),
		new(system.Authority),
		new(system.Dictionary),
		new(system.OperationRecord),
		new(system.DictionaryDetail),
		new(system.MenuParameter),

		new(extra.File),
		new(extra.SimpleUploader),
		new(extra.BreakpointContinue),
		new(extra.BreakpointContinueChunk),

		new(workflow.WorkflowNode),
		new(workflow.WorkflowEdge),
		new(workflow.WorkflowProcess),
		new(workflow.WorkflowEndPoint),
		new(workflow.WorkflowStartPoint),
	)
	if err != nil {
		color.Warn.Printf("[%v] --> 初始化数据表失败, err: %v\n", global.Config.DbType, err)
		os.Exit(0)
	}
	color.Info.Printf("[%v] --> 初始化数据表成功!\n", global.Config.DbType)
}
