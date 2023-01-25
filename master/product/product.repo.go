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

func (product *Product) GetProductById(id uint) (*Product,error){
	data := Product{}
	result := database.DB.Where(&Product{ID: id}).Find(&data)
	if result.RowsAffected == 0 {
		return nil,result.Error
	}

	return &data,nil
}

func (product *Product) UpdateProductStock(
	id uint,
	amount int,
	initialAmount int,
	types string) error {
		var totalAmount int
		if types == "sum"{
			totalAmount = initialAmount + amount
		}else{
			totalAmount = initialAmount - amount
		}

		result := database.DB.Model(&product).Where(&Product{ID: id}).Updates(Product{Quantity: totalAmount})
		if result.RowsAffected == 0 {
			return result.Error
		}

		return nil


}

func (product *Product) GetAllProductSeller(sellerId int)(*[]Product,error){
	data := []Product{}

	result := database.DB.Where(&Product{SellerId: sellerId}).Preload("Seller").Preload("Category").Find(&data)
	if result.RowsAffected == 0 {
		return nil,result.Error
	}

	return &data,nil
}

func (product *Product) Delete(sellerId int, id uint) error {
	result := database.DB.Where(&Product{ID: id,SellerId: sellerId}).Delete(&product)
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}