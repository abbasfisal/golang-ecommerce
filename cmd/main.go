package main

import (
	"database/sql"
	"github.com/abbasfisal/golang-ecommerce/cmd/api"
	"github.com/abbasfisal/golang-ecommerce/config"
	"github.com/abbasfisal/golang-ecommerce/db"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	db, err := db.NewMysqlStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Net:                  "tcp",
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		ParseTime:            true,
		AllowNativePasswords: true,
	})

	CheckDb(db)

	if err != nil {
		log.Fatal(err)
	}
	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func CheckDb(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Db: successfully connected ")
}
