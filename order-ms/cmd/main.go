// @title           Order Microservice API
// @version         1.0
// @description     Order microservice using Go, MongoDB, Chi, Swagger
// @termsOfService  http://example.com/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  MIT License
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8083
// @BasePath  /api
// @schemes   http
package main

import (
	"log"
	"net/http"
	"os"

	orderhttp "order-ms/internal/order/adapter/http"
	"order-ms/internal/order/adapter/mongo"
	"order-ms/internal/order/usecase"
	"order-ms/pkg/config"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "order-ms/docs"
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found or error loading .env")
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found. Using system environment variables.")
	}

	db := config.ConnectMongo()
	orderCol := db.Database("orderdb").Collection("orders")

	repo := mongo.NewOrderRepository(orderCol)
	uc := usecase.NewOrderUseCase(repo)
	handler := orderhttp.NewOrderHandler(uc)

	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Route("/api", func(r chi.Router) {
		handler.RegisterRoutes(r)
	})

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set in the environment")
	}

	log.Printf("🚀 Server running at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
