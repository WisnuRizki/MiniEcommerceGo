package transaction

import (
	"miniecommerce.wisnu.net/database"
	"miniecommerce.wisnu.net/modules"
)

type Transaction modules.Transaction

func GetTrans() *Transaction {
	tra := Transaction{}
	result := database.DB.Where(&Transaction{ID: 1}).Preload("History").Preload("User").Find(&tra)
	if result.RowsAffected == 0 {
		return nil
	}

	return &tra
}

func (transaction *Transaction) Create(t *Transaction) (int,error) {
	result := database.DB.Create(&t)
	if result.RowsAffected == 0 {
		return 0,result.Error
	}


	return int(transaction.ID),nil
}

func (transaction *Transaction) UpdateStatus(status string,id int) error {
	result := database.DB.Model(&transaction).Where(&Transaction{ID: uint(id)}).Updates(Transaction{Status: status})
		if result.RowsAffected == 0 {
			return result.Error
		}

		return nil
}