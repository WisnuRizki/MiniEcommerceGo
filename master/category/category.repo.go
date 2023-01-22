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