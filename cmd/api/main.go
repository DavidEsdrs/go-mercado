package main

import (
	"os"

	"github.com/DavidEsdrs/go-mercado/internal/config"
	"github.com/DavidEsdrs/go-mercado/internal/handler"
	"github.com/DavidEsdrs/go-mercado/internal/middleware"
	"github.com/DavidEsdrs/go-mercado/internal/model"
	"github.com/DavidEsdrs/go-mercado/internal/repository"
	service "github.com/DavidEsdrs/go-mercado/internal/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	gorm_logger "gorm.io/gorm/logger"
)

func main() {
	logger := setupAppLogger("app.log")
	defer logger.Sync()

	db, err := setupDatabase(logger)
	if err != nil {
		logger.Fatal("error while starting database: " + err.Error())
	}

	productHandler := CreateProductHandler(db)
	userHandler := CreateUserHandler(db, logger)

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.TimeLogging(logger))

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

	r.Run(":8080")
}

func setupDatabase(log *zap.Logger) (*gorm.DB, error) {
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

func CreateUserHandler(db *gorm.DB, logger *zap.Logger) *handler.UserHandler {
	repoService := repository.NewUserRepository(db)
	userService := service.NewUserService(repoService)
	return handler.NewUserHandler(userService, logger)
}

func setupAppLogger(logFileName string) *zap.Logger {
	userLogs, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic("unable to create users log file! error: " + err.Error())
	}

	writeSyncer := zapcore.AddSync(userLogs)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)

	core := zapcore.NewCore(fileEncoder, writeSyncer, zapcore.InfoLevel)

	return zap.New(core)
}
