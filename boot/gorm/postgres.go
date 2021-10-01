//+build postgres

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
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DbResolver = new(_postgres)

type _postgres struct{}

func (p *_postgres) AutoMigrate() error {
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

func (p *_postgres) LinkDatabase() error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  global.GinVueAdminConfig.Gorm.GetDsn(),
		PreferSimpleProtocol: false,
	}), Gorm.GenerateConfig())
	if err != nil {
		_err := errors.New(fmt.Sprintf("failed to connect to `host=%v user=%v database=%v`: server error (FATAL: database \"%v\" does not exist (SQLSTATE 3D000))", global.GinVueAdminConfig.Gorm.GetHost(), global.GinVueAdminConfig.Gorm.GetUsername(), global.GinVueAdminConfig.Gorm.GetDbName(), global.GinVueAdminConfig.Gorm.GetDbName()))
		if errors.As(err, &_err) {
			input := answer.Database{}
			if err = survey.Ask(question.Database, &input); err != nil {
				color.Warn.Printf("[mysql] --> 获取用户输入失败! error:%v\n", err)
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
			return p.LinkDatabase()
		}
		return err
	}
	global.Db = db
	return nil
}

func (p *_postgres) GetConfigPath() string {
	return "config/config.postgres.yaml"
}

func (p *_postgres) DataInitialize() {
	_ = interfaces.DataInitialize(
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

func (p *_postgres) CreateDatabase() error {
	_sql := fmt.Sprintf("CREATE DATABASE %s", global.GinVueAdminConfig.Gorm.GetDbName())
	db, err := sql.Open("postgres", global.GinVueAdminConfig.Gorm.GetDatabaseDsn())
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	_, err = db.Exec(_sql)
	return err
}
