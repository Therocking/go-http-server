package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"githup.com/Therocking/go-http/models"
	"githup.com/Therocking/go-http/services"
)

func GetAllProducts(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))

	products, err := services.GetAllProducts(limit, skip)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	return c.JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	productId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	product, err := services.GetProduct(productId)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	product := models.Product{}
	if err := c.BodyParser(&product); err != nil {
		return err
	}

	newProduct, err := services.CreateProduct(product)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(newProduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	productId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	product, err := services.DeleteProduct(productId)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	return c.JSON(product)
}
