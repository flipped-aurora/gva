package global

import (
	"database/sql"
	"fmt"
	"github.com/gogf/gf/i18n/gi18n"
	"gorm.io/gorm"
	"time"
)

var (
	Db     *gorm.DB
	I18n   *gi18n.Manager
	Config System
)

type Model struct {
	ID        uint           `orm:"id" json:"ID" gorm:"primarykey"`
	CreatedAt time.Time      `orm:"created_at" json:"CreatedAt"`
	UpdatedAt time.Time      `orm:"updated_at" json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `orm:"deleted_at" json:"-" gorm:"index"`
}

type System struct {
	DbType    string `json:"db_type"`
	Language  string `json:"language"`
	FrameType string `json:"frame_type"`

	Host     string `json:"host"`
	Port     string `json:"port"`
	Config   string `json:"config"`
	DbName   string `json:"db_name"`
	Username string `json:"username"`
	Password string `json:"password"`

	LogMode bool `json:"log_mode"`
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取dsn
func (s *System) Dsn() string {
	if s.Config == "" {
		s.Config = "charset=utf8mb4&parseTime=True&loc=Local"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", s.Username, s.Password, s.Host, s.Host, s.DbName, s.Config)
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建数据库(mysql)
func (s *System) CreateTable() error {
	_sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", s.DbName)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", s.Username, s.Password, s.Host, s.Port)
	db, err := sql.Open(s.DbType, dsn)
	if err != nil {
		return err
	}
	defer func() {
		_ = db.Close()
	}()
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(_sql)
	return err
}
