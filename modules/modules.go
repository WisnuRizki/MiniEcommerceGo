package modules

import (
	"time"
)

type Role struct {
	ID        	uint      	`json:"id"`
	Type 		string 		`json:"type"`
}

type User struct {
	ID        	uint      	`json:"id"`
	FirstName 	string    	`json:"first_name"`
	LastName  	string    	`json:"last_name"`
	Email     	string    	`json:"email"`
	Password  	string    	`json:"password"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
	RoleId 		int 		`json:"role"`
	Role 		Role 		`json:"user_role"`
	History 	[]History	`json:"history"`
}

type Admin struct {
	ID        		uint      	`json:"id"`
	Username     	string    	`json:"username"`
	Password  		string    	`json:"password"`
	RoleId 			int 		`json:"role"`
	Role 			Role 		`json:"user_role"`
}

type Seller struct {
	ID        	uint      	`json:"id"`
	Name  		string    	`json:"last_name"`
	NoHp 		string		`json:"no_hp"`
	Email     	string    	`json:"email"`
	Password  	string    	`json:"password"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
	RoleId 		int 		`json:"role"`
	Role 		Role 		`json:"user_role"`
	Products 	[]Product	`json:"product"`
}

type Balance struct {
	ID        	uint      	`json:"id"`
	UserId		int 		`json:"user_id"`
	Amount 		int64 		`json:"amount"`	
	User 		User 		`json:"user"`
}

type Product struct {
	ID        		uint      	`json:"id"`
	Name 			string 		`json:"name"`
	SellerId 		int 		`json:"seller_id"`
	CategoryId		int			`json:"category_id"`
	Price 			int64		`json:"price"`
	Quantity		int 		`json:"quantity"`
	Seller			Seller 		`json:"seller"`
	Category 		Category	`json:"category"` 
}

type Category struct {
	ID        	uint      	`json:"id"`
	Name 		string 		`json:"name"`
}

type History struct {
	ID        	uint      	`json:"id"`
	UserId		int 		`json:"user_id"`
	ProductId 	int 		`json:"product_id"`
	Quantity	int  		`json:"quantity"`
	TotalPrice 	int64 		`json:"total_price"`
	User 		User 		`json:"user"`
	Product 	Product		`json:"product"`
}