package main

import (
	"pkg.cocoad.mobi/x/cache"
	"pkg.cocoad.mobi/x/db"
	"pkg.cocoad.mobi/x/log"
)

type Application struct {
	qr    *db.Querier
	cache cache.Cacher
	cfg   *Config
}

func NewApplication(cfg *Config) (*Application, error) {
	database, err := db.Open(cfg.DSN)
	if err != nil {
		return nil, err
	}
	qr := db.NewQuerier(database)

	c, err := cache.NewRedisCache(cfg.Redis["host"], cfg.Redis["password"])
	if err != nil {
		log.Warn(err)
	}

	app := &Application{qr, c, cfg}
	return app, err
}

func (app *Application) Run() {
}
