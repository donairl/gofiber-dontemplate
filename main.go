package main

import (
	"github.com/donairl/gofiber-dontemplate/routers"
)

func main() {

	app := routers.New()

	// Start the server on port 3000
	app.Listen(":3000")
}
