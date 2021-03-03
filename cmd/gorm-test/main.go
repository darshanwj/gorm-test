package main

import (
	"darshanwj/gorm-test/internal"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// DB connection, read from config
	db, err := gorm.Open(
		mysql.Open("root:root@tcp(127.0.0.1:3326)/gokit"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	if err != nil {
		log.Fatalln(err)
	}
	// User service
	us := internal.NewUserService(db)
	// Http transport handler
	h := internal.NewHTTPHandler(us)
	// Running
	_ = http.ListenAndServe(":8082", h)
}
