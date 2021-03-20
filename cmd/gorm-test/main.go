package main

import (
	"darshanwj/gorm-test/internal"
	"darshanwj/gorm-test/internal/config"
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
	defer func() {
		_ = log.Sync()
	}()

	// load config
	conf, err := config.NewReader(log).Read()
	if err != nil {
		log.Fatal("could not load config", zap.Error(err))
	}
	log.Debug("loaded config", zap.Any("settings", conf.AllSettings()))

	// DB connection
	db, err := gorm.Open(
		mysql.Open(conf.GetString("mysql.dsn")),
		&gorm.Config{Logger: gormlogger.Default.LogMode(gormlogger.Info)},
	)
	if err != nil {
		log.Fatal("could not connect to mysql db", zap.Error(err))
	}
	log.Debug("connected to mysql db")

	// User service
	us := internal.NewUserService(db, log)

	// Http transport router
	h := internal.NewHTTPHandler(us)

	// decorate the router
	middleware := internal.RequestLoggingMiddleware(log)
	h = middleware(h)

	// Run web server
	log.Debug("serving...")
	err = http.ListenAndServe(":8082", h)
	if err != nil {
		log.Fatal("could not start server", zap.Error(err))
	}
}
