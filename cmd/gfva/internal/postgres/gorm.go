package postgres

import "github.com/flipped-aurora/gf-vue-admin/library/global"

var Postgres = new(_postgres)

type _postgres struct{}

func (p *_postgres) CreateDatabase() error {
	return global.Db.Exec("CREATE DATABASE " + global.Config.Gorm.Dsn.GetDefaultDbName()).Error
}
