package model

import "time"

type LatestRates struct {
	Rates map[string]float64
	Base  string
	Date  time.Time
}
