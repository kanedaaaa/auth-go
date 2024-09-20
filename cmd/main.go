package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/kanedaaaa/auth-go/cmd/api"
	"github.com/kanedaaaa/auth-go/config"
	"github.com/kanedaaaa/auth-go/db"
)

func main() {
	db := db.NewSQLStorage(mysql.Config{
		User:                 config.ENVS.DBUser,
		Passwd:               config.ENVS.DBPassword,
		Addr:                 config.ENVS.DBAddress,
		DBName:               config.ENVS.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	initDB(db)
	fmt.Printf("Running on PORT %v\n", config.ENVS.Port)

	server := api.NewAPIServer(config.ENVS.Port, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initDB(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB Connected")
}
