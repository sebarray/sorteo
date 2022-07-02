package main

import (
	"encoding/json"
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
	r.HandleFunc("/", Try).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+PORT, handlers.CORS(headers, methods, origin)(r)))
}

func Try(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("ccreo que ya pidio un cupon")

}
