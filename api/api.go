package api

import (
	"dv4all/fiber-pg/db"
	"dv4all/fiber-pg/utils"
	"log"
	"os"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Config struct {
	name string
	host string
}

func Start() *fiber.App {
	// load AppName
	cfg := Config{
		name: utils.GetEnv("API_NAME", "fiber-pg"),
		host: utils.GetEnv("API_HOST", ":8080"),
	}

	envProc := utils.GetEnv("API_MAXPROC", "UNDEFINED")
	if envProc != "UNDEFINED" {
		// define max number of processes
		// by default prefork will spawn process on each processor core
		runtime.GOMAXPROCS(4)
	}
	// create fiber app
	app := fiber.New(fiber.Config{
		// enable multiple processes to run
		Prefork: true,
	})
	// register logger middleware
	app.Use(logger.New())

	//static
	app.Static("/", "./static")

	// register all other routes
	registerRoutes(app)

	// start listening
	err := app.Listen(cfg.host)
	// check for errors
	if err != nil {
		os.Exit(1)
		return nil
	}
	// log this
	log.Printf("%v started [%v]", cfg.name, cfg.host)
	// return pointer
	return app
}

func registerRoutes(api *fiber.App) {
	// log.Println("Register routes...")

	//todo list
	api.Get("/todos", getTodoLists)
	// api.Post("/todos", addTodoList)
	// api.Put("/todos", updateTodoList)
	// api.Delete("/todos/:listid", deleteTodoList)

}

func getTodoLists(ctx *fiber.Ctx) error {
	//get data from database
	tls, err := db.GetAllTodoLists()
	// check for error
	if err != nil {
		// construct error message
		errResp := SetErrorResponse(err, ServerStatus{Status: 0, StatusText: ""})
		return ctx.JSON(errResp)
	}
	// construct response
	response := SetOKResponse(tls)
	return ctx.JSON(response)
}