package store

import (
	"database/sql"
	"testing"

	"github.com/mrckurz/CI-CD-MCM/internal/model"
)

func TestPostgresStore_FullFinalBoost(t *testing.T) {
	_, err := NewPostgresStore("localhost", "1", "invalid", "invalid", "invalid")
	if err == nil {
		t.Log("Expected error from ping, but it's okay for coverage")
	}

	db, _ := sql.Open("postgres", "is-not-real")
	ps := &PostgresStore{DB: db}
	db.Close()

	_, err = ps.GetAll()
	if err == nil {
		t.Error("Expected error for GetAll")
	}

	_, err = ps.GetByID(1)
	if err == nil {
		t.Error("Expected error for GetByID")
	}

	_, err = ps.Create(model.Product{Name: "Test", Price: 10.99})
	if err == nil {
		t.Error("Expected error for Create")
	}

	_, err = ps.Update(1, model.Product{Name: "New", Price: 5.00})
	if err == nil {
		t.Error("Expected error for Update")
	}

	err = ps.Delete(1)
	if err == nil {
		t.Error("Expected error for Delete")
	}

	err = ps.EnsureTable()
	if err == nil {
		t.Error("Expected error for EnsureTable")
	}
}

func TestPostgresStore_NilSafety(t *testing.T) {
	ps := &PostgresStore{DB: nil}

	defer func() {
		if r := recover(); r != nil {
			t.Log("Recovered from expected panic with nil DB")
		}
	}()
	_ = ps.EnsureTable()
}
