// @title User Microservice API
// @version 1.0
// @description This is a user microservice.
// @host localhost:8081
// @BasePath /api

package main

import (
	"log"
	"net/http"
	"os"

	userhttp "user-ms/internal/user/adapter/http"
	"user-ms/internal/user/adapter/mongo"
	"user-ms/internal/user/usecase"
	"user-ms/pkg/config"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	_ "user-ms/docs"

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
	userCol := db.Database("userdb").Collection("users")

	repo := mongo.NewUserRepository(userCol)
	uc := usecase.NewUserUseCase(repo)
	handler := userhttp.NewUserHandler(uc)

	r := chi.NewRouter()

	// ✅ Swagger route
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	// ✅ API route group
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
