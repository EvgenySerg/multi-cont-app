package adapter

import "udemy/docker/worker/internal/storage/model"

type Repo interface {
	Save(latestRate model.LatestRates) error
}

