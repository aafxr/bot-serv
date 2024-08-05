package routes

import (
	"github.com/aafxr/bot-serv/internal/app/apiserver"
	"github.com/aafxr/bot-serv/internal/app/store/sqlstore"
	"github.com/gorilla/mux"
)

type UserRoutes struct {
	userRepo sqlstore.UserRepository
}

func NewUserRoutes(r *mux.Router) []apiserver.Route {
	routes := []apiserver.Route{}
	return routes
}
