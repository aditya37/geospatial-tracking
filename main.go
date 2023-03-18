package main

import (
	"log"
	"os"

	"github.com/aditya37/geospatial-tracking/service"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	log.Println(os.Getenv("ELASTIC_APM_SERVICE_NAME"))
	log.Println(os.Getenv("ELASTIC_APM_SERVER_URL"))

	svc, err := service.NewGrpc()
	if err != nil {
		log.Fatal(err)
	}
	svc.Run()
}
