package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
	err error
)

const (
	port = 5432
	user = "admin"
	password = "postgres"
	dialect = "postgres"
	host = "localhost"
	dbname = "todo"
)

func InitializedDatabase(){
	handleDBConnection()
	createRequiredTable()
}

func handleDBConnection(){
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, user, password, port, dbname)


	db, err = sql.Open(dialect, psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}
}

func createRequiredTable(){

}