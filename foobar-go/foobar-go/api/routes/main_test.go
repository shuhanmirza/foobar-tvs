package routes

import (
	"database/sql"
	"foobar-go/api/controller"
	"foobar-go/api/service"
	"foobar-go/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var ginRouter infrastructure.GinRouter
var router *gin.Engine

var eventService service.EventService
var eventController controller.EventController
var eventRoute EventRoute
var configurationService service.ConfigurationService
var configurationController controller.ConfigurationController
var configurationRoute ConfigurationRoute

func TestMain(m *testing.M) {

	err := godotenv.Load("../../app.env")
	if err != nil {
		log.Println(err)
		log.Fatal("unable to load app.env")
	}

	DbDriver := os.Getenv("DB_DRIVER")
	DbSource := os.Getenv("DB_SOURCE")
	conn, err := sql.Open(DbDriver, DbSource)
	if err != nil {
		log.Fatal("can't connect to the database", err)
	}

	store := infrastructure.NewStore(conn)
	ginRouter = infrastructure.NewGinRouter()
	//// configuration feature
	configurationService = service.NewConfigurationService(store)
	configurationController = controller.NewConfigurationController(configurationService)
	configurationRoute = NewConfigurationRoute(configurationController, ginRouter)
	configurationRoute.Setup()

	//// Event Feature
	eventService = service.NewEventService(store)
	eventController = controller.NewEventController(eventService)
	eventRoute = NewEventRoute(eventController, ginRouter)
	eventRoute.Setup()

	router = ginRouter.Gin

	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
