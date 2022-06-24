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
	cadena := "b5ef58cbb8c454:0f5b1694@tcp(us-cdbr-east-04.cleardb.com)/heroku_ae44dd41be82398"
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
