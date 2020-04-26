package model

type Covid struct {
	Country        string `json:"country"`
	NewCases       string `json:"new_cases"`
	NewDeaths      string `json:"new_deaths"`
	TotalCases     string `json:"total_cases"`
	TotalDeaths    string `json:"total_deaths"`
	TotalRecovered string `json:"total_recovered"`
}
