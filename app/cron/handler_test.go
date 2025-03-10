package cron

import (
	"encoding/json"
	"errors"
	"testing"

	mockstorage "github.com/reinhardlinardi/atm-report/internal/mocks/filestorage"
	mockhistory "github.com/reinhardlinardi/atm-report/internal/mocks/history"
	mocktransaction "github.com/reinhardlinardi/atm-report/internal/mocks/transaction"
	"github.com/reinhardlinardi/atm-report/internal/transaction"
	"github.com/stretchr/testify/assert"
)

func TestHandleFile(t *testing.T) {
	e := errors.New("any error")

	data := []transaction.Transaction{}
	bytes, _ := json.Marshal(data)

	path := "A_20250225_1.json"
	atmId := "A"
	date := "20250225"
	seq := 1

	tts := []struct {
		caseName   string
		mockExpect func(*mockstorage.Storage,
			*mockhistory.Repository,
			*mocktransaction.Repository,
		)
		expectVal bool
		expectErr bool
	}{
		{
			caseName: "history check error",
			mockExpect: func(mstorage *mockstorage.Storage,
				mhist *mockhistory.Repository,
				mtx *mocktransaction.Repository,
			) {
				mhist.On("Check", atmId, date, seq).Return(false, e)
			},
			expectVal: false,
			expectErr: true,
		},
		{
			caseName: "file skipped",
			mockExpect: func(mstorage *mockstorage.Storage,
				mhist *mockhistory.Repository,
				mtx *mocktransaction.Repository,
			) {
				mhist.On("Check", atmId, date, seq).Return(true, nil)
			},
			expectVal: false,
			expectErr: false,
		},
		{
			caseName: "read file error",
			mockExpect: func(mstorage *mockstorage.Storage,
				mhist *mockhistory.Repository,
				mtx *mocktransaction.Repository,
			) {
				mhist.On("Check", atmId, date, seq).Return(false, nil)
				mstorage.On("Get", path).Return([]byte{}, e)
			},
			expectVal: false,
			expectErr: true,
		},
		{
			caseName: "parse file error",
			mockExpect: func(mstorage *mockstorage.Storage,
				mhist *mockhistory.Repository,
				mtx *mocktransaction.Repository,
			) {
				mhist.On("Check", atmId, date, seq).Return(false, nil)
				mstorage.On("Get", path).Return([]byte{}, nil)
			},
			expectVal: false,
			expectErr: true,
		},
		{
			caseName: "load data error",
			mockExpect: func(mstorage *mockstorage.Storage,
				mhist *mockhistory.Repository,
				mtx *mocktransaction.Repository,
			) {
				mhist.On("Check", atmId, date, seq).Return(false, nil)
				mstorage.On("Get", path).Return(bytes, nil)
				mtx.On("Load", data).Return(int64(0), e)
			},
			expectVal: false,
			expectErr: true,
		},
		{
			caseName: "append history error",
			mockExpect: func(mstorage *mockstorage.Storage,
				mhist *mockhistory.Repository,
				mtx *mocktransaction.Repository,
			) {
				mhist.On("Check", atmId, date, seq).Return(false, nil)
				mstorage.On("Get", path).Return(bytes, nil)
				mtx.On("Load", data).Return(int64(0), nil)
				mhist.On("Append", atmId, date, seq).Return(int64(0), e)
			},
			expectVal: false,
			expectErr: true,
		},
		{
			caseName: "success",
			mockExpect: func(mstorage *mockstorage.Storage,
				mhist *mockhistory.Repository,
				mtx *mocktransaction.Repository,
			) {
				mhist.On("Check", atmId, date, seq).Return(false, nil)
				mstorage.On("Get", path).Return(bytes, nil)
				mtx.On("Load", data).Return(int64(0), nil)
				mhist.On("Append", atmId, date, seq).Return(int64(0), nil)
			},
			expectVal: true,
			expectErr: false,
		},
	}

	for _, tt := range tts {
		t.Log(tt.caseName)

		mstorage := new(mockstorage.Storage)
		mhist := new(mockhistory.Repository)
		mtx := new(mocktransaction.Repository)

		tt.mockExpect(mstorage, mhist, mtx)
		cron := &Cron{fileStorage: mstorage, history: mhist, transaction: mtx}

		val, err := cron.handleFile(path)
		assert.Equal(t, tt.expectVal, val)

		if tt.expectErr {
			assert.NotEqual(t, err, nil)
		} else {
			assert.Equal(t, err, nil)
		}
	}
}
