// go:build postgres
// +build postgres

package gfva

import (
	"database/sql"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	boot "github.com/flipped-aurora/gf-vue-admin/boot/gorm"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/flipped-aurora/gva/answer"
	"github.com/flipped-aurora/gva/question"
	"github.com/gookit/color"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DbResolver = new(_postgres)

type _postgres struct {
	resolver
}

func (p *_postgres) GetConfigPath() string {
	return "config/config.mysql.yaml"
}

func (p *_postgres) CreateDatabase() error {
	_sql := "CREATE DATABASE " + global.Config.Gorm.Dsn.GetDefaultDbName()
	db, err := sql.Open("postgres", global.Config.Gorm.Dsn.GetEmptyDsn())
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

func (p *_postgres) LinkDatabase() error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  global.Config.Gorm.Dsn.GetDefaultDsn(global.Config.Gorm.Config),
		PreferSimpleProtocol: false,
	}), boot.Gorm.GenerateConfig())
	if err != nil {
		_err := errors.New(fmt.Sprintf(fmt.Sprintf("failed to connect to `host=127.0.0.1 user=root database=%v`: server error (FATAL: database \"%v\" does not exist (SQLSTATE 3D000))", global.Config.Gorm.Dsn.GetDefaultDbName(), global.Config.Gorm.Dsn.GetDefaultDbName())))
		if errors.As(err, &_err) {
			input := answer.Database{}
			if err = survey.Ask(question.Database, &input); err != nil {
				color.Warn.Printf("[postgres] --> 获取用户输入失败! error:%v\n", err)
				return err
			}
			switch input.Database {
			case "创建数据库":
				err = p.CreateDatabase()
				if err != nil {
					return err
				}
			case "自行创建", "退出程序":
				os.Exit(0)
			}
		}
		return p.LinkDatabase()
	}
	global.Db = db
	return nil
}
