package history

import (
	"miniecommerce.wisnu.net/database"
	"miniecommerce.wisnu.net/modules"
)

type History modules.History

func (history *History) CreateHistory(h []History) (error) {
	result := database.DB.Create(&h)
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (history *History) GetHistoryByUserId(userId int) (*[]History){
	data := []History{}
	result := database.DB.Where(&History{UserId: userId}).Preload("Product").Find(&data)
	if result.RowsAffected == 0 {
		return nil
	}

	return &data
}

func (history *History) DeleteByUserId(userId int) error {
	result := database.DB.Where(&History{UserId: userId}).Delete(&history)
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}