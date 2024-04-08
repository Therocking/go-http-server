package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"githup.com/Therocking/go-http/controllers"
	"githup.com/Therocking/go-http/db"
	"githup.com/Therocking/go-http/models"
)

func main() {
	if os.Getenv("ENV") == "dev" {
		envFile := ".env"

		err := godotenv.Load(envFile)
		if err != nil {
			panic(err)
		}
	}

	db.DbConnection()
	db.Db.AutoMigrate(&models.User{})
	db.Db.AutoMigrate(&models.Product{})

	app := fiber.New()

	usersGroup := app.Group("/users")
	productGroup := app.Group("/products")

	app.Use(logger.New())

	app.Get("", func(c *fiber.Ctx) error {
		return c.SendFile("./public/welcome.html")
	})
	app.Get("hello", HelloHandler)

	/*Users routes*/
	usersGroup.Get("", controllers.ShowUsers)
	usersGroup.Get(":id", controllers.ShowUser)
	usersGroup.Post("", controllers.CreateUser)
	usersGroup.Put(":id", controllers.UpdateUser)
	usersGroup.Delete(":id", controllers.DeleteUser)

	/*Products routes*/
	productGroup.Get("", controllers.GetAllProducts)
	productGroup.Get(":id", controllers.GetProduct)
	productGroup.Post("", controllers.CreateProduct)
	productGroup.Delete(":id", controllers.DeleteProduct)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if os.Getenv("ENV") == "prod" {
		port = ":2000"
	}
	app.Listen(port)
}

func HelloHandler(c *fiber.Ctx) error {

	return c.SendString("Hello World.")
}
