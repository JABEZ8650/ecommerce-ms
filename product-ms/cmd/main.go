// @title           Product Microservice API
// @version         1.0
// @description     This is a simple product microservice API using Go, Chi, MongoDB, and Swagger.
// @termsOfService  http://example.com/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  MIT License
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api

// @schemes http
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	// This line is crucial for Swagger to work. It initializes the generated docs.
	_ "product-ms/docs" // Import the generated docs package. Note the underscore alias.

	httpSwagger "github.com/swaggo/http-swagger"

	producthttp "product-ms/internal/product/adapter/http"
	"product-ms/internal/product/adapter/mongo"
	"product-ms/internal/product/usecase"
	"product-ms/pkg/config"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found. Using system environment variables.")
	}

	// Mongo connection
	db := config.ConnectMongo()
	productCollection := db.Database("productdb").Collection("products")

	// Dependency injection
	repo := mongo.NewProductRepository(productCollection)
	uc := usecase.NewProductUseCase(repo)
	handler := producthttp.NewProductHandler(uc)

	// Setup router
	r := chi.NewRouter()

	// Register Swagger UI endpoint
	// This makes the Swagger documentation accessible at http://localhost:8080/swagger/index.html
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	// Group your API routes under /api prefix
	r.Route("/api", func(r chi.Router) {
		// This line ensures all routes registered by handler.RegisterRoutes
		// will be prefixed with /api, e.g., /api/products
		handler.RegisterRoutes(r)
	})

	log.Println(" Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
