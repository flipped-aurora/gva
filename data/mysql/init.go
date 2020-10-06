package mysql

import (
	"os"

	"github.com/flipped-aurora/gva-ctl/data/datas"
	"github.com/flipped-aurora/gva-ctl/data/model"

	"github.com/gookit/color"
	"gorm.io/gorm"
)

func InitMysqlTables(db *gorm.DB) {
	var err error
	err = db.AutoMigrate(
		model.SysApi{},
		model.SysUser{},
		model.ExaFile{},
		model.ExaCustomer{},
		model.SysBaseMenu{},
		model.SysWorkflow{},
		model.SysAuthority{},
		model.JwtBlacklist{},
		model.ExaFileChunk{},
		model.SysDictionary{},
		model.ExaSimpleUploader{},
		model.SysOperationRecord{},
		model.SysWorkflowStepInfo{},
		model.SysDictionaryDetail{},
		model.SysBaseMenuParameter{},
		model.ExaFileUploadAndDownload{},
	)
	if err != nil {
		color.Warn.Printf("[Mysql]-->初始化数据表失败,err: %v\n", err)
		os.Exit(0)
	}
	color.Info.Println("[Mysql]-->初始化数据表成功")
}

func InitMysqlData(db *gorm.DB) {
	var err error
	err = datas.InitSysApi(db)
	err = datas.InitSysUser(db)
	err = datas.InitExaCustomer(db)
	err = datas.InitCasbinModel(db)
	err = datas.InitSysAuthority(db)
	err = datas.InitSysBaseMenus(db)
	err = datas.InitAuthorityMenu(db)
	err = datas.InitSysDictionary(db)
	err = datas.InitSysAuthorityMenus(db)
	err = datas.InitSysDataAuthorityId(db)
	err = datas.InitSysDictionaryDetail(db)
	err = datas.InitExaFileUploadAndDownload(db)
	if err != nil {
		color.Warn.Printf("[Mysql]-->初始化数据失败,err: %v\n", err)
		os.Exit(0)
	}
	color.Info.Println("[Mysql]-->初始化数据成功")
}
