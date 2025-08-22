package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/pbac_db?parseTime=true")
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("DB not reachable:", err)
	}
}
