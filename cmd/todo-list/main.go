package main

import (
	_ "github.com/eclipsemode/go-todo-app/docs"
	"github.com/eclipsemode/go-todo-app/internal/app"
	"github.com/joho/godotenv"
	logDefault "log"
)

// @title List to do API
// @version 1.0
// @description Simple Rest Api for to do list
// @termOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @externalDocs.description OpenAPI
// @externalDocs.url https://swagger.io/resources/open-api/
func main() {
	if err := godotenv.Load(); err != nil {
		logDefault.Fatal("Error loading .env file")
	}

	initApp()

	a, err := app.GetGlobalApp()
	if err != nil {
		logDefault.Fatal("Error getting global app")
	}

	err = a.StartHttpServer()
	if err != nil {
		logDefault.Fatal("Error starting http server")
	}
}

func initApp() {
	a, err := app.NewApp()
	if err != nil {
		panic(err)
	}

	app.SetGlobalApp(a)

}
