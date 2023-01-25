package user

import (
	"errors"

	"miniecommerce.wisnu.net/database"
	"miniecommerce.wisnu.net/modules"
)

type User modules.User

func (user *User) CreateUser(u *User) error {
	result := database.DB.Create(&u)

	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (user *User) CheckDataByEmail(email string) (*User,int64){
	data := User{}

	result := database.DB.Where(&User{Email: email}).Find(&data)

	if result.RowsAffected == 0 {
		return nil,0
	}

	return &data,result.RowsAffected
}

func (user *User) LoginUser(email string,password string) error {
	data := User{}

	result := database.DB.Where(&User{Email: email,Password: password}).Find(&data)
	if result.RowsAffected == 0 {
		return errors.New("Not Found")
	}

	return nil

}

func (user *User) Delete(id int) error {
	result := database.DB.Where(&User{ID: uint(id)}).Delete(&user)
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}