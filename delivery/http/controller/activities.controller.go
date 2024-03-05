package controller

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"todolist-api/delivery/http/dto/request"
	"todolist-api/domain/usecase"
	"todolist-api/shared/util"
)

type ActivitiesController interface {
	CreateActivityGroup(ctx *fiber.Ctx) error
	FindAllActivityGroup(ctx *fiber.Ctx) error
	FindActivityGroupById(ctx *fiber.Ctx) error
	UpdateActivityGroup(ctx *fiber.Ctx) error
	DeleteActivityGroup(ctx *fiber.Ctx) error
}

type activitiesControllerImpl struct {
	activitiesUsecase usecase.ActivitiesUsecase
}

func NewActivitiesController(activitiesUsecase usecase.ActivitiesUsecase) ActivitiesController {
	return &activitiesControllerImpl{activitiesUsecase: activitiesUsecase}
}

func (a activitiesControllerImpl) CreateActivityGroup(ctx *fiber.Ctx) error {
	requestBody := new(request.CreateActivityGroupRequest)
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

	result, err := a.activitiesUsecase.CreateActivityGroup(ctx.UserContext(), requestBody)
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

func (a activitiesControllerImpl) FindAllActivityGroup(ctx *fiber.Ctx) error {
	result, err := a.activitiesUsecase.FindAllActivityGroup(ctx.UserContext())
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

func (a activitiesControllerImpl) FindActivityGroupById(ctx *fiber.Ctx) error {
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

	result, err := a.activitiesUsecase.FindActivityGroupById(ctx.UserContext(), id)
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

func (a activitiesControllerImpl) UpdateActivityGroup(ctx *fiber.Ctx) error {
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

	requestBody := new(request.CreateActivityGroupRequest)
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

	result, err := a.activitiesUsecase.UpdateActivityGroup(ctx.UserContext(), id, requestBody)
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

func (a activitiesControllerImpl) DeleteActivityGroup(ctx *fiber.Ctx) error {
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

	err = a.activitiesUsecase.DeleteActivityGroup(ctx.UserContext(), id)
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
