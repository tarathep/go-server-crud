package main

import (
	"log"

	"github.com/tarathep/go-server-crud/apis"
	"github.com/tarathep/go-server-crud/db"
	"github.com/tarathep/go-server-crud/router"
)

func main() {
	// mongodb://admin:password@10.138.36.166:27017
	db, err := db.Init("mongodb://192.168.88.192:27018")
	if err != nil {
		log.Panic(err)
	}

	route := router.Router{apis.HelloHandler{db}, apis.TutorialHandler{db}}
	route.Route().Run(":8089")
}
