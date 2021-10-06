//+build mysql

package gfva

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/boot/gorm"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/flipped-aurora/gva/answer"
	data "github.com/flipped-aurora/gva/data/mysql/gf-vue-admin/system"
	"github.com/flipped-aurora/gva/interfaces"
	"github.com/flipped-aurora/gva/question"
	"github.com/gookit/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DbResolver = new(_mysql)

type _mysql struct{}

func (m *_mysql) GetConfigPath() string {
	return "config/config.mysql.yaml"
}

func (m *_mysql) LinkDatabase() error {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       global.Config.Gorm.Dsn.GetDefaultDsn(global.Config.Gorm.Config), // DSN data source name
		DefaultStringSize:         191,                                                             // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                            // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                            // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                            // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                           // 根据版本自动配置
	}), boot.Gorm.GenerateConfig())
	if err != nil {
		_err := errors.New(fmt.Sprintf("Unknown database '%v'", global.Config.Gorm.Dsn.GetDefaultDbName()))
		if errors.As(err, &_err) {
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
		return m.LinkDatabase()
	}
	global.Db = db
	return nil
}

func (m *_mysql) CreateDatabase() error {
	_sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", global.Config.Gorm.Dsn.GetDefaultDbName())
	db, err := sql.Open("mysql", global.Config.Gorm.Dsn.GetEmptyDsn())
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

func (m *_mysql) DataInitialize() {
	_ = interfaces.DataInitialize(
		data.Api,
		data.User,
		data.Menu,
		data.Casbin,
		data.Authority,
		data.Dictionary,
		data.DictionaryDetail,
		data.AuthoritiesMenus,
		data.AuthoritiesResources,
		data.AuthorityMenu,
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
