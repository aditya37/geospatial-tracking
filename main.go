package main

import (
	"log"

	"github.com/aditya37/geospatial-tracking/service"
)

// func init() {
// 	godotenv.Load()
// }

func main() {
	svc, err := service.NewGrpc()
	if err != nil {
		log.Fatal(err)
	}
	svc.Run()
}
