package main

import (
	"log"
	"net/http"
	"os"
	"sorteo/router"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"X-Requested-with", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origin := handlers.AllowedOrigins([]string{"*"})
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8084"
	}
	r.HandleFunc("/cupon", router.Coupon).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+PORT, handlers.CORS(headers, methods, origin)(r)))
}
