# Architecture Documentation

## 1. Overview

The Product Catalog API is a RESTful microservice built in Go using Gorilla Mux.  
It follows a layered architecture consisting of Router, Handler, Model, and Store layers.

The system supports two interchangeable storage implementations:
- In-memory storage (MemoryStore)
- PostgreSQL storage (PostgresStore)

The active implementation is selected at runtime via environment variables.

## 2. Request Flow

The application processes HTTP requests through the following pipeline:

Client → HTTP Server → Router → Handler → Model → Store → Memory / PostgreSQL

### Detailed Flow

1. The application starts an HTTP server using `http.ListenAndServe`.
2. Incoming requests are received by the Gorilla Mux router.
3. The router matches the request to a registered endpoint.
4. The corresponding handler function is executed.
5. The handler:
   - parses JSON request data
   - validates input using the model layer
   - calls the store layer for data operations
6. The store performs CRUD operations either in memory or in PostgreSQL.
7. The response is returned as JSON to the client.

---

## 3. Architecture Diagram

        Client
          ↓
  HTTP Server (main.go)
          ↓
  Router (Gorilla Mux)
          ↓
     Handler Layer
          ↓
Model Layer (validation)
          ↓
      Store Layer
          ↓
+------------------------+
| MemoryStore            |
| PostgresStore          |
+------------------------+
          ↓
  Memory / PostgreSQL

---

## 4. Model Layer

The model layer defines the core domain object of the application: `Product`.

```go
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
```

---

### Validation

The model includes business validation logic to ensure data integrity before data is processed further:

- The product name must not be empty
- The product price must not be negative

```go
func (p *Product) Validate() bool {
	if p.Name == "" {
		return false
	}
	if p.Price < 0 {
		return false
	}
	return true
}
```

This ensures that only valid data reaches the handler and storage layers.

## 5. Handler Layer

The handler layer is responsible for processing HTTP requests and generating responses.

Each handler function:
- parses incoming JSON data
- validates input using the model layer
- interacts with the store layer
- returns JSON responses with appropriate HTTP status codes

Two implementations exist:
- Handler → uses MemoryStore
- PostgresHandler → uses PostgresStore

Example Flow (Create Product)
1. Decode JSON request body into a Product struct
2. Validate product using Validate()
3. Call Store.Create(p)
4. Return HTTP 201 Created with the created product

This layer represents the business logic and HTTP interface of the application.

## 6. Store Layer

The store layer abstracts data persistence and provides a consistent interface for data operations.

### MemoryStore

The MemoryStore stores data in an in-memory map and is protected using a mutex for thread safety.

Characteristics:
- Fast execution (no database calls)
- Simple implementation
- Data is lost after application restart
- Suitable for testing and development

#### When to use MemoryStore

MemoryStore is used during development, testing, or when persistence is not required. It is ideal for fast prototyping and unit tests because it does not require any external dependencies.

### PostgresStore

The PostgresStore persists data in a PostgreSQL database using SQL queries.

Characteristics:
- Persistent storage
- Suitable for production systems
- Requires database setup and connection
- More complex due to SQL and error handling

#### When to use PostgresStore

PostgresStore is used in production environments where data persistence, reliability, and scalability are required.

### Trade-offs

The main trade-off between MemoryStore and PostgresStore is simplicity versus persistence.

- MemoryStore is lightweight and fast but does not persist data
- PostgresStore provides durability and scalability but adds complexity and overhead

## 7. Overall Design Principles

The application follows several software engineering principles:

### Separation of Concerns

Each layer has a distinct responsibility:

- Router → routing logic
- Handler → HTTP + business logic
- Model → validation rules
- Store → data persistence

### Dependency Injection

The handler receives its store implementation as a dependency, allowing flexibility between MemoryStore and PostgresStore.

### Flexibility

The storage backend can be switched at runtime using environment variables without changing application logic.

## 8. Conclusion

This architecture demonstrates a clean layered design that improves:

- maintainability
- testability
- scalability
- flexibility

It separates HTTP handling, business logic, validation, and persistence into independent layers, making the system easy to extend and adapt.
