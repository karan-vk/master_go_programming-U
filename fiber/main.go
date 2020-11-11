package main

import (
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// fbapp, err := firebase.NewApp(c.Background(), nil)
		// if err != nil {
		// 	log.Fatalf("error initializing app: %v\n", err)
		// }
	// Print current process
	if fiber.IsChild() {
		fmt.Printf("[%d] Child", os.Getppid())
	} else {
		fmt.Printf("[%d] Master", os.Getppid())
	}

	// Fiber instance
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	// Routes
	app.Get("/", hello)

	// Start server
	log.Fatal(app.Listen(":3000"))

	// Run the following command to see all processes sharing port 3000:
	// sudo lsof -i -P -n | grep LISTEN
}

// Handler
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
