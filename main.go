package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"githup.com/Therocking/go-http/controllers"
	"githup.com/Therocking/go-http/db"
)

func main() {
	// env := os.Getenv("ENV")
	// print(env)
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db.DbConnection()

	app := fiber.New()

	usersGroup := app.Group("/users")
	productGroup := app.Group("/products")

	app.Use(logger.New())

	app.Get("hello", controllers.HelloHandler)

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
	app.Listen(port)
}
