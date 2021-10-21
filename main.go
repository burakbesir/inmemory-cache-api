package main

import (
	"github.com/rahmanbesir/inmemory-cache-api/controller"
	"github.com/rahmanbesir/inmemory-cache-api/service"
	"log"
	"net/http"
)

func main() {
	cacheService := service.NewCacheService()
	contr := controller.New(cacheService)

	http.HandleFunc("/cache", contr.Handle)

	log.Fatalln(http.ListenAndServe(":8080", nil))

}
