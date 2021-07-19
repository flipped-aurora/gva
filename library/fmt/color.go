package fmt

import "github.com/gookit/color"

const (
	InitDataExist   = "\n[%v] --> %v 表的初始数据已存在!"
	InitDataSuccess = "\n[%v] --> %v 表初始数据成功!"
)

// Printf
// Author [SliverHorn](https://github.com/SliverHorn)
func Printf(format string, database string, tableName string) {
	color.Info.Printf(format, database, tableName)
}
