package client

import (
	"encoding/json"
	"github.com/AliSahib998/go-challanges/model"
	"github.com/AliSahib998/go-challanges/util"
	"github.com/go-resty/resty/v2"
)

type CovidClient struct {
	Client *resty.Client
}

type ICovidClient interface {
	GetCovidData() ([]model.Covid, util.ErrorResponse)
}

func (c *CovidClient) GetCovidData() ([]model.Covid, util.ErrorResponse) {

	var covid []model.Covid

	resp, err := c.Client.R().
		EnableTrace().
		Get("https://covid-az.herokuapp.com/api/")

	if err != nil {
		return nil, util.ErrorResponse{Code: resp.StatusCode(), Message: err.Error()}
	}

	error := json.Unmarshal(resp.Body(), &covid)

	if error != nil {
		return nil, util.ErrorResponse{Code: 400, Message: "Parse Exception"}
	}

	return covid, util.ErrorResponse{}
}
