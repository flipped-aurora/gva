// go:build postgres
// +build postgres

package interfaces

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/gookit/color"
)

func DataInitialize(inits ...InitData) error {
	var entity system.AuthorityMenu
	for i := 0; i < len(inits); i++ {
		if inits[i].TableName() == entity.TableName() {
			if k := inits[i].CheckDataExist(); k {
				color.Info.Printf(AuthorityMenu, "postgres", inits[i].TableName())
				continue
			}
		} else {
			if inits[i].CheckDataExist() {
				color.Info.Printf(InitDataExist, "postgres", inits[i].TableName())
				continue
			}
		}

		if err := inits[i].Initialize(); err != nil {
			color.Info.Printf(InitDataFailed, "postgres", err)
			continue
		} else {
			color.Info.Printf(InitDataSuccess, "postgres", inits[i].TableName())
		}
	}
	color.Info.Printf(InitSuccess, "postgres")
	return nil
}
