// @title User Microservice API
// @version 1.0
// @description This is a user microservice.
// @host localhost:8082
// @BasePath /api

package main

import (
	"log"
	"net/http"

	userhttp "user-ms/internal/user/adapter/http"
	"user-ms/internal/user/adapter/mongo"
	"user-ms/internal/user/usecase"
	"user-ms/pkg/config"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	_ "user-ms/docs" // ✅ Swagger docs

	httpSwagger "github.com/swaggo/http-swagger"
)

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

	log.Println("🚀 User service running at http://localhost:8082")
	log.Println("📄 Swagger UI available at http://localhost:8082/swagger/index.html")
	log.Fatal(http.ListenAndServe(":8082", r))
}
