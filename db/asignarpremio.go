package db

import (
	"fmt"
	"sorteo/model"
)

func Asignar(ip string) (bool, string) {
	var ipUser model.Premio

	registry, err := ConDB.Query(`SELECT * FROM premio WHERE   ip='nil';`)
	if err != nil {
		fmt.Println(err.Error())
		return false, " "
	}

	for registry.Next() {

		err := registry.Scan(&ipUser.Premio, &ipUser.Ip)
		if err != nil {
			fmt.Println(err.Error())

		}

		break

	}

	ipUser.Ip = ip
	UpdatePremio(ipUser)

	return true, ipUser.Premio

}
