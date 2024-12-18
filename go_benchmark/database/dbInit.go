package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

var DBSqlX *sqlx.DB

func DatabaseConnect() {
	dbStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_IP"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	db, err := sqlx.Connect("mysql", dbStr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected")
	DBSqlX = db
}
