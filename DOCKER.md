# Docker & Docker Compose Analysis

## 1. Multi-Stage Docker Build

The project uses a multi-stage Docker build.

### Build Stage
- Uses a Go base image
- Copies source code into the container
- Downloads dependencies
- Compiles the Go application into a binary

### Runtime Stage
- Uses a minimal base image
- Copies only the compiled binary
- Does not include source code or build tools

Result: smaller and more secure production image

---

## 2. CGO_ENABLED=0

`CGO_ENABLED=0` disables CGO (C-Go integration).

### Why this is important:
- Produces a fully static binary
- No dependency on system libraries (e.g. glibc)
- Enables usage of minimal container images like alpine or scratch
- Improves portability across environments

---

## 3. Image Size Comparison

### Multi-stage build
- Final image contains only the compiled binary
- Small image size (typically ~10–20MB depending on base image)

### Single-stage build
- Includes compiler, source code, and dependencies
- Much larger image (often hundreds of MB)

Conclusion: multi-stage builds significantly reduce image size and improve security

---

## 4. Docker Compose

Docker Compose is used to run the full application stack locally.

Typically includes:
- Go API service
- PostgreSQL database

Benefits:
- One command startup: docker compose up --build
- Reproducible development environment
- Easy service orchestration

---

## 5. CRUD Testing

The API was tested using curl commands.

### Health Check
```bash
curl http://localhost:8080/health
```

Response:
```json
{"status":"ok"}
```

### Create Products
```bash
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{"name":"Laptop","price":999.99}'
```

Response:
```json
{"id":1,"name":"Laptop","price":999.99}
```

```bash
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{"name":"Maus","price":29.99}'
```

Response:
```json
{"id":2,"name":"Maus","price":29.99}
```

```bash
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{"name":"Tastatur","price":59.99}'
```

Response:
```json
{"id":3,"name":"Tastatur","price":59.99}
```

### Read Products

```bash
curl http://localhost:8080/products
```

Response:
```json
[
  {"id":1,"name":"Laptop","price":999.99},
  {"id":2,"name":"Maus","price":29.99},
  {"id":3,"name":"Tastatur","price":59.99}
]
```

### Update Products

```bash
curl -X PUT http://localhost:8080/products/1 \
-H "Content-Type: application/json" \
-d '{"name":"Laptop Pro","price":1299.99}'
```

Response:
```json
{"id":1,"name":"Laptop Pro","price":1299.99}
```

### Delete Products

```bash
curl -X DELETE http://localhost:8080/products/3
```

Response:
```json
{"result":"success"}
```

### Final State

```bash
curl http://localhost:8080/products
```

Response:
```json
[
  {"id":1,"name":"Laptop Pro","price":1299.99},
  {"id":2,"name":"Maus","price":29.99}
]
```