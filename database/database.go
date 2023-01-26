package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"miniecommerce.wisnu.net/modules"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=MiniEcommerceGo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&modules.User{},
		&modules.Admin{},
		&modules.Seller{},
		&modules.Product{},
		&modules.Category{},
		&modules.History{},
		&modules.Balance{},
		&modules.Transaction{},
		
	)
	
	DB = db
}