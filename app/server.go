package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/reinhardlinardi/atm-report/internal/config"
	"github.com/reinhardlinardi/atm-report/internal/repository/transactionrepo"
)

type Server struct {
	config          *config.Config
	http            *http.Server
	transactionRepo transactionrepo.Repository
}

func NewServer(
	config *config.Config,
	transactionRepo transactionrepo.Repository,
) *Server {
	addr := fmt.Sprintf(":%d", config.Server.Port)
	server := &http.Server{Addr: addr, Handler: nil}

	return &Server{
		config:          config,
		http:            server,
		transactionRepo: transactionRepo,
	}
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
