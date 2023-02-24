package main

import (
	"log"

	v1 "github.com/ishanshre/2FA-with-golang/api/v1"
)

func main() {
	store, err := v1.NewPostgreStore()
	if err != nil {
		log.Fatalf("Error in connection in database: %s", err)
	}
	if err := store.Init(); err != nil {
		log.Fatalf("Error in creating table in database: %s", err)
	}
	server := v1.NewApiServer(":8000", store)
	server.Run()
}
