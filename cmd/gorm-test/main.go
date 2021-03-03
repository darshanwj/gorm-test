package main

import (
	"darshanwj/gorm-test/internal"
	"net/http"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

func main() {
	// logger
	log, _ := zap.NewDevelopment()
	// flushes buffer, if any
	defer log.Sync()

	// DB connection, read from config
	db, err := gorm.Open(
		mysql.Open("root:root@tcp(127.0.0.1:3326)/gokit"),
		&gorm.Config{Logger: gormlogger.Default.LogMode(gormlogger.Info)},
	)
	if err != nil {
		log.Fatal("could not connect to mysql db", zap.Error(err))
	}

	// User service
	us := internal.NewUserService(db, log)

	// Http transport handler
	h := internal.NewHTTPHandler(us)

	// Run web server
	err = http.ListenAndServe(":8082", h)
	if err != nil {
		log.Fatal("could not start server", zap.Error(err))
	}
}
