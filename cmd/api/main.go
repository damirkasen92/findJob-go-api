package main

import (
	"log"
	"net/http"

	"github.com/damir/jobfinder/app"
)

func main() {
	app, err := app.New()

	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(":9000", app.Router)

	if err != nil {
		log.Fatal(err)
	}
}
