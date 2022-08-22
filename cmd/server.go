package main

import (
	"kolo_marvel_project/config"
	"kolo_marvel_project/internal/server"
	"kolo_marvel_project/internal/server/handler"
	"kolo_marvel_project/pkg/cache"
	"kolo_marvel_project/pkg/marvel"
	"kolo_marvel_project/utils/initialize"

	"go.uber.org/fx"
)

func serverRun() {
	app := fx.New(
		fx.Provide(
			// postgresql
			// initialize.NewKoloMarvelprojecteDB,
			initialize.NewRedisWorker,
		),
		config.Module,
		initialize.Module,
		server.Module,
		handler.Module,
		marvel.Module,
		cache.Module,
	)
	// Run app forever
	app.Run()
}
