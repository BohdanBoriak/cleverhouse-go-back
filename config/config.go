package config

type Config struct {
	DbName     string
	DbHost     string
	DbUser     string
	DbPassword string
}

func GetConfig() Config {
	return Config{
		DbName:     "cleverhouse-db",
		DbHost:     "localhost",
		DbUser:     "postgres",
		DbPassword: "postgres",
	}
}
