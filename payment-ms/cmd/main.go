// @title Payment Microservice API
// @version 1.0
// @description This is the Payment microservice.
// @host localhost:8084
// @BasePath /api

package main

import (
	"log"
	"net/http"
	paymenthttp "payment-ms/internal/payment/adapter/http"

	"payment-ms/internal/payment/adapter/mongo"
	"payment-ms/internal/payment/usecase"
	"payment-ms/pkg/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "payment-ms/docs"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(" .env file not found. Using system environment variables.")
	}

	db := config.ConnectMongo()
	paymentCol := db.Database("paymentdb").Collection("payments")

	repo := mongo.NewPaymentRepository(paymentCol)
	uc := usecase.NewPaymentUseCase(repo)
	handler := paymenthttp.NewPaymentHandler(uc)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/api", func(r chi.Router) {
		handler.RegisterRoutes(r)
	})

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	log.Println(" Payment service running at http://localhost:8084")
	log.Fatal(http.ListenAndServe(":8084", r))
}
