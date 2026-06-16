package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/yaduvamsi/user-age-api/internal/service"

	"github.com/yaduvamsi/user-age-api/internal/models"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(
	service *service.UserService,
) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetUser(
	c *fiber.Ctx,
) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	user, err := h.service.GetUser(
		c.Context(),
		int32(id),
	)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.JSON(user)
}
func (h *UserHandler) CreateUser(
	c *fiber.Ctx,
) error {

	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}
	if err := validate.Struct(req); err != nil {
	return c.Status(400).JSON(fiber.Map{
		"error": err.Error(),
	})
}


	user, err := h.service.CreateUser(
	c.Context(),
	req,
)

if err != nil {
	return c.Status(400).JSON(
		fiber.Map{
			"error": err.Error(),
		},
	)
}

return c.Status(201).JSON(user)
}
func (h *UserHandler) ListUsers(
	c *fiber.Ctx,
) error {

	users, err := h.service.ListUsers(
		c.Context(),
	)

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"error": "failed to fetch users",
			},
		)
	}

	return c.JSON(users)
}
func (h *UserHandler) DeleteUser(
	c *fiber.Ctx,
) error {

	id, err := strconv.Atoi(
		c.Params("id"),
	)

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "invalid id",
			},
		)
	}

	err = h.service.DeleteUser(
		c.Context(),
		int32(id),
	)

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"error": "failed to delete user",
			},
		)
	}

	return c.JSON(
		fiber.Map{
			"message": "user deleted",
		},
	)
}

func (h *UserHandler) UpdateUser(
	c *fiber.Ctx,
) error {

	id, err := strconv.Atoi(
		c.Params("id"),
	)

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "invalid id",
			},
		)
	}

	var req models.UpdateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "invalid request",
			},
		)
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	user, err := h.service.UpdateUser(
		c.Context(),
		int32(id),
		req,
	)

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(user)
}