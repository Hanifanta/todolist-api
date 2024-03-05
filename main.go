package main

import (
	"log"
	"todolist-api/app"
	"todolist-api/shared/util"
)

func main() {
	viperConfig, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db := app.NewDatabaseConnection(viperConfig.DBDsn)
	fiber := app.NewFiber(viperConfig)
	validate := app.NewValidator()

	app.Bootstrap(&app.BootstrapConfig{
		DB:          db,
		App:         fiber,
		Validate:    validate,
		ViperConfig: viperConfig,
	})

	err = fiber.Listen(":" + viperConfig.PortApp)
	if err != nil {
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}
}
