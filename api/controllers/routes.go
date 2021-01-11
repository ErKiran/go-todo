package controllers

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/todo", server.CreateTodo).Methods("POST")
	server.Router.HandleFunc("/todo/{id}", server.UpdateTodo).Methods("PUT")
	server.Router.HandleFunc("/todo/{id}", server.DeleteTodo).Methods("DELETE")
}
