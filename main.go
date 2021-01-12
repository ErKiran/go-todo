package main

import (
	"todo-api/api"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// @title todo-api
// @version 1.0.1
// @description Basic Todo API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email kiruu1238@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:6000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not coming through %v", err)
	}
	api.Run()
}
