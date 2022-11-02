package infrastructure

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Println(err)
		log.Fatal("unable to load app.env")
	}
}
