package main

import (
	"database/sql"
	"foobar-go/api/controller"
	"foobar-go/api/routes"
	"foobar-go/api/service"
	"foobar-go/infrastructure"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	//env
	log.Println("Loading Environment Variables")
	infrastructure.LoadEnv()

	//database
	log.Println("Initializing Database Connection")
	DbDriver := os.Getenv("DB_DRIVER")
	DbSource := os.Getenv("DB_SOURCE")
	conn, err := sql.Open(DbDriver, DbSource)
	if err != nil {
		log.Fatal("can't connect to the database", err)
	}
	store := infrastructure.NewStore(conn)

	// apis
	log.Println("Initializing Routes")
	ginRouter := infrastructure.NewGinRouter()

	//// configuration feature
	configurationService := service.NewConfigurationService(store)
	configurationController := controller.NewConfigurationController(configurationService)
	configurationRoute := routes.NewConfigurationRoute(configurationController, ginRouter)
	configurationRoute.Setup()

	ServerAddress := os.Getenv("SERVER_ADDRESS")
	err = ginRouter.Gin.Run(ServerAddress)
	if err != nil {
		log.Println(err)
		log.Fatal("could not start APIs")
	}

}
