# Event Management REST API

A robust, scalable RESTful API built with Go designed to handle event management operations. This application provides secure user authentication, event creation, and ticket booking functionalities, paired with an automated Swagger UI for seamless API testing and documentation.

## 🚀 Tech Stack

* **Language:** Go (Golang)
* **Framework:** Gin (`gin-gonic/gin`) for high-performance HTTP routing
* **Database:** SQLite (lightweight relational data persistence)
* **Authentication:** JSON Web Tokens (JWT) for stateless, secure session management
* **Documentation:** Swagger / OpenAPI (`swaggo/gin-swagger`)

## ✨ Key Features

* **User Management:** Secure user registration and login with encrypted credentials.
* **Authentication Middleware:** Custom middleware to validate JWTs and protect private routes.
* **Event Operations (CRUD):** Authenticated users can create, update, delete, and fetch event details.
* **Ticket Booking:** Concurrency-safe ticket registration tied to specific event IDs.
* **Interactive UI:** Auto-generated Swagger documentation for testing endpoints directly in the browser.

## 🛠️ Getting Started

### Prerequisites
* [Go](https://go.dev/doc/install) (1.18 or higher)
* Git

1. **Clone the repository:**
   ```bash
   git clone https://github.com/saraghgh99/go-rest-api-starter-.git
   cd go-rest-api-starter-
   ```

2. **Install dependencies:**
```bash
go mod tidy
```

3. **Generate the Swagger Documentation:**
```bash
swag init
```

4. **Run the server:**
```bash
go run main.go
```

📖 API Documentation
This project uses Swagger for interactive API documentation. Once the server is running, you can access the UI in your browser at:

http://localhost:8080/swagger/index.html

From here, you can view all available endpoints, required payloads, and test the API responses directly.




