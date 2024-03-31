package main

import (
	"github.com/abbasfisal/golang-ecommerce/cmd/api"
	"log"
)

func main() {
	server := api.NewAPIServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
