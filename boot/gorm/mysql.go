//+build mysql,business

package boot

import (
	"database/sql"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/flipped-aurora/gva/answer"
	"github.com/flipped-aurora/gva/interfaces"
	"github.com/flipped-aurora/gva/library/global"
	system "github.com/flipped-aurora/gva/model/gin-vue-admin/system"
	model "github.com/flipped-aurora/gva/model/gin-vue-admin/system/data"
	"github.com/flipped-aurora/gva/question"
	"github.com/gookit/color"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"os"
)

var DbResolver = new(_mysql)

type _mysql struct {
	dsn string
}

// GetSources 获取主库的 gorm.Dialector 切片对象
// Author [SliverHorn](https://github.com/SliverHorn)
func (m *_mysql) GetSources() (directories []gorm.Dialector) {
	length := len(global.GinVueAdminConfig.Gorm.Dsn.Sources)
	directories = make([]gorm.Dialector, 0, length)
	for i := 0; i < length; i++ {
		if !global.GinVueAdminConfig.Gorm.Dsn.Sources[i].IsEmpty() {
			dsn := global.GinVueAdminConfig.Gorm.Dsn.Sources[i].GetDsn(global.GinVueAdminConfig.Gorm.Config)
			if i == 0 {
				m.dsn = dsn
			}
			directories = append(directories, mysql.Open(dsn))
		} else {
			continue
		}
	}
	return directories
}

// GetReplicas 获取从库库的 gorm.Dialector 切片对象
// Author [SliverHorn](https://github.com/SliverHorn)
func (m *_mysql) GetReplicas() (directories []gorm.Dialector) {
	length := len(global.GinVueAdminConfig.Gorm.Dsn.Replicas)
	directories = make([]gorm.Dialector, 0, length)
	for i := 0; i < length; i++ {
		if !global.GinVueAdminConfig.Gorm.Dsn.Replicas[i].IsEmpty() {
			dsn := global.GinVueAdminConfig.Gorm.Dsn.Replicas[i].GetDsn(global.GinVueAdminConfig.Gorm.Config)
			directories = append(directories, mysql.Open(dsn))
		} else {
			continue
		}
	}
	return directories
}

// GetResolver 通过主库与从库的链接组装 gorm.Plugin
// Author [SliverHorn](https://github.com/SliverHorn)
func (m *_mysql) GetResolver() gorm.Plugin {
	sources := m.GetSources()
	resolver := dbresolver.Register(dbresolver.Config{
		Sources:  sources,
		Replicas: m.GetReplicas(),
		Policy:   dbresolver.RandomPolicy{}, // sources/replicas 负载均衡策略
	})
	resolver.SetMaxIdleConns(global.GinVueAdminConfig.Gorm.GetMaxOpenConnes())
	resolver.SetMaxOpenConns(global.GinVueAdminConfig.Gorm.GetMaxOpenConnes())
	resolver.SetConnMaxIdleTime(global.GinVueAdminConfig.Gorm.GetConnMaxIdleTime())
	resolver.SetConnMaxLifetime(global.GinVueAdminConfig.Gorm.GetConnMaxLifetime())
	return resolver
}

// GetGormDialector 获取数据库的 gorm.Dialector
// Author [SliverHorn](https://github.com/SliverHorn)
func (m *_mysql) GetGormDialector() gorm.Dialector {
	return mysql.New(mysql.Config{
		DSN:                       m.dsn, // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: true,  // 根据版本自动配置
	})
}

func (m *_mysql) GetConfigPath() string {
	return "config/config.mysql.yaml"
}

func (m *_mysql) LinkDatabase() error {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       global.GinVueAdminConfig.Gorm.GetMysqlDsn(), // DSN data source name
		DefaultStringSize:         191,                                         // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                        // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                        // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                        // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                       // 根据版本自动配置
	}), Gorm.GenerateConfig())
	if err != nil {
		if err = m.CheckDatabase(err); err != nil {
			return err
		}
		return m.LinkDatabase()
	}
	global.Db = db
	return nil
}

func (m *_mysql) CreateDatabase() error {
	_sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", global.GinVueAdminConfig.Gorm.GetDbName())
	db, err := sql.Open("driver", global.GinVueAdminConfig.Gorm.GetMysqlDatabaseDsn())
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(_sql)
	return err
}

func (m *_mysql) CheckDatabase(err error) error {
	message := fmt.Sprintf("Unknown database '%v'", global.GinVueAdminConfig.Gorm.GetDbName())
	if errors.As(err, errors.New(message)) {
		input := answer.Database{}
		if err = survey.Ask(question.Database, &input); err != nil {
			color.Warn.Printf("[mysql] --> 获取用户输入失败! error:%v\n", err)
			return err
		}
		switch input.Database {
		case "创建数据库":
			err = m.CreateDatabase()
			if err != nil {
				return err
			}
		case "自行创建", "退出程序":
			os.Exit(0)
		}
	}
	return nil
}

func (m *_mysql) DataInitialize() {
	_ = interfaces.DataInitialize("mysql",
		model.Api,
		model.User,
		model.Menu,
		model.Casbin,
		model.Authority,
		model.Dictionary,
		model.DictionaryDetail,
		model.AuthoritiesMenus,
		model.AuthoritiesResources,
		model.AuthorityMenu,
	)
}

func (m *_mysql) AutoMigrate() error {
	return global.Db.AutoMigrate(
		new(system.Api),
		new(system.Menu),
		new(system.User),
		new(system.Casbin),
		new(system.Authority),
		new(system.Dictionary),
		new(system.JwtBlacklist),
		new(system.MenuParameter),
		new(system.OperationRecord),
		new(system.DictionaryDetail),
	)
}
