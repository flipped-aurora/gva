package library

import (
	"go.uber.org/zap"
	"os"
)

var Directory = new(directory)

type directory struct{}

// PathExists 文件目录是否存在
// Author: [SliverHorn](https://github.com/SliverHorn)
func (d *directory) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CreateDir 批量创建文件夹
// Author: [SliverHorn](https://github.com/SliverHorn)
func (d *directory) CreateDir(dirs ...string) error {
	for _, v := range dirs {
		if exist, err := d.PathExists(v); err != nil {
			return err
		} else {
			if !exist {
				zap.L().Debug("create directory" + v)
				err = os.MkdirAll(v, os.ModePerm)
				if err != nil {
					zap.L().Error("create directory"+v, zap.Any(" error:", err))
				}
			}
		}
	}
	return nil
}
