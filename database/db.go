package database

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "keshika"
	dbname   = "godb"
)

var DB *gorm.DB

func Database() {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

}
