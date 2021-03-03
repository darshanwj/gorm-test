package main

import (
	"darshanwj/gorm-test/internal"
	"net/http"
)

func main() {
	us := internal.NewUserService()
	h := internal.NewHTTPHandler(us)

	_ = http.ListenAndServe(":8082", h)
}
