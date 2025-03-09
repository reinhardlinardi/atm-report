package model

import "time"

type FileLoad struct {
	Id    int64     `db:"id"`
	AtmId string    `db:"atm_id"`
	Date  time.Time `db:"date"`
}
