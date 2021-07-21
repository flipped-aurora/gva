package config

type Snowflake struct {
	StartTime string `mapstructure:"start-time" json:"startTime" yaml:"start-time"`
	MachineId uint16 `mapstructure:"machine-id" json:"machineId" yaml:"machine-id"`
}
