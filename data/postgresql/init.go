package postgresql

import (
	"gva/data/datas"
	"gva/data/model"
	"os"

	"github.com/gookit/color"

	gormadapter "github.com/casbin/gorm-adapter/v3"

	"gorm.io/gorm"
)

func InitPostgresqlData(db *gorm.DB) {
	var err error
	err = datas.InitSysApi(db)
	err = datas.InitSysUser(db)
	err = datas.InitExaCustomer(db)
	err = datas.InitCasbinModel(db)
	err = datas.InitSysAuthority(db)
	err = datas.InitSysBaseMenus(db)
	err = datas.InitAuthorityMenu(db)
	err = datas.InitSysAuthorityMenus(db)
	err = datas.InitSysDataAuthorityId(db)
	err = datas.InitSysDictionaryDetail(db)
	err = datas.InitExaFileUploadAndDownload(db)
	err = InitSysDictionaryToPostgresql(db)
	if err != nil {
		color.Error.Printf("[Postgresql]-->初始化数据失败,err: %v\n", err)
		os.Exit(0)
	}
	color.Info.Println("[Postgresql]-->初始化数据成功")
}

func InitPostgresqlTables(db *gorm.DB) {
	var err error
	if !db.Migrator().HasTable("casbin_rule") {
		err = db.Migrator().CreateTable(&gormadapter.CasbinRule{})
	}
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
		model.ExaSimpleUploader{},
		model.SysOperationRecord{},
		model.SysWorkflowStepInfo{},
		model.SysDictionaryDetail{},
		model.SysBaseMenuParameter{},
		model.ExaFileUploadAndDownload{},
		model.SysDictionaryToPostgresql{},
	)
	if err != nil {
		color.Error.Printf("[Postgresql]-->初始化数据表失败,err: %v\n", err)
		os.Exit(0)
	}
	color.Info.Println("[Postgresql]-->初始化数据表成功")
}
