# 🛍️ Ecommerce Order System

A complete order processing microservice built with Go, featuring event-driven architecture with RabbitMQ and MongoDB persistence.

## ✨ Features

- **RESTful API** for order management
- **Event-driven architecture** using RabbitMQ
- **MongoDB** data persistence
- **Docker** with Docker Compose
- **Comprehensive testing** with mocking
- **Error handling** and validation
- **UUID generation** for orders
- **Order status tracking**

## 🏗️ System Workflow

HTTP Client → Gin Router → Order Handler → PlaceOrder UseCase → RabbitMQ Publisher → MongoDB Repository

## 🛠️ Tech Stack

- **Go 1.21+** - Backend programming language
- **Gin** - HTTP web framework
- **RabbitMQ** - Message broker for events
- **MongoDB** - NoSQL database
- **Docker** - Containerization
- **Testify** - Testing framework
- **UUID** - Unique identifier generation

### 🐳 Docker Deployment

- Docker
- Docker Compose
- Git

## 🚀 Quick Start
```
# Clone the repository
git clone https://github.com/your-username/ecommerce-order-system.git
cd ecommerce-order-system

# Start all services with Docker Compose
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

## Docker Commands Explained
```
# Start services in detached mode (background)
docker-compose up -d

# Build and start services (force rebuild)
docker-compose up --build

# View running containers
docker-compose ps

# View logs
docker-compose logs

# Stop and remove containers
docker-compose down

# Restart specific service
docker-compose restart api
```


### 📡 API Documentation

#### POST /orders

##### Creates a new order

Request Body
```
    {
  "client_name": "John Doe",
  "client_email": "john@example.com",
  "shipping_value": 15.9,
  "address": {
    "cep": 12345678,
    "street": "Main Street"
  },
  "payment_method": "CREDIT",
  "items": [
    {
      "item_id": 1,
      "item_description": "Premium T-Shirt",
      "item_value": 59.9,
      "item_quantity": 2,
      "discount": 10
    }
  ]
}
```

Success Response:

```
{
  "message": "Order processed successfully",
  "status": "success",
  "data": {
    "order_id": "f64294ba-4c40-4ba2-a375-c3739750013c",
    "order_date": "2025-08-25T15:29:34.430708Z",
    "order_status": "OPEN",
    "client_name": "John Doe",
    "client_email": "john@example.com",
    "shipping_value": 15.9,
    "address": {
      "cep": 12345678,
      "street": "Main Street"
    },
    "payment_method": "CREDIT",
    "items": [
      {
        "item_id": 1,
        "item_description": "Premium T-Shirt",
        "item_value": 59.9,
        "item_quantity": 2,
        "discount": 10,
        "total_value": 109.8
      }
    ],
    "total_value": 125.7
  }
}
```

### 🧪 Testing

 Run the test suite:

```
# Run all tests
go test ./...

# Run tests with coverage
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run specific package tests
go test ./internal/application/usecase -v
```

