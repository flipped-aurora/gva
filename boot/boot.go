package boot

func Initialize(path ...string) {
	Viper.Initialize(path...)
	DbResolver.Initialize()
	Zap.Initialize()
}
