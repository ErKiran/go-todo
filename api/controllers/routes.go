package controllers

import (
	"net/http"
	"todo-api/api/middleware"
)

func (server *Server) ProtectedRoute(path string, next func(http.ResponseWriter, *http.Request), method string) {
	server.Router.HandleFunc(path, middleware.SetMiddlewareAuthentication(next, server.DB)).Methods(method, "OPTIONS")
}

func (server *Server) OwnerOnlyRoute(path string, next func(http.ResponseWriter, *http.Request), method string) {
	server.Router.HandleFunc(path, middleware.TodoOwner(next, server.DB)).Methods(method, "OPTIONS")
}

func (server *Server) initializeRoutes() {
	server.ProtectedRoute("/todo", server.CreateTodo, "POST")
	server.ProtectedRoute("/todo", server.ListTodos, "GET")
	server.OwnerOnlyRoute("/todo/{id}", server.TodoDetail, "GET")
	server.OwnerOnlyRoute("/todo/{id}", server.UpdateTodo, "PUT")
	server.OwnerOnlyRoute("/todo/{id}", server.DeleteTodo, "DELETE")
	server.Router.HandleFunc("/user", server.CreateUser).Methods("POST")
	server.Router.HandleFunc("/user/login", server.LoginUser).Methods("POST")
}
