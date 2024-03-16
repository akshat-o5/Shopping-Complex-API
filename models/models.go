package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Shop struct {
	gorm.Model
	ShopNumber int    `json:"shopNumber" gorm:"primary key"`
	Name       string `json:"name"`
	Owner      string `jason:"owner"`
	Phone      int64  `json:"phone"`
	Rent       int64  `json:"rent"`
	Status     bool   `json:"status" gorm:"default:false"`
}

func ConnectDB() *gorm.DB {
	dsn := "root:klaus@17320@tcp(127.0.0.1:3306)/shoppingComplex?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error is : ", err)
		panic(err)
	}
	database.AutoMigrate(&Shop{})
	fmt.Println("Database Connection Established")
	db = database
	return db
}
