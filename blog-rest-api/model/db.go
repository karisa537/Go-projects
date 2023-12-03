package model

import (
	"database/sql"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/go-sql-driver/mysql"

	
)

var db *sql.DB

type connection struct {
	Host string
	Port string
	User string
	Password string
	DBName string
}

func Init() {
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", &err.env())
		return
	}

	connInfo := connection {
		Host: os.Getenv("POSTGRES_URL"),
		Port: os.Getenv("POSTGRES_PORT"),
		User: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_USER"),
		DBName: os.Getenv("POSTGRES_DBNAME"),
}


	//try to pen the db connection with our connection information

	db, err = sql.open("postgres", connToString(connInfo))
	if err != nil {
			fmt.Printf("Error connecting to the DB: %s\n", err.Error())
			return
	} else {
		fmt.Printf("DB Connection is OK\n")
	}

	//check if we can ping our DB
	err = db.ping()
	if err != nil {
			fmt.Printf("Error, coul not ping DB: %s\n", err.Error())
			return
	} else {
		fmt.Printf("DB pinged successfully")
	}
}

//Take our connectionstruct and convert to a string for our db connection info
func connToString(info connection) string{
	return fmt.Sprint("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			info.Host, info.Port, info.User, info.Password, info.DBName,)
}


