// @title Payment Microservice API
// @version 1.0
// @description This is the Payment microservice.
// @host localhost:8084
// @BasePath /api

package main

import (
	"log"
	"net/http"
	"os"

	_ "payment-ms/docs"

	paymenthttp "payment-ms/internal/payment/adapter/http"
	"payment-ms/internal/payment/adapter/mongo"
	"payment-ms/internal/payment/usecase"
	"payment-ms/pkg/config"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
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
	col := db.Database("paymentdb").Collection("payments")

	repo := mongo.NewPaymentRepository(col)
	uc := usecase.NewPaymentUseCase(repo)
	handler := paymenthttp.NewPaymentHandler(uc)

	r := chi.NewRouter()

	// Register Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	// Register Payment Routes
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
