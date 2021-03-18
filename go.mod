module github.com/flipped-aurora/gva

go 1.16

require (
	github.com/casbin/gorm-adapter/v3 v3.0.2
	github.com/flipped-aurora/tool/cmd/gva v0.0.0-20210304080206-f5a2786e21c1
	github.com/fsnotify/fsnotify v1.4.9
	github.com/go-redis/redis/v8 v8.7.1
	github.com/gookit/color v1.3.8
	github.com/mitchellh/go-homedir v1.1.0
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	golang.org/x/crypto v0.0.0-20210317152858-513c2a44f670
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.3
)

replace github.com/casbin/gorm-adapter/v3 => github.com/casbin/gorm-adapter/v3 v3.0.2
