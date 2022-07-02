package db

import (
	"fmt"
	"sorteo/model"
)

func Validation(ip string) bool {
	var ipUser model.User
	ipUser.Ip = ""
	registry, err := ConDB.Query(`SELECT * FROM direcction where ip= '` + ip + `'`)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	for registry.Next() {

		err := registry.Scan(&ipUser.Ip)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(ipUser.Ip, "ip hh")

	}
	if ipUser.Ip == "" {
		fmt.Println("ip no encontrado")
		CreateIp(ip)
		return false

	}
	fmt.Println("el ip encontrado es ", ipUser.Ip, " !!!")

	return true

}
