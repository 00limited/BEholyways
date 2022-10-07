package routes

import (
	"holyways/handlers"
	"holyways/pkg/middleware"
	"holyways/pkg/mysql"
	"holyways/repositories"

	"github.com/gorilla/mux"
)

func Auth(r *mux.Router) {
	authRepo := repositories.RepositoryAuth(mysql.DB)

	h := handlers.HandlerAuth(authRepo)

	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/check-auth", middleware.Auth(h.CheckAuth)).Methods("GET")
}
