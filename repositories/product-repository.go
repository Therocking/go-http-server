package repositories

import (
	"githup.com/Therocking/go-http/db"
	"githup.com/Therocking/go-http/models"
)

func GetAllProducts(limit, skip int) ([]models.Product, error) {
	var products []models.Product
	db.Db.Limit(limit).Offset(skip).Find(&products)

	return products, nil
}

func GetProduct(productId int64) (models.Product, error) {
	var product models.Product
	productFound := db.Db.First(&product, productId)
	if err := productFound.Error; err != nil {
		return models.Product{}, err
	}

	db.Db.Model(&product).Association("UserId").Find(&product.UserId)

	return product, nil
}

func CreateProduct(product models.Product) (models.Product, error) {
	createdProduct := db.Db.Create(&product)
	if err := createdProduct.Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func UpdateProduct(id int64, product models.Product) (models.Product, error) {
	productFound := db.Db.First(&product, id)
	if err := productFound.Error; err != nil {
		return models.Product{}, err
	}

	db.Db.Save(&product)

	return product, nil
}

func DeleteProduct(productId int64) (models.Product, error) {
	var product models.Product
	productFound := db.Db.First(&product, productId)
	if err := productFound.Error; err != nil {
		return models.Product{}, nil
	}

	db.Db.Delete(&product)
	return product, nil
}
