📦 Go Backend API with Gin, GORM, PostgreSQL

A RESTful API backend built with Go (Gin-Gonic), PostgreSQL, and GORM. Includes JWT authentication, Docker support, and clean MVC architecture.

📁 Project Structure
.
├── config/           # DB configuration

├── controllers/      # Route logic

├── middleware/       # JWT auth

├── models/           # DB schemas

├── routes/           # Router groups

├── utils/            # Helper functions (JWT etc.)

├── Dockerfile

├── docker-compose.yml

├── go.mod / go.sum

└── main.go

🚀 Getting Started

🐳 Run with Docker

docker-compose up --build

API: http://localhost:8080/api/v1

PostgreSQL: localhost:5433

🔐 Authentication Flow

POST /api/v1/auth/register – Register a new user

POST /api/v1/auth/login – Returns JWT token

Use Authorization: Bearer <token> header for secured routes

📦 API Endpoints

Auth

POST /api/v1/auth/register

POST /api/v1/auth/login

Customers

GET /api/v1/customers

GET /api/v1/customers/:id

POST /api/v1/customers

PUT /api/v1/customers/:id

DELETE /api/v1/customers/:id

Orders

GET /api/v1/orders

GET /api/v1/orders/:id

POST /api/v1/orders

PUT /api/v1/orders/:id

DELETE /api/v1/orders/:id
