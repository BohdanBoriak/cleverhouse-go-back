package container

import (
	"cleverhouse-go-back/config"
	"cleverhouse-go-back/database"
	"log"

	"github.com/upper/db/v4/adapter/postgresql"
)

type Container struct {
	UserRepo database.UserRepository
}

func New() Container {
	config := config.GetConfig()

	settings := postgresql.ConnectionURL{
		Database: config.DbName,
		Host:     config.DbHost,
		User:     config.DbUser,
		Password: config.DbPassword,
	}

	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatalf("Db connection error: %s", err)
	}

	ur := database.NewUserRepository(sess)
	return Container{UserRepo: ur}
}
