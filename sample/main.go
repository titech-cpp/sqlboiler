package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/titech-cpp/sqlboiler/sample/models"
)

func main() {
	_db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))+"?parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4")
	if err != nil {
		panic(fmt.Errorf("Connecting DB Error: %w", err))
	}

	db := models.NewDB(_db)

	err = db.Migrate()
	if err != nil {
		panic(fmt.Errorf("Migrate Error: %w", err))
	}
}
