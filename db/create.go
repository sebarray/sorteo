package db

import "fmt"

func CreateIp(ip string) {

	query := `INSERT INTO direcction (ip) VALUE ('` + ip + `');`
	create, err := ConDB.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())

	}
	create.Exec()
}
