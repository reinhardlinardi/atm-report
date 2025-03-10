package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/reinhardlinardi/atm-report/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title ATM Report Service API
// @version 1.0

// @Host localhost:8000
// @BasePath /api/v1

func (server *Server) RegisterHandlers() {
	r := server.router

	r.Handle("/docs/*", http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))

	url := fmt.Sprintf("http://localhost:%d/docs/swagger.json", server.config.Port)
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(url)))
	// r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/count", func(r chi.Router) {
			r.Get("/daily", server.countDaily)
			r.Get("/type", server.countByType)
			r.Get("/daily-type", server.countDailyByType)
		})

		r.Route("/max/withdraw", func(r chi.Router) {
			r.Get("/daily", server.maxWithdrawDaily)
		})
	})
}

func (server *Server) countDaily(w http.ResponseWriter, r *http.Request) {
	// repo := server.transactionRepo
	// res, err := repo.CountDaily()

	// if err != nil {
	// 	fmt.Printf("err count daily: %s\n", err.Error())
	// 	httpresp.Write(w, httpresp.ErrServer, nil)
	// 	return
	// }

	// httpresp.Write(w, nil, res)
}

func (server *Server) countByType(w http.ResponseWriter, r *http.Request) {
	// repo := server.transactionRepo
	// res, err := repo.CountByType()

	// if err != nil {
	// 	fmt.Printf("err count by type: %s\n", err.Error())
	// 	httpresp.Write(w, httpresp.ErrServer, nil)
	// 	return
	// }

	// httpresp.Write(w, nil, res)
}

func (server *Server) countDailyByType(w http.ResponseWriter, r *http.Request) {
	// repo := server.transactionRepo
	// res, err := repo.CountDailyByType()

	// if err != nil {
	// 	fmt.Printf("err count daily by type: %s\n", err.Error())
	// 	httpresp.Write(w, httpresp.ErrServer, nil)
	// 	return
	// }

	// httpresp.Write(w, nil, res)
}

func (server *Server) maxWithdrawDaily(w http.ResponseWriter, r *http.Request) {
	// repo := server.transactionRepo
	// res, err := repo.MaxWithdrawDaily()

	// if err != nil {
	// 	fmt.Printf("err get max withdraw daily: %s\n", err.Error())
	// 	httpresp.Write(w, httpresp.ErrServer, nil)
	// 	return
	// }

	// httpresp.Write(w, nil, res)
}
