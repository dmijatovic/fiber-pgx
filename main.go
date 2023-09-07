package main

import (
	"dv4all/fiber-pg/api"
	"dv4all/fiber-pg/db"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func onQuit(app *fiber.App, pgdb *pgxpool.Pool, close chan bool) {
	// Create channel to signify interupt signal is being sent
	interupt := make(chan os.Signal, 1)
	// When an interrupt is sent, notify the channel
	signal.Notify(interupt, os.Interrupt)

	// shutdown api
	app.Shutdown()
	log.Println("Shutdown api...")
	// shutdown database
	pgdb.Close()
	log.Println("Closing database...")

	// notify others now
	close <- true
}

func main() {
	// log.Println("Starting...")
	// connect DB
	pgdb := db.Connect()

	// start api
	app := api.Start()

	// Create channel to signify interupt signal is being sent
	close := make(chan bool, 1)

	// listen to quit event
	go onQuit(app, pgdb, close)

	// wait for close sign
	_ = <-close

	// log.Println("Everything closed...")
}
