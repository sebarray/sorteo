package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var ConDB = ConnectionDB()

func ConnectionDB() (conn *sql.DB) {
	driver := "mysql"
	cadena := "b6f9638509247c:aa8cd3ce@tcp(us-cdbr-east-05.cleardb.net)/heroku_e41bfd2bf0034e4"
	conn, err := sql.Open(driver, cadena)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("conectado con exito")
	return conn

}

func Checkconnection() bool {
	err := ConDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	return true
}
