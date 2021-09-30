//+build mysql,gfva

package boot

import (
	"fmt"
	"github.com/flipped-aurora/gva/library/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
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
		DSN:                       global.GinVueAdminConfig.Gorm.Dsn.GetDsn(), // DSN data source name
		DefaultStringSize:         191,                                        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                      // 根据版本自动配置
	}), Gorm.GenerateConfig())
	if err != nil {
		return err
	}
	global.Db = db
	return nil
}

func (m *_mysql) CheckDatabase() {
	unknownDatabase := fmt.Sprintf("Unknown database '%v'", global.GinVueAdminConfig.Gorm.Dsn.GetDatabase())

}