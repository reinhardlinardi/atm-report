package server

import "github.com/reinhardlinardi/atm-report/internal/transaction"

type DailyAllResponse struct {
	Total  []transaction.DailyCount     `json:"total"`
	Detail []transaction.DailyTypeCount `json:"detail"`
}
