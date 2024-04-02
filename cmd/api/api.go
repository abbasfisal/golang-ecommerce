package api

import (
	"database/sql"
	"github.com/abbasfisal/golang-ecommerce/service/user"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s APIServer) Run() error {

	//init router

	//init routes

	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	//user routes

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)
	//
	log.Println("Listening On :", s.addr)
	return http.ListenAndServe(s.addr, router)
}
