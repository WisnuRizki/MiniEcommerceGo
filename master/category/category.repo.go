package category

import (
	"miniecommerce.wisnu.net/database"
	"miniecommerce.wisnu.net/modules"
)

type Category modules.Category

func (category *Category) Create(u *Category) error{
	result := database.DB.Create(&u)
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (category *Category) IsCategoryExist(name string) error{
	result := database.DB.Where(&Category{Name: name}).Find(&category)
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (category *Category) Delete(id int) error {
	result := database.DB.Where(&Category{ID: uint(id)}).Delete(&category)
	if result.RowsAffected == 0{
		return result.Error
	}

	return nil
}

func (category *Category) GetAll() (*[]Category){
	data := []Category{}
	result := database.DB.Find(&data)
	if result.RowsAffected == 0 {
		return nil
	}

	return &data
}

func (category *Category) GetById(id int) (*Category){
	result := database.DB.Where(&Category{ID: uint(id)}).Find(&category)
	if result.RowsAffected == 0 {
		return nil
	}

	return category
}

func (category *Category) Update(id int, c *Category) error {
	result := database.DB.Model(&category).Where(&Category{ID: uint(id)}).Updates(&c)
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}