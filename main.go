package main

import (
	"fmt"
	"log"

	"github.com/DavidEsdrs/go-mercado/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupDatabase() error {
	dsn := "root:mysqlPW@tcp(db:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		&models.Product{},
		&models.Cart{},
		&models.ProductCart{},
		&models.Vendor{},
	)

	if err != nil {
		return err
	}

	fmt.Println("Auto-migration completed successfully.")
	return nil
}

func main() {
	fmt.Println("starting...")
	err := setupDatabase()
	if err != nil {
		log.Fatal("error while starting database: ", err.Error())
	}

	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, "{success: true}")
	})

	r.Run(":8080")
}
