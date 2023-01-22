package admin

import (
	"errors"

	"miniecommerce.wisnu.net/database"
	"miniecommerce.wisnu.net/modules"
)

type Admin modules.Admin

func (admin *Admin) CreateAdmin(a *Admin) error {
	result := database.DB.Create(&a)

	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (admin *Admin) CheckDataByUsername(username string) (*Admin,int64){
	data := Admin{}

	result := database.DB.Where(&Admin{Username: username}).Find(&data)

	if result.RowsAffected == 0 {
		return nil,0
	}

	return &data,result.RowsAffected
}

func (admin *Admin) CheckDataById(id uint) error{
	data := Admin{}

	result := database.DB.Where(&Admin{ID: id}).Find(&data)

	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (admin *Admin) LoginUser(username string,password string) error {
	data := Admin{}

	result := database.DB.Where(&Admin{Username: username,Password: password}).Find(&data)
	if result.RowsAffected == 0 {
		return errors.New("not Found")
	}

	return nil

}