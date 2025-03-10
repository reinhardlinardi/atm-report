package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/reinhardlinardi/atm-report/docs"
	"github.com/reinhardlinardi/atm-report/internal/httpjson"
	"github.com/reinhardlinardi/atm-report/internal/transaction"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

//	@title			ATM Report API
//	@version		1.0
//	@description	ATM report service API server
//	@host			localhost:8000
//	@BasePath		/api/v1/daily

func (server *Server) RegisterHandlers() {
	r := server.router
	r.Get("/swagger/*", httpSwagger.WrapHandler)

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

// countDaily godoc
//
//	@Summary		Count daily transactions
//	@Description	Get number of transactions per day
//	@Tags			Count
//	@Produce		json
//	@Success		200	{object}	[]transaction.DailyCount
//	@Failure		500	{object}	httpjson.Response
//	@Router			/count [get]
func (server *Server) countDaily(w http.ResponseWriter, r *http.Request) {
	data, err := server.transaction.CountDaily()
	if err != nil {
		fmt.Printf("err count daily: %s\n", err.Error())
		httpjson.InternalError(w, nil)
		return
	}

	httpjson.OK(w, data)
}

// countDailyByType godoc
//
//	@Summary		Count daily transactions per type
//	@Description	Get number of transactions per day per transaction type
//	@Tags			Count
//	@Produce		json
//	@Success		200	{object}	[]transaction.DailyTypeCount
//	@Failure		500	{object}	httpjson.Response
//	@Router			/count/type [get]
func (server *Server) countDailyByType(w http.ResponseWriter, r *http.Request) {
	data, err := server.transaction.CountDailyByType()
	if err != nil {
		fmt.Printf("err count daily by type: %s\n", err.Error())
		httpjson.InternalError(w, nil)
		return
	}

	httpjson.OK(w, data)
}

// countDailyAll godoc
//
//	@Summary		Count daily transactions, with count per type
//	@Description	Get number of transactions per day, and number of transactions per day per type
//	@Tags			Count
//	@Produce		json
//	@Success		200	{object}	DailyAllResponse
//	@Failure		500	{object}	httpjson.Response
//	@Router			/count/all [get]
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

// getDailyMaxWithdraw godoc
//
//	@Summary		ATM with max withdraw per day
//	@Description	Get ATM with max withdraw amount per day
//	@Tags			Max
//	@Produce		json
//	@Success		200	{object}	[]transaction.DailyMaxWithdraw
//	@Failure		500	{object}	httpjson.Response
//	@Router			/max/withdraw [get]
func (server *Server) getDailyMaxWithdraw(w http.ResponseWriter, r *http.Request) {
	data, err := server.transaction.GetDailyMaxWithdraw()
	if err != nil {
		fmt.Printf("err get daily max withdraw: %s\n", err.Error())
		httpjson.InternalError(w, nil)
		return
	}

	httpjson.OK(w, data)
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
