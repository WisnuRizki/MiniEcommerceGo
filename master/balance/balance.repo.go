package balance

import (
	"miniecommerce.wisnu.net/database"
	"miniecommerce.wisnu.net/modules"
)

type Balance modules.Balance

func (balance *Balance) CreateBalance(userId int) error {
	balance.UserId = userId
	balance.Amount = 0

	result := database.DB.Create(balance)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
	
}