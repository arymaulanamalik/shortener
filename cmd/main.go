package main

import (
	"log"

	"github.com/arymaulanamalik/shortener/config"
	"github.com/arymaulanamalik/shortener/database"
)

//

func main() {
	conf := config.NewConfig()

	//connection database
	log.Println("connecting to database...")
	_, err := database.NewDatabase(*conf)
	if err != nil {
		panic(err)
	}
	log.Println("connecting to database success")
}
