package main

import (
	"github.com/abbasfisal/golang-ecommerce/config"
	"github.com/abbasfisal/golang-ecommerce/db"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	_, err := db.NewMysqlStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Net:                  "tcp",
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		ParseTime:            true,
		AllowNativePasswords: true,
	})

	if err != nil {
		log.Fatal(err)
	}
}
