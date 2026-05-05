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

The API exposes the following endpoints:

- POST /products (create)
- GET /products (list)
- GET /products/{id} (retrieve)
- PUT /products/{id} (update)
- DELETE /products/{id} (delete)

---

## 6. Persistence Test

To verify persistence:

```bash
docker compose down
docker compose up --build
```

If PostgreSQL is used correctly, data should persist after restart.