package interfaces

type Boot interface {
	Initialize()
}

type GormData interface {
	AutoMigrate() error
	LinkDatabase() error
	GetConfigPath() string
	DataInitialize()
	CreateDatabase() error
}
