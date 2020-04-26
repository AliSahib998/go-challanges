package client

import (
	"github.com/AliSahib998/go-challanges/model"
	"github.com/AliSahib998/go-challanges/util"
	"github.com/go-resty/resty/v2"
)

type CovidClient struct {
	Client *resty.Client
}

type ICovidClient interface {
	GetCovidData() ([]model.Covid, util.RestErrorResponse)
}

func (c *CovidClient) GetCovidData() ([]model.Covid, util.RestErrorResponse) {

	var covid []model.Covid

	resp, err := c.Client.R().
		EnableTrace().SetResult(&covid).
		Get("https://covid-az.herokuapp.com/api/")

	if err != nil {
		return nil, util.RestErrorResponse{Code: resp.StatusCode(), Message: err.Error()}
	}

	return covid, util.RestErrorResponse{}
}
