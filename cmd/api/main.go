package main

import (
	"os"

	"github.com/DavidEsdrs/go-mercado/internal/config"
	"github.com/DavidEsdrs/go-mercado/internal/handler"
	"github.com/DavidEsdrs/go-mercado/internal/middleware"
	"github.com/DavidEsdrs/go-mercado/internal/model"
	"github.com/DavidEsdrs/go-mercado/internal/repository"
	service "github.com/DavidEsdrs/go-mercado/internal/services"
	"github.com/DavidEsdrs/go-mercado/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	gorm_logger "gorm.io/gorm/logger"
)

var log *logger.Logger

func main() {
	log = logger.New(os.Stdout, "APP", logger.LstdFlags|logger.Ltime)
	log.SetLevel(logger.INFO)

	db, err := setupDatabase(log)
	if err != nil {
		log.Fatal("error while starting database: %v", err.Error())
	}

	productHandler := CreateProductHandler(db)
	userHandler := CreateUserHandler(db)

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.TimeLogging(log))

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": true,
		})
	})

	r.POST("/login", userHandler.Login)
	r.POST("/signup", userHandler.Signup)

	r.POST("/product", productHandler.CreateProduct)
	r.GET("/product/:id", productHandler.ReadProduct)
	r.GET("/product", productHandler.ReadProducts)

	log.Fatal("%v", r.Run(":8080"))
}

func setupDatabase(log *logger.Logger) (*gorm.DB, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	silentLogger := gorm_logger.New(nil, gorm_logger.Config{})

	db, err := gorm.Open(mysql.Open(cfg.DatabaseURL), &gorm.Config{
		Logger: silentLogger,
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Cart{},
		&model.ProductCart{},
		&model.Vendor{},
	)

	if err != nil {
		return nil, err
	}

	log.Info("Auto-migration completed successfully.")
	return db, nil
}

func CreateProductHandler(db *gorm.DB) *handler.ProductHandler {
	repoService := repository.NewProductRepository(db)
	productService := service.NewProductService(repoService)
	return handler.NewProductHandler(productService)
}

func CreateUserHandler(db *gorm.DB) *handler.UserHandler {
	var userLogger *logger.Logger
	userLogs, err := os.OpenFile("user-logs.txt", os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic("unable to create users log file!")
	}
	userLogger = logger.New(userLogs, "USER", logger.LstdFlags|logger.Ltime)

	repoService := repository.NewUserRepository(db)
	userService := service.NewUserService(repoService)
	return handler.NewUserHandler(userService, userLogger)
}
