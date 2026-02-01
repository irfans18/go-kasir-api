package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/database"
	"kasir-api/handler"
	"kasir-api/repository"
	"kasir-api/service"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	// Server Config
	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// Setup database
	db, err := database.InitDB(config.DBConn)

	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// GET localhost:8080/api/v1/products/{id}
	// PUT localhost:8080/api/v1/products/{id}
	// DELETE localhost:8080/api/v1/products/{id}
	http.HandleFunc("/api/v1/products", productHandler.HandleProducts)

	// GET localhost:8080/api/v1/products
	// POST localhost:8080/api/v1/products
	http.HandleFunc("/api/v1/products/", productHandler.HandleProductByID)

	// GET localhost:8080/api/v1/categories/{id}
	// PUT localhost:8080/api/v1/categories/{id}
	// DELETE localhost:8080/api/v1/categories/{id}
	http.HandleFunc("/api/v1/categories/", categoryHandler.HandleCategoryByID)

	// GET localhost:8080/api/v1/categories
	// POST localhost:8080/api/v1/categories
	http.HandleFunc("/api/v1/categories", categoryHandler.HandleCategorys)

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("gagal running server", err)
	}
}
