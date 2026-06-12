package main

import (
	"log"
)

func main() {
	config := config{
		address: "localhost",
		port:    8080,
	}

	app := Application{
		config: config,
	}

	if err := app.run(app.mount()); err != nil {
		log.Fatal("Application is shutting down to error: ", err)
	}
}
