package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"sample/models"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))+"?parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4")
	if err != nil {
		panic(fmt.Errorf("Connecting DB Error: %w", err))
	}
	defer _db.Close()

	db := models.NewDB(_db)

	err = db.Migrate()
	if err != nil {
		panic(fmt.Errorf("Migrate Error: %w", err))
	}

	user, err := db.Users().Find()
	if err == models.RECORD_NOT_FOUND {
		log.Println("Record Not Found")
	} else if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", user)

	users, err := db.Users().Select()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", users)

	name := "mazrean"
	password := "testtest"
	err = db.Users().Insert(&models.Users{
		Name:     &name,
		Password: &password,
	})
	if err != nil {
		panic(err)
	}

	user, err = db.Users().Find()
	if err == models.RECORD_NOT_FOUND {
		log.Println("Record Not Found")
	} else if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", user)

	password = "nyannyan"
	err = db.Users().Where(&models.Users{
		Id: &user.Id,
	}).Update(&models.Users{
		Password: &password,
	})
	if err != nil {
		panic(err)
	}

	user, err = db.Users().Where(&models.Users{
		Id: &user.Id,
	}).Find()
	if err == models.RECORD_NOT_FOUND {
		log.Println("Record Not Found")
	} else if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", user)

	err = db.Users().Where(&models.Users{
		Id: &user.Id,
	}).Delete()
	if err != nil {
		panic(err)
	}

	user, err = db.Users().Where(&models.Users{
		Id: &user.Id,
	}).Find()
	if err == models.RECORD_NOT_FOUND {
		log.Println("Record Not Found")
	} else if err != nil {
		panic(err)
	}
}
