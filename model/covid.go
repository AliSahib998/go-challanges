package model

type Covid struct {
	Id             int64  `sql:"id,pk" json:"id"`
	Country        string `sql:"country" json:"country"`
	NewCases       string `sql:"new_cases" json:"new_cases"`
	NewDeaths      string `sql:"new_deaths" json:"new_deaths"`
	TotalCases     string `sql:"total_cases" json:"total_cases"`
	TotalDeaths    string `sql:"total_deaths" json:"total_deaths"`
	TotalRecovered string `sql:"total_recovered" json:"total_recovered"`
}
