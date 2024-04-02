package main

import (
	"cleverhouse-go-back/config"
	"fmt"
	"log"

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

	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatalf("Db connection error: %s", err)
	}

	defer sess.Close()

	fmt.Println(sess)
}
