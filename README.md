# eCommerce Microservices (Go + MongoDB)

This project is a microservices-based eCommerce backend built in Go using Chi and MongoDB. It includes:

- 🛍️ Product Service
- 📦 Order Service
- 👤 User Service
- 💳 Payment Service

### Features
- ✅ CRUD Operations
- ✅ Swagger API Documentation
- ✅ Validation
- ✅ Hexagonal Architecture
- ✅ MongoDB Integration

### Technologies
- Go (Chi Router)
- MongoDB
- Swaggo (Swagger UI)
- go-playground/validator

> Each service is built independently and can run on its own.

### How to Run
1. Open the folder for any service (e.g., `payment-ms`)
2. Run:
   ```bash
   go run main.go