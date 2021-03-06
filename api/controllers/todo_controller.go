package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"todo-api/api/auth"
	"todo-api/api/models"
	"todo-api/api/responses"

	"github.com/gorilla/mux"
)

func (server *Server) CreateTodo(w http.ResponseWriter, r *http.Request) {
	uid, _ := auth.ExtractTokenID(r)
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	todo := models.Todo{}
	err = json.Unmarshal(body, &todo)
	todo.UserID = uid
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	todo.Prepare()
	err = todo.Validate()

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	newTodo, err := todo.Save(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, newTodo)
}

func (server *Server) ListTodos(w http.ResponseWriter, r *http.Request) {
	uid, _ := auth.ExtractTokenID(r)
	todo := models.Todo{}
	todos, err := todo.FindAllTodoByUser(server.DB, uid)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, todos)
}

func (server *Server) TodoDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	todo := models.Todo{}
	todoDetails, err := todo.Find(server.DB, uint(id))

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, todoDetails)
}

func (server *Server) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	todo := models.Todo{}

	err = json.Unmarshal(body, &todo)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	_, err = todo.Update(server.DB, id)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	updatedTodo, err := todo.Find(server.DB, uint(id))
	responses.JSON(w, http.StatusCreated, updatedTodo)
}

func (server *Server) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	todo := models.Todo{}

	err = json.Unmarshal(body, &todo)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	_, err = todo.Find(server.DB, uint(id))

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	_, err = todo.Delete(server.DB, uint(id))

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, map[string]string{
		"status": "Ok",
	})
}
