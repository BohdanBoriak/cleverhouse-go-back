package main

import (
	"cleverhouse-go-back/config"

	"github.com/upper/db/v4/adapter/postgresql"
)

func main() {
	config := config.GetConfig()

	settings := postgresql.ConnectionURL{
		Database: config.DbName,
		Host:     config.DbHost,
		User:     config.DbUser,
		Password: config.DbPassword,
	}
}
