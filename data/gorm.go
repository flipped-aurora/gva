package data

import (
	"github.com/flipped-aurora/gva/global"
	"github.com/gookit/color"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"

	// "gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"os"
)

var err error

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() (db *gorm.DB) {
	switch global.Config.System.DbType {
	case "mysql":
		db = GormMysql()
	case "postgresql":
		GormPostgreSql()
	//case "sqlite":
	//	GormSqlite()
	case "sqlserver":
		GormSqlServer()
	}
	return
}

// GormMysql 初始化Mysql数据库
func GormMysql() (db *gorm.DB) {
	m := global.Config.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	gormConfig := config(m.LogMode)
	if db, err = gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
		color.Warn.Printf("MySQL启动异常, err:%v\n", err)
		os.Exit(0)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	}
	return db
}

// GormPostgreSql 初始化PostgreSql数据库
func GormPostgreSql() {
	p := global.Config.Postgresql
	dsn := "user=" + p.Username + " password=" + p.Password + " dbname=" + p.Dbname + " port=" + p.Port + " " + p.Config
	postgresConfig := postgres.Config{
		DSN:                  dsn,                    // DSN data source name
		PreferSimpleProtocol: p.PreferSimpleProtocol, // 禁用隐式 prepared statement
	}
	gormConfig := config(p.Logger)
	if global.Db, err = gorm.Open(postgres.New(postgresConfig), gormConfig); err != nil {
		color.Warn.Printf("PostgreSql启动异常, err:%v\n", err)
		os.Exit(0)
	} else {
		sqlDB, _ := global.Db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
	}
}

// GormSqlite 初始化Sqlite数据库 sqlite需要gcc支持 windows用户需要自行安装gcc 如需使用打开注释即可
//func GormSqlite() {
//	s := global.Config.Sqlite
//	gormConfig := config(s.Logger)
//	if global.Db, err = gorm.Open(sqlite.Open(s.Path), gormConfig); err != nil {
//		color.Warn.Printf("Sqlite启动异常, err:%v\n", err)
//		os.Exit(0)
//	} else {
//		sqlDB, _ := global.Db.DB()
//		sqlDB.SetMaxIdleConns(s.MaxIdleConns)
//		sqlDB.SetMaxOpenConns(s.MaxOpenConns)
//	}
//}

// GormSqlServer 初始化SqlServer数据库
func GormSqlServer() {
	ss := global.Config.Sqlserver
	dsn := "sqlserver://" + ss.Username + ":" + ss.Password + "@" + ss.Path + "?database=gorm"
	if global.Db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{}); err != nil {
		color.Warn.Printf("SqlServer启动异常, err:%v\n", err)
		os.Exit(0)
	} else {
		sqlDB, _ := global.Db.DB()
		sqlDB.SetMaxIdleConns(ss.MaxIdleConns)
		sqlDB.SetMaxOpenConns(ss.MaxOpenConns)
	}
}

// config 根据配置决定是否开启日志
func config(mod bool) (c *gorm.Config) {
	if mod {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	} else {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}
	return
}
