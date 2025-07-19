# 🛍️ E-Commerce Microservices API in Golang

A modular and production-ready **eCommerce API** built using **Golang**, following **Hexagonal Architecture** and **Microservice Design**.  
Each service (User, Product, Order, Payment) is containerized using Docker, documented with Swagger, and fully tested via Postman.

---

## 📦 Microservices Included

| Service       | Port  | Description                   |
|---------------|-------|-------------------------------|
| `user-ms`     | 8081  | Handles user management       |
| `product-ms`  | 8082  | Manages products              |
| `order-ms`    | 8083  | Processes customer orders     |
| `payment-ms`  | 8084  | Handles payment transactions  |

---

## 🧱 Tech Stack

- **Language**: Go (Golang 1.21+)
- **Routing**: [Chi Router](https://github.com/go-chi/chi)
- **Database**: MongoDB
- **Architecture**: Hexagonal (Ports and Adapters)
- **Documentation**: Swagger (via swaggo)
- **Validation**: go-playground/validator
- **Containerization**: Docker, Docker Compose
- **Testing**: Postman Collection

---

## ⚙️ Getting Started

### 📁 Clone the repo

```bash
git clone https://github.com/jabez8650/ecommerce-ms.git
cd ecommerce-ms

```

##🗂️ Folder Structure

ecommerce-ms/
│
├── user-ms/
│   ├── internal/
│   ├── pkg/
│   ├── docs/
│   ├── main.go
│   ├── Dockerfile
│   └── go.mod
├── product-ms/
├── order-ms/
├── payment-ms/
│
├── docker-compose.yml
├── .env (used by all services)
├── ecommerce.postman_collection.json
└── README.md


## 🐳 Run All Services with Docker

```bash
docker-compose up --build

```


## 📚 API Documentation

Each microservice includes Swagger documentation.

Service	Swagger URL
User	http://localhost:8081/swagger/index.html
Product	http://localhost:8082/swagger/index.html
Order	http://localhost:8083/swagger/index.html
Payment	http://localhost:8084/swagger/index.html


## 🔁 Testing with Postman
A Postman collection is included to test all services.


## ✅ Features Covered
Create

Read All

Read by ID

Update

Delete

Edge Case Handling

Invalid Payloads


## 🔽 Import Postman Collection
File: ecommerce.postman_collection.json

Open Postman → Import → Choose file → Select collection

Start testing each endpoint


## ✅ Features Summary
 Modular Microservice Structure

 MongoDB Integration

 REST APIs with Chi Router

 Hexagonal Architecture

 Swagger Documentation per service

 Field Validation

 Docker & Docker Compose setup

 Postman Test Collection included


## 🧪 Example Endpoints
User Service

```GET    /api/users/
POST   /api/users/
GET    /api/users/{id}
PUT    /api/users/{id}
DELETE /api/users/{id}
Product / Order / Payment
```

Product / Order / Payment
```
GET    /api/{resource}/
POST   /api/{resource}/
GET    /api/{resource}/{id}
PUT    /api/{resource}/{id}
DELETE /api/{resource}/{id}
```

## 📬 Feedback
Suggestions, improvements, and issues are always welcome.
This project is a foundation — feel free to extend it!