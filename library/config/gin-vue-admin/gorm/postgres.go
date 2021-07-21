//+build postgres

package config

import "time"

type Dsn struct {
	MaxIdleConnes   int           `mapstructure:"max-idle-connes" json:"maxIdleConnes" yaml:"max-idle-connes"`
	MaxOpenConnes   int           `mapstructure:"max-open-connes" json:"maxOpenConnes" yaml:"max-open-connes"`
	ConnMaxLifetime time.Duration `mapstructure:"conn-max-lifetime" json:"connMaxLifetime" yaml:"conn-max-lifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn-max-idle-time" json:"connMaxIdleTime" yaml:"conn-max-idle-time"`
	Sources         []Source      `mapstructure:"sources" json:"sources" yaml:"sources"`
	Replicas        []Replica     `mapstructure:"replicas" json:"replicas" yaml:"replicas"`
}

type Source struct {
	Host        string `mapstructure:"host" json:"host" yaml:"host"`
	Port        string `mapstructure:"port" json:"port" yaml:"port"`
	DbName      string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username    string `mapstructure:"username" json:"username" yaml:"username"`
	Password    string `mapstructure:"password" json:"password" yaml:"password"`
	OtherConfig string `mapstructure:"other-config" json:"otherConfig" yaml:"other-config"`
}

func (s *Source) IsEmpty() bool {
	if s.Host == "" || s.Port == "" || s.DbName == "" || s.Username == "" || s.Password == "" {
		return true
	}
	return false
}

func (s *Source) GetDsn(config string) string {
	if s.OtherConfig != "" {
		return "host=" + s.Host + " user=" + s.Username + " password=" + s.Password + " dbname=" + s.DbName + " port=" + s.Port + " " + s.OtherConfig
	}
	return "host=" + s.Host + " user=" + s.Username + " password=" + s.Password + " dbname=" + s.DbName + " port=" + s.Port + " " + config
}

type Replica struct {
	Host        string `mapstructure:"host" json:"host" yaml:"host"`
	Port        string `mapstructure:"port" json:"port" yaml:"port"`
	DbName      string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username    string `mapstructure:"username" json:"username" yaml:"username"`
	Password    string `mapstructure:"password" json:"password" yaml:"password"`
	OtherConfig string `mapstructure:"other-config" json:"otherConfig" yaml:"other-config"`
}

func (r *Replica) IsEmpty() bool {
	if r.Host == "" || r.Port == "" || r.DbName == "" || r.Username == "" || r.Password == "" {
		return true
	}
	return false
}

func (r *Replica) GetDsn(config string) string {
	if r.OtherConfig != "" {
		return "host=" + r.Host + " user=" + r.Username + " password=" + r.Password + " dbname=" + r.DbName + " port=" + r.Port + " " + r.OtherConfig
	}
	return "host=" + r.Host + " user=" + r.Username + " password=" + r.Password + " dbname=" + r.DbName + " port=" + r.Port + " " + config
}
