package data

import (
	"github.com/flipped-aurora/gva/global"
	model "github.com/flipped-aurora/gva/model/gf/workflow"

	"time"

	"gorm.io/gorm"
)

var Process = new(_process)

type _process struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: workflow_processes 表数据初始化
func (p *_process) Init() error {
	processes := []model.WorkflowProcess{
		{ID: "leaveFlow", CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: "leaveFlow", Clazz: "process", Label: global.I18n.T("LeaveProcess"), HideIcon: false, Description: global.I18n.T("LeaveProcess"), View: "view/iconList/index.vue"},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&processes).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 定义表名
func (p *_process) TableName() string {
	return "workflow_processes"
}
