package ratestclient

import "time"

type Ingestion interface {
	GetRates() (LatestResponse, error)
}


func NewWebIngestion(url string) *WebIngestion{
	return &WebIngestion{url:url}
}

type WebIngestion struct {
	url string
}

func (wi *WebIngestion) GetRates() (LatestResponse, error) {

}

type LatestResponse struct {
	Rates map[string]float64
	Base  string
	Date  time.Time
}
