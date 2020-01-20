package main

import (
	"log"

	"github.com/riphidon/evo/config"
	"github.com/riphidon/evo/router"
	"github.com/riphidon/evo/server"
)

func main() {
	var err error

	// Get the configuration data
	config.Data.ParseConfigFile()

	//Connect to Database
	// db.InitDB()

	r := router.InitRouter()

	addr := config.Data.ServerPort

	srv := server.New(r, addr)

	router.SetupRoutes(r)

	err = srv.ListenAndServe()
	log.Fatal(err)
}
