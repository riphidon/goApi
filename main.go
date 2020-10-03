package main

import (
	"log"

	"github.com/riphidon/evo/config"
	"github.com/riphidon/evo/router"
	"github.com/riphidon/evo/server"
	"github.com/rs/cors"
)

func main() {
	var err error

	// Get the configuration data
	config.Data.ParseConfigFile()

	//Connect to Database
	// db.InitDB()

	addr := config.Data.ServerPort

	r := router.InitRouter()
	handler := cors.Default().Handler(r)
	srv := server.New(handler, addr)
	router.SetupRoutes(r)

	err = srv.ListenAndServe()
	log.Fatal(err)
}
