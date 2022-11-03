package service

import (
	"database/sql"
	"foobar-go/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var store infrastructure.Store
var eventService EventService

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

	eventService = NewEventService(store)

	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
