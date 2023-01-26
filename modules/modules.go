package modules

import (
	"time"
)


type User struct {
	ID        	uint      		`json:"id"`
	FirstName 	string    		`json:"first_name"`
	LastName  	string    		`json:"last_name"`
	Email     	string    		`json:"email"`
	Password  	string    		`json:"password"`
	CreatedAt 	time.Time 		`json:"created_at"`
	UpdatedAt 	time.Time 		`json:"updated_at"`
	Transaction *[]Transaction  `json:"transaction"`
}

type Admin struct {
	ID        		uint      	`json:"id"`
	Username     	string    	`json:"username" `
	Password  		string    	`json:"password" `
}

type Seller struct {
	ID        	uint      	`json:"id"`
	Name  		string    	`json:"name"`
	NoHp 		string		`json:"no_hp"`
	Email     	string    	`json:"email" `
	Password  	string    	`json:"password"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
	Products 	[]Product	`json:"product"`
}

type Balance struct {
	ID        	uint      	`json:"id"`
	UserId		int 		`json:"user_id"`
	Amount 		int64 		`json:"amount"`	
	User 		*User 		`json:"user"`
}

type Product struct {
	ID        		uint      	`json:"id"`
	Name 			string 		`json:"name"`
	SellerId 		int 		`json:"seller_id"`
	CategoryId		int			`json:"category_id"`
	Price 			int64		`json:"price"`
	Quantity		int 		`json:"quantity"`
	Seller			*Seller 		`json:"seller"`
	Category 		*Category	`json:"category"` 
}

type Category struct {
	ID        	uint      	`json:"id"`
	Name 		string 		`json:"name"`
}

type Transaction struct {
	ID        		uint      	`json:"id"`
	UserId			int 		`json:"user_id"`
	Status 			string 		`json:"status"`
	TransNumber 	string  	`json:"trans_number"`
	TotalPayment	int64 		`json:"total_price"`
	CreatedAt 		time.Time 	`json:"created_at"`
	UpdatedAt 		time.Time 	`json:"updated_at"`
	User 			*User  
	History 		[]History  `json:"history" gorm:"foreignKey:TransId"` 
}

type History struct {
	ID        	uint      	 `json:"id"`
	UserId		int 		 `json:"user_id"`
	ProductId 	int 		 `json:"product_id"`
	Quantity	int  		 `json:"quantity"`
	TotalPrice 	int64 		 `json:"total_price"`
	TransId		int			 `json:"trans_id"`
	Status 		string 		`json:"status"`
	Product 	*Product	 `json:"product"`
	Transaction *Transaction  `json:"transaction" gorm:"foreignKey:TransId"`
}



