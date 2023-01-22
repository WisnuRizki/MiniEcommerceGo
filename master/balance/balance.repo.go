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

func (balance *Balance) CheckBalanceByUserId(userId int)(*Balance,int){
	data := Balance{}
	result := database.DB.Where(&Balance{UserId: userId}).Find(&data)
	if result.RowsAffected == 0 {
		return nil,0
	}

	return &data,int(result.RowsAffected)
}

func (balance *Balance) AddBalance(userId int,amount int64,initialAmount int64) error {
	totalAmount := amount + initialAmount
	result := database.DB.Model(&balance).Where(&Balance{UserId: userId}).Updates(Balance{Amount: totalAmount})
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (balance *Balance) ReduceBalance(userId int,amount int64,initialAmount int64) error {
	totalAmount := initialAmount -amount
	result := database.DB.Model(&balance).Where(&Balance{UserId: userId}).Updates(Balance{Amount: totalAmount})
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}