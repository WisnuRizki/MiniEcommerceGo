package seller

import (
	"errors"

	"miniecommerce.wisnu.net/database"
	"miniecommerce.wisnu.net/modules"
)

type Seller modules.Seller

func (seller *Seller) CreateSeller(u *Seller) error {
	result := database.DB.Create(&u)

	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (seller *Seller) CheckDataSellerExists(email string) (*Seller,int64){
	data := Seller{}

	result := database.DB.Where(&Seller{Email: email}).Find(&data)

	if result.RowsAffected == 0 {
		return nil,0
	}

	return &data,result.RowsAffected
}

func (seller *Seller) CheckDataSellerExistById(id uint) (*Seller,int64){
	data := Seller{}

	result := database.DB.Where(&Seller{ID: id}).Find(&data)

	if result.RowsAffected == 0 {
		return nil,0
	}

	return &data,result.RowsAffected
}

func (seller *Seller) LoginUser(email string,password string) error {
	data := Seller{}

	result := database.DB.Where(&Seller{Email: email,Password: password}).Find(&data)
	if result.RowsAffected == 0 {
		return errors.New("not Found")
	}

	return nil

}


