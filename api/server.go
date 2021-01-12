package api

import (
	"os"
	"todo-api/api/controllers"
)

var server = controllers.Server{}

func Run() {
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	server.Run(os.Getenv("SERVER_PORT"))
}
