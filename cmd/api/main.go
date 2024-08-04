package main

import (
	"fmt"
	"log"

	"github.com/DavidEsdrs/go-mercado/internal/config"
	"github.com/DavidEsdrs/go-mercado/internal/handler"
	"github.com/DavidEsdrs/go-mercado/internal/model"
	"github.com/DavidEsdrs/go-mercado/internal/repository"
	service "github.com/DavidEsdrs/go-mercado/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupDatabase() (*gorm.DB, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&model.Product{},
		&model.Cart{},
		&model.ProductCart{},
		&model.Vendor{},
	)

	if err != nil {
		return nil, err
	}

	fmt.Println("Auto-migration completed successfully.")
	return db, nil
}

func main() {
	db, err := setupDatabase()
	if err != nil {
		log.Fatal("error while starting database: ", err.Error())
	}

	repoService := repository.NewProductRepository(db)

	productService := service.NewProductService(repoService)

	productHandler := handler.NewProductHandler(productService)

	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]interface{}{
			"success": true,
		})
	})

	r.POST("/product", productHandler.CreateProduct)

	r.Run(":8080")
}
