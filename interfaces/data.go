package interfaces

import (
	"github.com/gookit/color"
)

const (
	InitSuccess     = "\n[%v] --> 初始数据成功!\n"
	AuthorityMenu   = "\n[%v] --> %v 视图已存在!\n"
	InitDataExist   = "\n[%v] --> %v 表的初始数据已存在!\n"
	InitDataFailed  = "\n[%v] --> %v 表初始数据失败! \nerr: %+v\n"
	InitDataSuccess = "\n[%v] --> %v 表初始数据成功!\n"
)

type InitDateFunc interface {
	Init() error
	TableName() string
}

// PostgresInitDb postgres 数据库 批量初始化表数据
// Author: [SliverHorn](https://github.com/SliverHorn)
func PostgresInitDb(inits ...InitDateFunc) error {
	for _, init := range inits {
		if err := init.Init(); err != nil {
			color.Warn.Printf("\n[postgres] --> %v 表初始数据失败, err: %v\n", init.TableName(), err)
			return err
		} else {
			color.Info.Printf("\n[postgres] --> %v 表初始数据成功!\n", init.TableName())
		}
	}
	color.Info.Printf("\n[postgres] --> 初始化数据成功!\n")
	return nil
}

type InitData interface {
	TableName() string
	Initialize() error
	CheckDataExist() bool
}

func DataInitialize(dbType string, inits ...InitData) error {
	for i := 0; i < len(inits); i++ {
		if inits[i].TableName() == "authority_menu" {
			if k := inits[i].CheckDataExist(); k {
				color.Info.Printf(AuthorityMenu, dbType, inits[i].TableName())
				continue
			}
		} else {
			if inits[i].CheckDataExist() {
				color.Info.Printf(InitDataExist, dbType, inits[i].TableName())
				continue
			}
		}

		if err := inits[i].Initialize(); err != nil {
			color.Info.Printf(InitDataFailed, dbType, err)
			continue
		} else{
			color.Info.Printf(InitDataSuccess, dbType, inits[i].TableName())
		}
	}
	color.Info.Printf(InitSuccess, dbType)
	return nil
}
