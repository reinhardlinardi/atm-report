package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mocktransaction "github.com/reinhardlinardi/atm-report/internal/mocks/transaction"
	"github.com/reinhardlinardi/atm-report/internal/transaction"
	"github.com/stretchr/testify/assert"
)

func TestCountDaily(t *testing.T) {
	e := errors.New("any error")

	tts := []struct {
		caseName   string
		mockExpect func(*mocktransaction.Repository)
		expect     int
	}{
		{
			caseName: "count daily error",
			mockExpect: func(m *mocktransaction.Repository) {
				m.On("CountDaily").Return(nil, e)
			},
			expect: http.StatusInternalServerError,
		},
		{
			caseName: "count daily ok",
			mockExpect: func(m *mocktransaction.Repository) {
				m.On("CountDaily").Return([]transaction.DailyCount{}, nil)
			},
			expect: http.StatusOK,
		},
	}

	for _, tt := range tts {
		t.Log(tt.caseName)

		m := new(mocktransaction.Repository)
		tt.mockExpect(m)

		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		server := &Server{transaction: m}
		server.countDaily(w, req)

		resp := w.Result()
		assert.Equal(t, tt.expect, resp.StatusCode)
	}
}

func TestCountDailyByType(t *testing.T) {
	e := errors.New("any error")

	tts := []struct {
		caseName   string
		mockExpect func(*mocktransaction.Repository)
		expect     int
	}{
		{
			caseName: "count daily by type error",
			mockExpect: func(m *mocktransaction.Repository) {
				m.On("CountDailyByType").Return(nil, e)
			},
			expect: http.StatusInternalServerError,
		},
		{
			caseName: "count daily by type ok",
			mockExpect: func(m *mocktransaction.Repository) {
				m.On("CountDailyByType").Return([]transaction.DailyTypeCount{}, nil)
			},
			expect: http.StatusOK,
		},
	}

	for _, tt := range tts {
		t.Log(tt.caseName)

		m := new(mocktransaction.Repository)
		tt.mockExpect(m)

		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		server := &Server{transaction: m}
		server.countDailyByType(w, req)

		resp := w.Result()
		assert.Equal(t, tt.expect, resp.StatusCode)
	}
}

func TestCountDailyAll(t *testing.T) {
	e := errors.New("any error")

	tts := []struct {
		caseName   string
		mockExpect func(*mocktransaction.Repository)
		expect     int
	}{
		{
			caseName: "count daily all error",
			mockExpect: func(m *mocktransaction.Repository) {
				m.On("CountDailyByType").Return(nil, e)
			},
			expect: http.StatusInternalServerError,
		},
		{
			caseName: "count daily all ok",
			mockExpect: func(m *mocktransaction.Repository) {
				m.On("CountDailyByType").Return([]transaction.DailyTypeCount{}, nil)
			},
			expect: http.StatusOK,
		},
	}

	for _, tt := range tts {
		t.Log(tt.caseName)

		m := new(mocktransaction.Repository)
		tt.mockExpect(m)

		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		server := &Server{transaction: m}
		server.countDailyAll(w, req)

		resp := w.Result()
		assert.Equal(t, tt.expect, resp.StatusCode)
	}
}

func TestGetDailyMaxWithdraw(t *testing.T) {
	e := errors.New("any error")

	tts := []struct {
		caseName   string
		mockExpect func(*mocktransaction.Repository)
		expect     int
	}{
		{
			caseName: "get daily max withdraw error",
			mockExpect: func(m *mocktransaction.Repository) {
				m.On("GetDailyMaxWithdraw").Return(nil, e)
			},
			expect: http.StatusInternalServerError,
		},
		{
			caseName: "get daily max withdraw ok",
			mockExpect: func(m *mocktransaction.Repository) {
				m.On("GetDailyMaxWithdraw").Return([]transaction.DailyMaxWithdraw{}, nil)
			},
			expect: http.StatusOK,
		},
	}

	for _, tt := range tts {
		t.Log(tt.caseName)

		m := new(mocktransaction.Repository)
		tt.mockExpect(m)

		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		server := &Server{transaction: m}
		server.getDailyMaxWithdraw(w, req)

		resp := w.Result()
		assert.Equal(t, tt.expect, resp.StatusCode)
	}
}
