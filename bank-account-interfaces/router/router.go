package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/otaxhu/go-methods-interfaces/bank-account-interfaces/controllers"
)

func newServer(handler http.Handler) *http.Server {
	return &http.Server{Addr: ":8080", Handler: handler}
}

func InitializeRouter() {
	router := mux.NewRouter()
	server := newServer(router)

	router.HandleFunc("/", controllers.Home).Methods(http.MethodGet)

	log.Fatal(server.ListenAndServe())
}
