package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sorteo/model"
)

func Coupon(w http.ResponseWriter, r *http.Request) {
	var ip model.User
	reqbody, err := ioutil.ReadAll(r.Body) //recibo datos que envia el cliente
	if err != nil {
		fmt.Fprintln(w, "error al recibir datos del cliente")
	}
	json.Unmarshal(reqbody, &ip)
	// asigno los valores  recibido al struct
	//newTask.ID = sqlg.Insert(newTask.Name, newTask.Content)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ip)
}
