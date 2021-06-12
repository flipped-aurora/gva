package interfaces

import (
	"github.com/gookit/color"
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
