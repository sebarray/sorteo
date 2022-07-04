package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sorteo/db"
	"sorteo/model"
	"strings"
)

func Coupon(w http.ResponseWriter, r *http.Request) {
	var ip model.User
	direcction := strings.Split(r.RemoteAddr, ":")[0]

	reqbody, err := ioutil.ReadAll(r.Body) //recibo datos que envia el cliente
	if err != nil {
		fmt.Fprintln(w, "error al recibir datos del cliente")
	}
	json.Unmarshal(reqbody, &ip)
	if ip.Ip == "" {
		json.NewEncoder(w).Encode("error")
	} else {

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)

		if db.Validation(direcction) {
			json.NewEncoder(w).Encode("creo que ya pidio un cupon")
		} else {
			_, P := db.Asignar(direcction)
			json.NewEncoder(w).Encode(P)
		}
	}
}
