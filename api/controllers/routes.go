package controllers

import "github.com/FarOutWest/genity/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

  // Data Route
	s.Router.HandleFunc("/post-data", middlewares.SetMiddlewareJSON(s.CreateData)).Methods("POST")
	s.Router.HandleFunc("/get-data", middlewares.SetMiddlewareJSON(s.GetDatas)).Methods("GET")
	s.Router.HandleFunc("/get-data/{id}", middlewares.SetMiddlewareJSON(s.GetData)).Methods("GET")
}
