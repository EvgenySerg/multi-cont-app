package internal

import (
	"udemy/docker/worker/internal/adapter"
	"udemy/docker/worker/internal/ratestclient"
)

func NewHandler(dbRepo adapter.Repo, cache adapter.Repo, client ratestclient.Ingestion) *Handler {
	return &Handler{
		db:        dbRepo,
		cache:     cache,
		ingestion: client,
	}
}

type Handler struct {
	db        adapter.Repo
	cache     adapter.Repo
	ingestion ratestclient.Ingestion
}


func (s *Handler) ProcessRates() error{
	data, err:=s.ingestion.GetRates()
	if err!=nil{
		return err
	}
	err=s.cache.Save(adapter.ToStorage(data))
	if err!=nil{
		return err
	}
	err=s.db.Save(adapter.ToCache(data))
	if err!=nil{
		return err
	}

	return nil
}


