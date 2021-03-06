package controllers

import "github.com/fawziridwan/backstack/api/middleware"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middleware.SetMiddlewareJson(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middleware.SetMiddlewareJson(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middleware.SetMiddlewareJson(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middleware.SetMiddlewareJson(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middleware.SetMiddlewareJson(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middleware.SetMiddlewareJson(middleware.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middleware.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/posts", middleware.SetMiddlewareJson(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", middleware.SetMiddlewareJson(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middleware.SetMiddlewareJson(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middleware.SetMiddlewareJson(middleware.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", middleware.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")
}
