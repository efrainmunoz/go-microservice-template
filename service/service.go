package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/efrainmunoz/go-microservice-template/service/common"
	"github.com/efrainmunoz/go-microservice-template/service/routers"
)

// Entry point of the service
func main() {
	// call startup logic
	common.StartUp()

	// get the mux router object
	router := routers.InitRoutes()

	// create negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
