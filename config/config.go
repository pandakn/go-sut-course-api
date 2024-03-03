package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

type IAppConfig interface {
	Url() string
	Name() string
}

type IConfig interface {
	App() IAppConfig
}

type config struct {
	app *app
}

type app struct {
	host string
	port string
	name string
}

func LoadConfig(path string) IConfig {
	envMap, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("load dotenv failed: %v", err)
	}

	return &config{
		app: &app{
			host: envMap["APP_HOST"],
			port: envMap["APP_PORT"],
			name: envMap["APP_NAME"],
		},
	}
}

func (c *config) App() IAppConfig { return c.app }

func (a *app) Url() string {
	return fmt.Sprintf("%s:%s", a.host, a.port)
}

func (a *app) Name() string { return a.name }
