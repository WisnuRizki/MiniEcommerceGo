package product

import (
	"miniecommerce.wisnu.net/database"
	"miniecommerce.wisnu.net/modules"
)

type Product modules.Product

func (product *Product) CreateProduct(p []Product)error {
	result := database.DB.Create(&p)
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}