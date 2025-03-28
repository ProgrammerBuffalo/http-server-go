package router

import (
	"http-app/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ServerRouter struct {
	accountHandler *handler.AccountHandler
}

func NewRouter(accountHandler *handler.AccountHandler) *ServerRouter {
	return &ServerRouter{
		accountHandler: accountHandler,
	}
}

func (sr *ServerRouter) Routes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/accounts", sr.accountHandler.GetAccounts)

	return r
}
