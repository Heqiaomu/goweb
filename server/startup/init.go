package startup

func InitServer() {
	InitViper()
	InitConfig()
	InitLogger()
	InitMySQL()

	InitRedis()
}
