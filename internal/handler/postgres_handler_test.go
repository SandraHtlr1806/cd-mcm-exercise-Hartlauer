package handler

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mrckurz/CI-CD-MCM/internal/store"
)

func TestPostgresHandler_FinalBoost(t *testing.T) {
	db, _ := sql.Open("postgres", "is-not-real")
	db.Close()
	s := &store.PostgresStore{DB: db}
	h := NewPostgresHandler(s)

	run := func(method, url, body string, vars map[string]string, fn http.HandlerFunc) {
		req := httptest.NewRequest(method, url, strings.NewReader(body))
		if vars != nil {
			req = mux.SetURLVars(req, vars)
		}
		fn(httptest.NewRecorder(), req)
	}

	h.RegisterRoutes(mux.NewRouter()) // Deckt die Registrierung ab

	run("GET", "/health", "", nil, h.Health)
	run("GET", "/products", "", nil, h.GetProducts)
	run("GET", "/products/1", "", map[string]string{"id": "1"}, h.GetProduct)
	run("DELETE", "/products/1", "", map[string]string{"id": "1"}, h.DeleteProduct)

	run("POST", "/products", `{"name":"T","price":10}`, nil, h.CreateProduct)
	run("POST", "/products", `{invalid}`, nil, h.CreateProduct)              // Error-Zweig 1
	run("POST", "/products", `{"name":"","price":-1}`, nil, h.CreateProduct) // Error-Zweig 2

	run("PUT", "/products/1", `{"name":"T","price":10}`, map[string]string{"id": "1"}, h.UpdateProduct)
	run("PUT", "/products/1", `{invalid}`, map[string]string{"id": "1"}, h.UpdateProduct)
}
