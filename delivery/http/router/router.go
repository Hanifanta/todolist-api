package router

import (
	"github.com/gofiber/fiber/v2"
	"todolist-api/delivery/http/controller"
)

type RouteConfig struct {
	App                  *fiber.App
	AuthMiddleware       fiber.Handler
	ActivitiesController controller.ActivitiesController
	TodosController      controller.TodosController
}

func (c *RouteConfig) Setup() {
	c.SetupRoute()
}

func (c *RouteConfig) SetupRoute() {
	// Activities
	c.App.Get("/activity-groups", c.ActivitiesController.FindAllActivityGroup)
	c.App.Get("/activity-groups/:id", c.ActivitiesController.FindActivityGroupById)
	c.App.Post("/activity-groups", c.ActivitiesController.CreateActivityGroup)
	c.App.Patch("/activity-groups/:id", c.ActivitiesController.UpdateActivityGroup)
	c.App.Delete("/activity-groups/:id", c.ActivitiesController.DeleteActivityGroup)

	// Todos
	c.App.Get("/todo-items", c.TodosController.FindAllTodoItem)
	c.App.Get("/todo-items/:id", c.TodosController.FindTodoItemById)
	c.App.Post("/todo-items", c.TodosController.CreateTodoItem)
	c.App.Patch("/todo-items/:id", c.TodosController.UpdateTodoItem)
	c.App.Delete("/todo-items/:id", c.TodosController.DeleteTodoItem)
}
