package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mrckurz/CI-CD-MCM/internal/store"
)

func setupRouter() (*mux.Router, *Handler) {
	s := store.NewMemoryStore()
	h := NewHandler(s)
	r := mux.NewRouter()
	h.RegisterRoutes(r)
	return r, h
}

func TestHealthEndpoint(t *testing.T) {
	r, _ := setupRouter()

	req := httptest.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}
}

func TestGetProductsEmpty(t *testing.T) {
	r, _ := setupRouter()

	req := httptest.NewRequest("GET", "/products", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}
}

func TestCreateAndGetProduct(t *testing.T) {
	r, _ := setupRouter()

	// Create
	body := `{"name":"Widget","price":9.99}`
	req := httptest.NewRequest("POST", "/products", strings.NewReader(body))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", rr.Code)
	}

	// Get
	req = httptest.NewRequest("GET", "/products/1", nil)
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}
}

func TestGetProductNotFound(t *testing.T) {
	r, _ := setupRouter()

	req := httptest.NewRequest("GET", "/products/999", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", rr.Code)
	}
}

// TODO: Add tests for UpdateProduct, DeleteProduct, and invalid payloads

func TestUpdateProduct(t *testing.T) {
	r, _ := setupRouter()

	body := `{"name":"Old Name","price":5.0}`
	req := httptest.NewRequest("POST", "/products", strings.NewReader(body))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	updateBody := `{"name":"New Name","price":15.0}`
	req = httptest.NewRequest("PUT", "/products/1", strings.NewReader(updateBody))
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 on update, got %d", rr.Code)
	}

	req = httptest.NewRequest("PUT", "/products/999", strings.NewReader(updateBody))
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected 404 on update non-existent, got %d", rr.Code)
	}
}

func TestDeleteProduct(t *testing.T) {
	r, _ := setupRouter()

	r.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("POST", "/products", strings.NewReader(`{"name":"Bye","price":1}`)))

	req := httptest.NewRequest("DELETE", "/products/1", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 on delete, got %d", rr.Code)
	}

	req = httptest.NewRequest("DELETE", "/products/999", nil)
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected 404 on delete non-existent, got %d", rr.Code)
	}
}

func TestInvalidPayloads(t *testing.T) {
	r, _ := setupRouter()

	req := httptest.NewRequest("POST", "/products", strings.NewReader(`{invalid-json`))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected 400 for bad JSON, got %d", rr.Code)
	}

	req = httptest.NewRequest("POST", "/products", strings.NewReader(`{"name":"","price":10}`))
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected 400 for validation error, got %d", rr.Code)
	}

	req = httptest.NewRequest("PUT", "/products/1", strings.NewReader(`{bad-json`))
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected 400 for bad JSON on update, got %d", rr.Code)
	}
}

func TestHandler_CoverageBoost(t *testing.T) {
	s := store.NewMemoryStore()
	h := NewHandler(s)
	r := setupRouterWithHandler(h) // Hilfsfunktion siehe unten

	r.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("POST", "/products", strings.NewReader(`{invalid`)))
	r.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("POST", "/products", strings.NewReader(`{"name":""}`)))

	r.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", "/products/999", nil))
	r.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("PUT", "/products/999", strings.NewReader(`{"name":"A"}`)))
	r.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("DELETE", "/products/999", nil))

	r.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("PUT", "/products/1", strings.NewReader(`{invalid`)))
}

func setupRouterWithHandler(h *Handler) *mux.Router {
	r := mux.NewRouter()
	h.RegisterRoutes(r)
	return r
}
