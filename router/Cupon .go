package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sorteo/db"
	"sorteo/model"
)

func Coupon(w http.ResponseWriter, r *http.Request) {
	var ip model.User
	fmt.Println(r.RemoteAddr)
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

		if db.Validation(r.RemoteAddr) {
			json.NewEncoder(w).Encode("ccreo que ya pidio un cupon")
		} else {
			_, P := db.Asignar(r.RemoteAddr)
			json.NewEncoder(w).Encode(P)
		}
	}
}
