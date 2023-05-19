package database

import (
	"fmt"
	"log"
	"os"

	"github.com/fydhfzh/fp-4/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	err error
)

var (
	port = os.Getenv("PG_PORT")
	user = os.Getenv("PG_USER")
	password = os.Getenv("PG_PASSWORD")
	host = os.Getenv("PG_HOST")
	dbname = os.Getenv("PG_DBNAME")
)

func InitializedDatabase(){
	handleDBConnection()
	createRequiredTable()
}

func handleDBConnection(){
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", host, user, password, port, dbname)

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