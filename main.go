package main

import (
	"fmt"
	"log"

	"github.com/riphidon/evo/config"
	"github.com/riphidon/evo/router"
	"github.com/riphidon/evo/server"
)

func main() {
	var err error

	// Get the configuration data
	data, err := config.Data.ParseConfigFile()
	if err != nil {
		fmt.Println("can't connect")
		return
	}

	//Connect to Database
	// db.InitDB()

	r := router.InitRouter()

	addr := data.ServerPort

	srv := server.New(r, addr)

	router.SetupRoutes(r)

	err = srv.ListenAndServe()
	log.Fatal(err)
}
