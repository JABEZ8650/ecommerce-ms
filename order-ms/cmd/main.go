// @title           Order Microservice API
// @version         1.0
// @description     Order microservice using Go, MongoDB, Chi, Swagger
// @termsOfService  http://example.com/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  MIT License
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8081
// @BasePath  /api
// @schemes   http
package main

import (
	"log"
	"net/http"

	orderhttp "order-ms/internal/order/adapter/http"
	"order-ms/internal/order/adapter/mongo"
	"order-ms/internal/order/usecase"
	"order-ms/pkg/config"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "order-ms/docs"
)

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

	log.Println("Order service running at http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
