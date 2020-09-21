package mysql

import (
	"gva/data/datas"
	"gva/data/model"
	"gva/global"
	"os"

	"gorm.io/gorm"

	"go.uber.org/zap"
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
		global.Log.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.Log.Info("register table success")
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
		global.Log.Error("initialize data failed", zap.Any("err", err))
	}
	global.Log.Info("initialize data success")
}
