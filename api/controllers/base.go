package controllers

import (
	"fmt"
	"net/http"

	"todo-api/api/models"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)

		if err != nil {
			log.Fatalf("Cannot connect to %s database: %v", Dbdriver, err)
		}

		log.Info("COnnected to Database .....")
	}
	server.Migrate()
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Migrate() {
	server.DB.Debug().AutoMigrate(
		models.User{},
		models.Todo{},
	)
}

func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(": %v", addr), server.Router))
	log.Info("Server is running on port", addr)
}
