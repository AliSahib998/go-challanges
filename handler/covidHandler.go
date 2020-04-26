package handler

import (
	"encoding/json"
	"github.com/AliSahib998/go-challanges/client"
	"github.com/AliSahib998/go-challanges/service"
	"github.com/AliSahib998/go-challanges/util"
	"github.com/go-chi/chi"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type covidHandler struct {
	covidServcie service.ICovidService
}

func NewCovidHandler(router *chi.Mux) *chi.Mux {
	c := &covidHandler{
		covidServcie: &service.CovidService{CovidClient: &client.CovidClient{Client: resty.New()}},
	}
	router.Get("/data", c.getCovidData)

	return router
}

func (c *covidHandler) getCovidData(w http.ResponseWriter, r *http.Request) {
	resp, err := c.covidServcie.GetCovidData()

	if err.Code > 0 {
		util.ErrorHandler(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
