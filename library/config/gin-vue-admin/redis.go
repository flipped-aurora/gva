package config

import "time"

type Redis struct {
	DB         int           `mapstructure:"db" json:"db" yaml:"db"`
	Expire     time.Duration `mapstructure:"expire" json:"expire" yaml:"expire"` // 过期时间
	Address    string        `mapstructure:"address" json:"address" yaml:"address"`
	Password   string        `mapstructure:"password" json:"password" yaml:"password"`
	ExpireNil  time.Duration `mapstructure:"expire-nil" json:"expireNil" yaml:"expire-nil"` // 数据为空的过期时间
	Serializer string        `mapstructure:"serializer" json:"serializer" yaml:"serializer"` // 序列化方式
}
