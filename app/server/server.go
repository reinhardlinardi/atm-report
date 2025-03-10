package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/reinhardlinardi/atm-report/internal/config"
	"github.com/reinhardlinardi/atm-report/internal/transaction"
)

type Server struct {
	config      *config.ServerConfig
	http        *http.Server
	router      *chi.Mux
	transaction transaction.Repository
}

func New(config *config.ServerConfig, transaction transaction.Repository) *Server {
	addr := fmt.Sprintf(":%d", config.Port)
	router := chi.NewRouter()

	server := &http.Server{Addr: addr, Handler: router}
	return &Server{config: config, http: server, router: router, transaction: transaction}
}

func (server *Server) Run(ctx context.Context, cancel context.CancelFunc) {
	fmt.Printf("Listening on :%d\n", server.config.Port)
	err := server.http.ListenAndServe()

	if !errors.Is(err, http.ErrServerClosed) && err != nil {
		fmt.Printf("err server listen: %s\n", err.Error())
		cancel()
	}
}

func (server *Server) Shutdown(ctx context.Context) {
	server.http.Shutdown(ctx)
}
