package main

import (
	"log"

	"github.com/donairl/gofiber-template/routers"
)

func main() {

	app := routers.New()

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
