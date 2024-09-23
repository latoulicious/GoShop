package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/latoulicious/GoShop/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// Import each package based on the directory structure
	cartModel "github.com/latoulicious/GoShop/internal/cart"
	cartProductModel "github.com/latoulicious/GoShop/internal/cartproduct"
	categoryModel "github.com/latoulicious/GoShop/internal/category"
	orderModel "github.com/latoulicious/GoShop/internal/order"
	paymentModel "github.com/latoulicious/GoShop/internal/payment"
	productModel "github.com/latoulicious/GoShop/internal/product"
	userModel "github.com/latoulicious/GoShop/internal/user"
)

var DB *gorm.DB
var DSN string

// ConnectDB establishes a connection to the PostgreSQL database.
// It parses the DB connection settings from environment variables,
// constructs a connection string, opens a gorm DB connection,
// runs auto-migration on the model structs, and returns any errors.
func ConnectDB() error {
	var err error
	p := config.GetEnv("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Failed to parse database port")
		return err
	}

	if os.Getenv("ENVIRONMENT") == "development" {
		DSN = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.GetEnv("DB_HOST"), port, config.GetEnv("DB_USER"), config.GetEnv("DB_PASSWORD"), config.GetEnv("DB_NAME"))
		log.Println("Connecting to Database with DSN:", DSN)
	}

	DB, err = gorm.Open(postgres.Open(DSN))

	if err != nil {
		log.Println("Failed to connect to the database")
		return err
	}

	log.Println("Connection Opened to Database")

	// Auto Migrate Models
	err = DB.AutoMigrate(
		&userModel.User{},               // From internal/user/model.go
		&cartModel.ShoppingCart{},       // From internal/cart/model.go
		&cartProductModel.CartProduct{}, // From internal/cartproduct/model.go
		&categoryModel.Category{},       // From internal/category/model.go
		&orderModel.Order{},             // From internal/order/model.go
		&paymentModel.Payment{},         // From internal/payment/model.go
		&productModel.Product{},         // From internal/product/model.go
	)

	if err != nil {
		log.Println("Failed to auto migrate models:", err)
		return err
	}
	log.Println("Database Migrated")
	return nil
}
