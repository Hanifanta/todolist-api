package controller

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"todolist-api/delivery/http/dto/request"
	"todolist-api/domain/usecase"
	"todolist-api/shared/util"
)

type TodosController interface {
	CreateTodoItem(ctx *fiber.Ctx) error
	FindAllTodoItem(ctx *fiber.Ctx) error
	FindTodoItemById(ctx *fiber.Ctx) error
	UpdateTodoItem(ctx *fiber.Ctx) error
	DeleteTodoItem(ctx *fiber.Ctx) error
}

type todosControllerImpl struct {
	todosUsecase usecase.TodosUsecase
}

func NewTodosController(todosUsecase usecase.TodosUsecase) TodosController {
	return &todosControllerImpl{todosUsecase: todosUsecase}
}

func (t todosControllerImpl) CreateTodoItem(ctx *fiber.Ctx) error {
	requestBody := new(request.CreateTodoRequest)
	err := ctx.BodyParser(requestBody)
	if err != nil {
		resp, statusCode := util.ConstructBaseResponse(
			util.BaseResponse{
				Code:   fiber.StatusBadRequest,
				Status: err.Error(),
			},
		)
		return ctx.Status(statusCode).JSON(resp)
	}

	result, err := t.todosUsecase.CreateTodoItem(ctx.UserContext(), requestBody)
	if err != nil {
		resp, statusCode := util.ConstructBaseResponse(
			util.BaseResponse{
				Code:   fiber.StatusBadRequest,
				Status: err.Error(),
			},
		)
		return ctx.Status(statusCode).JSON(resp)
	}

	resp, statusCode := util.ConstructBaseResponse(
		util.BaseResponse{
			Code:   fiber.StatusCreated,
			Status: "Success",
			Data:   result,
		},
	)
	return ctx.Status(statusCode).JSON(resp)
}

func (t todosControllerImpl) FindAllTodoItem(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Query("activity_group_id", "0"))
	result, err := t.todosUsecase.FindAllTodoItem(ctx.UserContext(), id)
	if err != nil {
		resp, statusCode := util.ConstructBaseResponse(
			util.BaseResponse{
				Code:   fiber.StatusBadRequest,
				Status: err.Error(),
			},
		)
		return ctx.Status(statusCode).JSON(resp)
	}

	resp, statusCode := util.ConstructBaseResponse(
		util.BaseResponse{
			Code:   fiber.StatusOK,
			Status: "Success",
			Data:   result,
		},
	)
	return ctx.Status(statusCode).JSON(resp)
}

func (t todosControllerImpl) FindTodoItemById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		resp, statusCode := util.ConstructBaseResponse(
			util.BaseResponse{
				Code:   fiber.StatusBadRequest,
				Status: err.Error(),
			},
		)
		return ctx.Status(statusCode).JSON(resp)
	}

	result, err := t.todosUsecase.FindTodoItemById(ctx.UserContext(), id)
	if err != nil {
		resp, statusCode := util.ConstructBaseResponse(
			util.BaseResponse{
				Code:   fiber.StatusBadRequest,
				Status: err.Error(),
			},
		)
		return ctx.Status(statusCode).JSON(resp)
	}

	resp, statusCode := util.ConstructBaseResponse(
		util.BaseResponse{
			Code:   fiber.StatusOK,
			Status: "Success",
			Data:   result,
		},
	)
	return ctx.Status(statusCode).JSON(resp)
}

func (t todosControllerImpl) UpdateTodoItem(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		resp, statusCode := util.ConstructBaseResponse(
			util.BaseResponse{
				Code:   fiber.StatusBadRequest,
				Status: err.Error(),
			},
		)
		return ctx.Status(statusCode).JSON(resp)
	}

	requestBody := new(request.CreateTodoRequest)
	err = ctx.BodyParser(requestBody)
	if err != nil {
		resp, statusCode := util.ConstructBaseResponse(
			util.BaseResponse{
				Code:   fiber.StatusBadRequest,
				Status: err.Error(),
			},
		)
		return ctx.Status(statusCode).JSON(resp)
	}

	result, err := t.todosUsecase.UpdateTodoItem(ctx.UserContext(), id, requestBody)
	if err != nil {
		resp, statusCode := util.ConstructBaseResponse(
			util.BaseResponse{
				Code:   fiber.StatusBadRequest,
				Status: err.Error(),
			},
		)
		return ctx.Status(statusCode).JSON(resp)
	}

	resp, statusCode := util.ConstructBaseResponse(
		util.BaseResponse{
			Code:   fiber.StatusOK,
			Status: "Success",
			Data:   result,
		},
	)
	return ctx.Status(statusCode).JSON(resp)
}

func (t todosControllerImpl) DeleteTodoItem(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		resp, statusCode := util.ConstructBaseResponse(
			util.BaseResponse{
				Code:   fiber.StatusBadRequest,
				Status: err.Error(),
			},
		)
		return ctx.Status(statusCode).JSON(resp)
	}

	err = t.todosUsecase.DeleteTodoItem(ctx.UserContext(), id)
	if err != nil {
		resp, statusCode := util.ConstructBaseResponse(
			util.BaseResponse{
				Code:   fiber.StatusBadRequest,
				Status: err.Error(),
			},
		)
		return ctx.Status(statusCode).JSON(resp)
	}

	resp, statusCode := util.ConstructBaseResponse(
		util.BaseResponse{
			Code:   fiber.StatusOK,
			Status: "Success",
		},
	)
	return ctx.Status(statusCode).JSON(resp)
}
