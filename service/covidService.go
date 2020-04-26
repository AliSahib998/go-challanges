package service

import (
	"github.com/AliSahib998/go-challanges/client"
	"github.com/AliSahib998/go-challanges/model"
	"github.com/AliSahib998/go-challanges/util"
	log "github.com/sirupsen/logrus"
)

type CovidService struct {
	CovidClient client.ICovidClient
}

type ICovidService interface {
	GetCovidData() ([]model.Covid, util.RestErrorResponse)
}

func (c *CovidService) GetCovidData() ([]model.Covid, util.RestErrorResponse) {

	log.Debug("start GetCovidData")

	resp, err := c.CovidClient.GetCovidData()

	if err.Code > 0 {
		return nil, err
	}

	log.Debug("end GetCovidData")

	return resp, util.RestErrorResponse{}
}
