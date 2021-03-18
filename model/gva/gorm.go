package model

import (
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	extra "github.com/flipped-aurora/gva/model/gva/extra"
	system "github.com/flipped-aurora/gva/model/gva/system"
	workflow "github.com/flipped-aurora/gva/model/gva/workflow"
	"github.com/gookit/color"
	"gorm.io/gorm"
	"os"
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: gorm 同步模型 生成mysql表
func GinVueAdminAutoMigrateTables(db *gorm.DB) {
	err := db.AutoMigrate(
		new(system.SysApi),
		new(system.SysUser),
		new(gormAdapter.CasbinRule),
		new(system.SysBaseMenu),
		new(system.JwtBlacklist),
		new(system.SysAuthority),
		new(system.SysDictionary),
		new(system.SysOperationRecord),
		new(system.SysDictionaryDetail),
		new(system.SysBaseMenuParameter),

		new(extra.ExaFile),
		new(extra.ExaCustomer),
		new(extra.ExaFileChunk),
		new(extra.ExaSimpleUploader),
		new(extra.ExaFileUploadAndDownload),

		new(workflow.WorkflowNode),
		new(workflow.WorkflowEdge),
		new(workflow.WorkflowProcess),
		new(workflow.WorkflowEndPoint),
		new(workflow.WorkflowStartPoint),
	)
	if err != nil {
		color.Warn.Printf("[Mysql] --> 初始化数据表失败, err: %v\n", err)
		os.Exit(0)
	}
	color.Info.Println("[Mysql] --> 初始化数据表成功! ")
}
