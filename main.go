package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/pandakn/go-sut-course-api/config"
	"github.com/pandakn/go-sut-course-api/internal/server"
)

func envPath() string {
	// e.g., ./command-line-arguments a b c d (5 arguments)
	// os.Args[1] is a

	if len(os.Args) == 1 {
		return ".env.local"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())
	server := server.New(cfg)

	server.RegisterFiberRoutes()
	// port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := server.Listen(cfg.App().Url())
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
