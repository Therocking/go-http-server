package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"githup.com/Therocking/go-http/models"
	"githup.com/Therocking/go-http/services"
)

func HelloHandler(c *fiber.Ctx) error {

	return c.SendString("Hello World.")
}

func ShowUsers(c *fiber.Ctx) error {
	users, err := services.GetAllUsers()
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	return c.JSON(users)
}

func ShowUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	user, err := services.GetUser(userId)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := models.User{}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	userSaved, err := services.CreateUser(user)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(userSaved)
}

func UpdateUser(c *fiber.Ctx) error {
	user := models.User{}

	id := c.Params("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.Id = userId

	userUpdated, err := services.UpdateUser(user)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	return c.JSON(userUpdated)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	user, err := services.DeleteUser(userId)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	return c.JSON(user)
}
