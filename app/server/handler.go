package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/reinhardlinardi/atm-report/docs"
	"github.com/reinhardlinardi/atm-report/internal/transaction"
	"github.com/reinhardlinardi/atm-report/pkg/httpjson"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title ATM Report Service API
// @version 1.0

// @Host localhost:8000
// @BasePath /api/v1

func (server *Server) RegisterHandlers() {
	r := server.router

	r.Handle("/docs/*", http.StripPrefix(
		"/docs/",
		http.FileServer(http.Dir("./docs")),
	))

	url := fmt.Sprintf("http://localhost:%d/docs/swagger.json", server.config.Port)
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(url)))
	// r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Route("/api/v1/daily", func(r chi.Router) {
		r.Route("/count", func(r chi.Router) {
			r.Get("/", server.countDaily)
			r.Get("/type", server.countDailyByType)
			r.Get("/all", server.countDailyAll)
		})

		r.Route("/max", func(r chi.Router) {
			r.Get("/withdraw", server.getDailyMaxWithdraw)
		})
	})
}

func (server *Server) countDaily(w http.ResponseWriter, r *http.Request) {
	res, err := server.transaction.CountDaily()
	if err != nil {
		fmt.Printf("err count daily: %s\n", err.Error())
		httpjson.InternalError(w, nil)
		return
	}

	httpjson.OK(w, res)
}

func (server *Server) countDailyByType(w http.ResponseWriter, r *http.Request) {
	res, err := server.transaction.CountDailyByType()
	if err != nil {
		fmt.Printf("err count daily by type: %s\n", err.Error())
		httpjson.InternalError(w, nil)
		return
	}

	httpjson.OK(w, res)
}

func (server *Server) countDailyAll(w http.ResponseWriter, r *http.Request) {
	data, err := server.transaction.CountDailyByType()
	if err != nil {
		fmt.Printf("err count daily all: %s\n", err.Error())
		httpjson.InternalError(w, nil)
		return
	}

	total := getTotalCount(data)
	res := DailyAllResponse{
		Total:  total,
		Detail: data,
	}

	httpjson.OK(w, res)
}

func (server *Server) getDailyMaxWithdraw(w http.ResponseWriter, r *http.Request) {
	res, err := server.transaction.GetDailyMaxWithdraw()
	if err != nil {
		fmt.Printf("err get daily max withdraw: %s\n", err.Error())
		httpjson.InternalError(w, nil)
		return
	}

	httpjson.OK(w, res)
}

func getTotalCount(data []transaction.DailyTypeCount) []transaction.DailyCount {
	cnt := make(map[string]int)
	total := []transaction.DailyCount{}

	for _, t := range data {
		if _, ok := cnt[t.Date]; !ok {
			cnt[t.Date] = 0
		}
		cnt[t.Date] += t.Count
	}

	for key, val := range cnt {
		total = append(total, transaction.DailyCount{
			Date:  key,
			Count: val,
		})
	}
	return total
}
