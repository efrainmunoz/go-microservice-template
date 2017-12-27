package routers

import (
	"github.com/gorilla/mux"
)

// InitRoutes initializes the services routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// Routes for the User entity
	router = SetUserRoutes(router)
	return router
}
