// +build mysql

package interfaces

import "github.com/gookit/color"

func DataInitialize(inits ...InitData) error {
	for i := 0; i < len(inits); i++ {
		if inits[i].TableName() == "system_authority_menu" {
			if k := inits[i].CheckDataExist(); k {
				color.Info.Printf(AuthorityMenu, "mysql", inits[i].TableName())
				continue
			}
		} else {
			if inits[i].CheckDataExist() {
				color.Info.Printf(InitDataExist, "mysql", inits[i].TableName())
				continue
			}
		}

		if err := inits[i].Initialize(); err != nil {
			color.Info.Printf(InitDataFailed, "mysql", err)
			continue
		} else {
			color.Info.Printf(InitDataSuccess, "mysql", inits[i].TableName())
		}
	}
	color.Info.Printf(InitSuccess, "mysql")
	return nil
}
