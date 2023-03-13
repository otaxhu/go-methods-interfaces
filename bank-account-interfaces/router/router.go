package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/otaxhu/go-methods-interfaces/bank-account-interfaces/controllers"
	"github.com/otaxhu/go-methods-interfaces/bank-account-interfaces/database"
	"github.com/otaxhu/go-methods-interfaces/bank-account-interfaces/models"
)

func newServer(handler http.Handler) *http.Server {
	return &http.Server{Addr: os.Getenv("server_port"), Handler: handler}
}

func InitializeRouter() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	tx, err := database.NewTransaction()
	if err != nil {
		log.Fatal(err)
	}
	if err := tx.AutoMigrate(&models.User{}, &models.BankAccount{}); err != nil {
		log.Fatal(err)
	}
	tx.Commit()
	router := mux.NewRouter()
	server := newServer(router)

	router.HandleFunc("/", controllers.Home).Methods(http.MethodGet)
	router.HandleFunc("/signup", controllers.RegisterUser).Methods(http.MethodPost)

	log.Fatal(server.ListenAndServe())
}
