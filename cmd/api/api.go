package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Dhruv2638/ecom-go/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (a *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subRouter)

	log.Println("Listning on : ", a.addr)
	return http.ListenAndServe(a.addr, router)
}
