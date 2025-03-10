package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/reinhardlinardi/atm-report/docs"
	"github.com/reinhardlinardi/atm-report/internal/httpresp"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title ATM Report Service API
// @version 1.0

// @Host localhost:8000
// @BasePath /api/v1

func (srv *Server) RegisterHandlers() {
	r := srv.router

	r.Handle("/docs/*", http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))

	url := fmt.Sprintf("http://localhost:%d/docs/swagger.json", srv.config.Port)
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(url)))

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/count", func(r chi.Router) {
			r.Get("/daily", srv.countDaily)
			r.Get("/type", srv.countByType)
			r.Get("/daily-type", srv.countDailyByType)
		})

		r.Route("/max/withdraw", func(r chi.Router) {
			r.Get("/daily", srv.maxWithdrawDaily)
		})
	})
}

func (srv *Server) countDaily(w http.ResponseWriter, r *http.Request) {
	repo := srv.transactionRepo
	res, err := repo.CountDaily()

	if err != nil {
		fmt.Printf("err count daily: %s\n", err.Error())
		httpresp.Write(w, httpresp.ErrServer, nil)
		return
	}

	httpresp.Write(w, nil, res)
}

func (srv *Server) countByType(w http.ResponseWriter, r *http.Request) {
	repo := srv.transactionRepo
	res, err := repo.CountByType()

	if err != nil {
		fmt.Printf("err count by type: %s\n", err.Error())
		httpresp.Write(w, httpresp.ErrServer, nil)
		return
	}

	httpresp.Write(w, nil, res)
}

func (srv *Server) countDailyByType(w http.ResponseWriter, r *http.Request) {
	repo := srv.transactionRepo
	res, err := repo.CountDailyByType()

	if err != nil {
		fmt.Printf("err count daily by type: %s\n", err.Error())
		httpresp.Write(w, httpresp.ErrServer, nil)
		return
	}

	httpresp.Write(w, nil, res)
}

func (srv *Server) maxWithdrawDaily(w http.ResponseWriter, r *http.Request) {
	repo := srv.transactionRepo
	res, err := repo.MaxWithdrawDaily()

	if err != nil {
		fmt.Printf("err get max withdraw daily: %s\n", err.Error())
		httpresp.Write(w, httpresp.ErrServer, nil)
		return
	}

	httpresp.Write(w, nil, res)
}
