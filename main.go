package main

import (
	"fmt"
	"log"

	"github.com/DavidEsdrs/go-mercado/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:mysqlPW@tcp(db:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(
		&models.Product{},
		&models.Cart{},
		&models.ProductCart{},
		&models.Vendor{},
	)

	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	fmt.Println("Auto-migration completed successfully.")
}
