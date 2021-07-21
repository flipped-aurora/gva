//+build sqlite

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
	DbName      string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	IsTemporary bool   `mapstructure:"is-temporary" json:"isTemporary" yaml:"is-temporary"`
}

func (s *Source) IsEmpty() bool {
	if s.DbName == "" {
		return true
	}
	return false
}

func (s *Source) GetDsn() string {
	if s.IsTemporary {
		return "file:" + s.DbName + "?mode=memory&cache=shared"
	}
	return s.DbName
}

type Replica struct {
	DbName      string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	IsTemporary bool   `mapstructure:"is-temporary" json:"isTemporary" yaml:"is-temporary"`
}

func (r *Replica) IsEmpty() bool {
	if r.DbName == "" {
		return true
	}
	return false
}

func (r *Replica) GetDsn() string {
	if r.IsTemporary {
		return "file:" + r.DbName + "?mode=memory&cache=shared"
	}
	return r.DbName
}
