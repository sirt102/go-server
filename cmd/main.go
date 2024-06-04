package main

import (
	"go-server/config"
	"go-server/internal/infrastructure/router"
	"go-server/internal/registry"
	"go-server/pkg/mongo"
)

func main() {
	config.LoadConfig()
	mongoDB := mongo.NewMongo(config.C.MongoDB.URLString, config.C.MongoDB.DatabaseName)
	rg := registry.NewInteractor(mongoDB)

	masterHandler := rg.NewAppHandler()

	router.Initialize(masterHandler)
}
