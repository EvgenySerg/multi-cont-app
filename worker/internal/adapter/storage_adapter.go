package adapter

import (
	"udemy/docker/worker/internal/ratestclient"
	"udemy/docker/worker/internal/storage/model"
)

func ToStorage(webResp ratestclient.LatestResponse) model.LatestRates {
	return model.LatestRates{
		Rates: webResp.Rates,
		Base:  webResp.Base,
		Date:  webResp.Date,
	}
}

func ToCache(webResp ratestclient.LatestResponse)model.LatestRates {
	return model.LatestRates{
		Rates: webResp.Rates,
		Base:  webResp.Base,
		Date:  webResp.Date,
	}

}
