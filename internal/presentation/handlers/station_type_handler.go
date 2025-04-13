package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hamwiwatsapon/train-booking-go/internal/application/services"
	"github.com/hamwiwatsapon/train-booking-go/internal/domain/entities"
)

type TrainHandler struct {
	services *services.TrainService
}

func NewTrainHandler(services *services.TrainService) *TrainHandler {
	return &TrainHandler{
		services: services,
	}
}

func (h *TrainHandler) GetTrainTypes(c *fiber.Ctx) error {
	trainStations, err := h.services.GetTrainStationTypes()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch train station types",
		})
	}
	return c.JSON(trainStations)
}

func (h *TrainHandler) CreateTrainType(c *fiber.Ctx) error {
	type createTrainTypeRequest struct {
		Code string `json:"code" validate:"required"`
		Name string `json:"name" validate:"required"`
	}

	var req createTrainTypeRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	var trainType entities.StationType

	userID := c.Locals("user")

	trainType.Code = req.Code
	trainType.Name = req.Name
	if id, ok := userID.(uint); ok {
		trainType.ModifyBy = id
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	createdStation, err := h.services.CreateTrainStationType(trainType)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create train station type",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(createdStation)
}
