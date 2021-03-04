module github.com/flipped-aurora/tool/cmd/gfva

go 1.16

require (
	github.com/casbin/gorm-adapter/v3 v3.0.2
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gookit/color v1.3.8
	github.com/mitchellh/go-homedir v1.1.0
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0 // indirect
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	golang.org/x/sys v0.0.0-20210112080510-489259a85091 // indirect
	golang.org/x/text v0.3.4 // indirect
	gorm.io/driver/mysql v1.0.4
	gorm.io/gorm v1.20.12
)

replace github.com/casbin/gorm-adapter/v3 => github.com/casbin/gorm-adapter/v3 v3.0.2
