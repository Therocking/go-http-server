package services

import (
	"fmt"

	"githup.com/Therocking/go-http/models"
	"githup.com/Therocking/go-http/repositories"
)

func GetAllProducts() ([]models.Product, error) {
	products, err := repositories.GetAllProducts()
	if err != nil {
		return []models.Product{}, err
	}

	return products, nil
}

func GetProduct(id int64) (models.Product, error) {
	product, err := repositories.GetProduct(id)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func CreateProduct(product models.Product) (string, error) {

	result, err := repositories.CreateProduct(product)
	if err != nil {
		return "", err
	}

	id, _ := result.LastInsertId()

	return fmt.Sprintf("models.Product with id: %v was added", id), nil
}

func DeleteProduct(id int64) (models.Product, error) {
	product, err := repositories.DeleteProduct(id)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}
