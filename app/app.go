package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"todolist-api/delivery/http/controller"
	"todolist-api/delivery/http/router"
	"todolist-api/domain/repository"
	"todolist-api/domain/usecase"
	"todolist-api/shared/util"
)

type BootstrapConfig struct {
	DB          *gorm.DB
	App         *fiber.App
	Validate    *validator.Validate
	ViperConfig util.Config
}

func Bootstrap(config *BootstrapConfig) {

	// setup repositories
	activitiesRepository := repository.NewActivitiesRepository()
	todosRepository := repository.NewTodosRepository()

	// setup usecases
	activitiesUsecase := usecase.NewActivitiesUsecase(config.DB, config.Validate, activitiesRepository)
	todosUsecase := usecase.NewTodosUsecase(config.DB, config.Validate, todosRepository, activitiesRepository)

	// setup controller
	activitiesController := controller.NewActivitiesController(activitiesUsecase)
	todosController := controller.NewTodosController(todosUsecase)

	routeConfig := router.RouteConfig{
		App:                  config.App,
		ActivitiesController: activitiesController,
		TodosController:      todosController,
	}

	routeConfig.Setup()
}
