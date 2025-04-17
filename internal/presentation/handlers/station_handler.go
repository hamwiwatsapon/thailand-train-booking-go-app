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

// StationType Handler
func (h *TrainHandler) GetStationTypes(c *fiber.Ctx) error {
	trainStations, err := h.services.GetTrainStationTypes()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch train station types",
		})
	}
	return c.JSON(trainStations)
}

func (h *TrainHandler) CreateStationType(c *fiber.Ctx) error {
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

	userID, ok := c.Locals("user").(uint)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid or missing user ID",
		})
	}

	role := c.Locals("role").(string)
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}

	trainType.Code = req.Code
	trainType.Name = req.Name
	trainType.ModifyBy = userID

	createdStation, err := h.services.CreateTrainStationType(trainType)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create train station type",
			"details": err.Error(), // Optional: Include error details for debugging
		})
	}
	return c.Status(fiber.StatusCreated).JSON(createdStation)
}

func (h *TrainHandler) UpdateStationType(c *fiber.Ctx) error {
	type updateTrainTypeRequest struct {
		Code string `json:"code" validate:"required"`
		Name string `json:"name" validate:"required"`
	}

	var req updateTrainTypeRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	role := c.Locals("role").(string)
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
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

	updatedStation, err := h.services.UpdateTrainStationType(trainType)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update train station type",
		})
	}
	return c.JSON(updatedStation)
}

func (h *TrainHandler) DeleteStationType(c *fiber.Ctx) error {
	type deleteTrainTypeRequest struct {
		Code string `json:"code" validate:"required"`
	}

	var req deleteTrainTypeRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	role := c.Locals("role").(string)
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}

	err := h.services.DeleteTrainStationType(req.Code)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete train station type",
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// Station Handler
func (h *TrainHandler) GetStations(c *fiber.Ctx) error {
	// Helper function to extract query parameters
	getQueryParam := func(key string) string {
		value := c.Query(key)
		if value != "" {
			return value
		}
		return ""
	}

	// Build filters using a helper function
	filters := map[string]interface{}{
		"station_type_code": getQueryParam("station_type_code"),
		"name":              getQueryParam("name"),
		"postal_code":       getQueryParam("postal_code"),
		"province":          getQueryParam("province"),
		"district":          getQueryParam("district"),
		"sub_district":      getQueryParam("sub_district"),
	}

	// Fetch train stations
	trainStations, err := h.services.GetTrainStations(filters)
	if err != nil {
		// Log the error (if logging is implemented)
		// log.Errorf("Failed to fetch train stations: %v", err)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to fetch train stations",
			"details": err.Error(), // Optional: Include error details for debugging
		})
	}

	return c.JSON(trainStations)
}

func (h *TrainHandler) BulkCreateStation(c *fiber.Ctx) error {
	type createTrainStationRequest struct {
		Name            string `json:"name" validate:"required"`
		StationTypeCode string `json:"station_type_code" validate:"required"`
		PostalCode      string `json:"postal_code"`
		Province        string `json:"province"`
		District        string `json:"district"`
		SubDistrict     string `json:"sub_district"`
		Latitude        string `json:"latitude"`
		Longitude       string `json:"longitude"`
	}

	var req []createTrainStationRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	role := c.Locals("role").(string)
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}

	userID, ok := c.Locals("user").(uint)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid or missing user ID",
		})
	}

	trainStations := make([]entities.TrainStation, len(req))
	for i, station := range req {
		trainStations[i] = entities.TrainStation{
			Name:            station.Name,
			Province:        station.Province,
			District:        station.District,
			SubDistrict:     station.SubDistrict,
			PostalCode:      station.PostalCode,
			Latitude:        station.Latitude,
			Longitude:       station.Longitude,
			StationTypeCode: station.StationTypeCode,
			ModifyBy:        userID,
		}
	}
	createdStations, err := h.services.BulkCreateTrainStation(trainStations)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create train stations",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(createdStations)
}
