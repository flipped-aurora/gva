package config

import "time"

type JWT struct {
	Issuer      string        `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
	NotBefore   int64         `mapstructure:"not-before" json:"notBefore" yaml:"not-before"`
	ExpiresAt   int64         `mapstructure:"expires-at" json:"expiresAt" yaml:"expires-at"`
	BufferTime  int64         `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`
	SigningKey  string        `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`
	ExpireRedis time.Duration `mapstructure:"expire-redis" json:"expireRedis" yaml:"issuer"`
}
