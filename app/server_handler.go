package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/reinhardlinardi/atm-report/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title ATM Report Service API
// @version 1.0

func (srv *Server) RegisterHandlers() {
	r := srv.router

	url := fmt.Sprintf("http://localhost:%d/docs/swagger.json", srv.config.Port)
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(url)))

	r.Handle("/docs/*", http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))

	r.Route("/api/v1/daily", func(r chi.Router) {
		r.Get("/transaction-count", countTransactionHandler)
		r.Get("/max-withdraw", maxWithdrawHandler)
	})
}

func countTransactionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("count transaction"))
}

func maxWithdrawHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("max withdraw"))
}
