package main

import (
	"kolo_marvel_project/config"
	"kolo_marvel_project/internal/server"
	"kolo_marvel_project/internal/server/handler"
	"kolo_marvel_project/pkg/dummy"
	"kolo_marvel_project/utils/initialize"

	"go.uber.org/fx"
)

func serverRun() {
	app := fx.New(
		fx.Provide(
			// postgresql
			initialize.NewKoloMarvelprojecteDB,
		),
		config.Module,
		initialize.Module,
		server.Module,
		handler.Module,
		dummy.Module,
	)
	// Run app forever
	app.Run()
}
