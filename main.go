package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"X-Requested-with", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origin := handlers.AllowedOrigins([]string{"*"})
	PORT := os.Getenv("PORT")

	//router.HandleFunc("/", met.Indexrouter).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+PORT, handlers.CORS(headers, methods, origin)(router)))
}
