package database

import (
	"fmt"
	"log"

	"github.com/fydhfzh/fp-4/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	err error
)

const (
	port = 5432
	user = "admin"
	password = "postgres"
	dialect = "postgres"
	host = "localhost"
	dbname = "toko_belanja"
)

func InitializedDatabase(){
	handleDBConnection()
	createRequiredTable()
}

func handleDBConnection(){
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, user, password, port, dbname)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

}

func createRequiredTable(){
	err := db.AutoMigrate(&entity.User{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entity.Category{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entity.Product{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entity.Transaction{})

	if err != nil {
		log.Fatal(err)
	}

}

func GetDatabaseInstance() *gorm.DB{
	return db
}