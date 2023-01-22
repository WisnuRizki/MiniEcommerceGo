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