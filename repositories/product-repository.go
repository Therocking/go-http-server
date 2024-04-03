package repositories

import (
	"database/sql"

	"githup.com/Therocking/go-http/db"
	"githup.com/Therocking/go-http/models"
)

func GetAllProducts() ([]models.Product, error) {
	rows, err := db.Db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CreatedBy)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func GetProduct(productId int64) (models.Product, error) {
	query := "SELECT * FROM products WHERE id = ?"

	var product models.Product
	err := db.Db.QueryRow(query, productId).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CreatedBy)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func CreateProduct(product models.Product) (sql.Result, error) {
	query := "INSERT INTO products (name, description, price, createdBy) VALUES (?, ?, ?, ?)"
	result, err := db.Db.Exec(query, product.Name, product.Description, product.Price, product.CreatedBy)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateProduct(product models.Product) (models.Product, error) {
	query := "UPDATE products SET name = ?, description = ?, price = ?, createdBy = ? WHERE id = ?"
	err := db.Db.QueryRow(query, product.Name, product.Description, product.Price, product.CreatedBy, product.Id).Scan(
		&product.Id, &product.Name, &product.Description, &product.Price, &product.CreatedBy)

	if err != nil {
		return models.Product{}, err

	}

	return product, nil
}

func DeleteProduct(productId int64) (models.Product, error) {
	query := "DELETE FROM products WHERE id = ?"

	var product models.Product
	err := db.Db.QueryRow(query, productId).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CreatedBy)

	if err != nil {
		return models.Product{}, err

	}

	return product, nil
}
