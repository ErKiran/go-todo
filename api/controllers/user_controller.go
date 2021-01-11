package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"todo-api/api/auth"
	"todo-api/api/models"
	"todo-api/api/responses"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate()

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		log.Fatalf("Hashing Error")
	}
	user.Password = string(hash)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	newUser, err := user.Save(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, newUser)
}

func (server *Server) LoginUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.Validate()

	userInfo, err := user.FindByEmail(server.DB, user.Email)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if user.Password != userInfo.Password {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	signedToken, err := auth.CreateToken(userInfo.ID, user.Email)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, map[string]string{
		"token": signedToken,
	})
}
