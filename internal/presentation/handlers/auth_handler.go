package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hamwiwatsapon/train-booking-go/internal/application/services"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	type RegisterRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	type RegisterReponse struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	_, err := h.service.RegisterUser(req.Email, req.Password, req.Role)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	token, refresh, err := h.service.LoginUser(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	response := RegisterReponse{
		Token:        token,
		RefreshToken: refresh,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type LoginResponse struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	token, refresh, err := h.service.LoginUser(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	response := LoginResponse{
		Token:        token,
		RefreshToken: refresh,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	type RefreshRequest struct {
		RefreshToken string `json:"refresh_token"`
	}

	type RefreshResponse struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}

	var req RefreshRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	token, refresh, err := h.service.GetNewToken(req.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	response := RefreshResponse{
		Token:        token,
		RefreshToken: refresh,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h *AuthHandler) CheckUser(c *fiber.Ctx) error {
	type CheckUserRequest struct {
		Email string `json:"email"`
	}

	var req CheckUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	err := h.service.CheckUserExist(req.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "user exists"})
}

func (h *AuthHandler) OTPLogin(c *fiber.Ctx) error {
	type OTPLoginRequest struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}

	type OTPLoginResponse struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}

	var req OTPLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	token, refresh, err := h.service.OTPLogin(req.Email, req.OTP)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	response := OTPLoginResponse{
		Token:        token,
		RefreshToken: refresh,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
