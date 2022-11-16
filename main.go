package main

import (
	"github.com/Big-Sh4rk/Inventory-Project/settings"
	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(
			settings.New,
		),
		fx.Invoke(),
	)

	app.Run()
}
