module github.com/flipped-aurora/tool/cmd/gva

go 1.11

require (
	github.com/casbin/gorm-adapter/v3 v3.2.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gookit/color v1.3.8
	github.com/mitchellh/go-homedir v1.1.0
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	gorm.io/driver/mysql v1.0.4
	gorm.io/gorm v1.20.12
)

replace github.com/casbin/gorm-adapter/v3 => github.com/casbin/gorm-adapter/v3 v3.0.2
