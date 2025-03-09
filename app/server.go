package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/reinhardlinardi/atm-report/internal/config"
	"github.com/reinhardlinardi/atm-report/internal/repository/transactionrepo"
)

type Server struct {
	config          *config.Config
	http            *http.Server
	router          *chi.Mux
	transactionRepo transactionrepo.Repository
}

func NewServer(config *config.Config, transactionRepo transactionrepo.Repository) *Server {
	addr := fmt.Sprintf(":%d", config.Server.Port)
	router := chi.NewRouter()

	server := &http.Server{Addr: addr, Handler: router}
	return &Server{config: config, http: server, router: router, transactionRepo: transactionRepo}
}

func (srv *Server) Run(ctx context.Context, cancel context.CancelFunc) {
	fmt.Printf("Listening on :%d\n", srv.config.Server.Port)
	err := srv.http.ListenAndServe()

	if !errors.Is(err, http.ErrServerClosed) && err != nil {
		fmt.Printf("err server listen: %s\n", err.Error())
		cancel()
	}
}

func (srv *Server) Shutdown(ctx context.Context) {
	srv.http.Shutdown(ctx)
}

func (srv *Server) RegisterHandlers() {
	r := srv.router

	r.Route("/api/v1/daily", func(r chi.Router) {
		r.Get("/transaction-count", countTransactionHandler)
		r.Get("/max-withdraw", maxWithdrawHandler)
	})
}
