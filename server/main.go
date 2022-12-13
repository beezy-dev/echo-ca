package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/beezy-dev/echo-ca/server/pkg/types"
	"github.com/gofiber/fiber/v2"
)

func main() {

	 // define current version status for echo-ca health endpoint
	data := types.Build{
		AppName: "Echo-CA",
		Version: "v1beta",
	}

	// set up fiber with the version status 
	server := fiber.New(fiber.Config{
		AppName: data.AppName,
		DisableStartupMessage: true,
	})

	// route to health
	server.Get("/api/health", func(c *fiber.Ctx) error {
		log.Println("Request from |", c.IP(),"|", c.Method(), "|", c.Path())
		return c.JSON(data)
	})

	// set up the func for gracefull shutdown of fiber service
	go func() {
		if err := server.Listen(":8080"); err != nil {
			log.Panic(err)
		}
	}()

	log.Println("Starting Echo-CA service", "version", data.Version)

	// set up a channel to send the signal to and set the notification
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	_ = <-c 
	log.Println("SIGTERM received! Initiating gracefull shutdown...")
	_ = server.Shutdown()
	log.Println("Cleaning up running task...")
	// tasks to perform prior to shutdown
	// dump JSON structures to files 
	log.Println("Echo-CA was successfully shutdown.")

}
