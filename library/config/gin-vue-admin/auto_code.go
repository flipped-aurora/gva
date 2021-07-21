package config

import (
	"path/filepath"
)

type AutoCode struct {
	Web             Web    `mapstructure:"web" json:"web" yaml:"web"`
	Root            string `mapstructure:"root" json:"root" yaml:"root"`
	Server          Server `mapstructure:"server" json:"server" yaml:"server"`
	TransferRestart bool   `mapstructure:"transfer-restart" json:"transferRestart" yaml:"transfer-restart"`
}

type Server struct {
	Api     string `mapstructure:"api" json:"api" yaml:"api"`
	Boot    string `mapstructure:"boot" json:"boot" yaml:"boot"`
	Model   string `mapstructure:"model" json:"model" yaml:"model"`
	Router  string `mapstructure:"router" json:"router" yaml:"router"`
	Server  string `mapstructure:"server" json:"server" yaml:"server"`
	Request string `mapstructure:"request" json:"request" yaml:"request"`
	Service string `mapstructure:"service" json:"service" yaml:"service"`
}

type Web struct {
	Web   string `mapstructure:"web" json:"web" yaml:"web"`
	Api   string `mapstructure:"api" json:"api" yaml:"api"`
	Form  string `mapstructure:"form" json:"form" yaml:"form"`
	Flow  string `mapstructure:"flow" json:"flow" yaml:"flow"`
	Table string `mapstructure:"table" json:"table" yaml:"table"`
}

func (a *AutoCode) GetGormFilePath() string {
	return filepath.Join(a.Root, a.Server.Server, a.Server.Boot, "gorm.go")
}

