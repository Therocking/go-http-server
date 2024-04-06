package services

import (
	"githup.com/Therocking/go-http/models"
	"githup.com/Therocking/go-http/repositories"
)

func GetAllProducts(limit, skip int) ([]models.Product, error) {
	products, err := repositories.GetAllProducts(limit, skip)
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

func CreateProduct(product models.Product) (models.Product, error) {

	product, err := repositories.CreateProduct(product)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func DeleteProduct(id int64) (models.Product, error) {
	product, err := repositories.DeleteProduct(id)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}
