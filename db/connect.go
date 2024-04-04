package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DbConnection()  {
  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

  var err error

  Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err  != nil {
    log.Fatal(err)
  }else {
    log.Println("Database connected!")
  }
}

/*Connetion de forma nátiva*/
// var Db *sql.DB
//
// func DbConnection() {
// 	cfg := mysql.Config{
// 		User:   os.Getenv("DB_USER"),
// 		Passwd: os.Getenv("DB_PASS"),
// 		Net:    "tcp",
// 		Addr:   os.Getenv("DB_ADDR"),
// 		DBName: os.Getenv("DB_NAME"),
// 	}
// 	// Get a database handle.
// 	var err error
// 	Db, err = sql.Open("mysql", cfg.FormatDSN())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	pingErr := Db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr)
// 	}
// 	fmt.Println("Connected!")
// }
