package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Dhruv2638/ecom-go/cmd/api"
	"github.com/Dhruv2638/ecom-go/config"
	"github.com/Dhruv2638/ecom-go/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}) // This is just for opening the connection
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db) // this connectes to DB
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB is Connected!")
}
