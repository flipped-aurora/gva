package interfaces

import (
	"github.com/flipped-aurora/gva/global"
	"github.com/gookit/color"
)

type InitDateFunc interface {
	Init() error
	TableName() string
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量初始化表数据
func InitDb(inits ...InitDateFunc) error {
	for _, init := range inits {
		if err := init.Init(); err != nil {
			color.Warn.Printf("\n[%v] --> %v 表初始数据失败, err: %v\n", global.Config.DbType, init.TableName(), err)
			return err
		} else {
			color.Info.Printf("\n[%v] --> %v 表初始数据成功!\n", global.Config.DbType, init.TableName())
		}
	}
	color.Info.Printf("\n[%v] --> 初始化数据成功!\n", global.Config.DbType)
	return nil
}
