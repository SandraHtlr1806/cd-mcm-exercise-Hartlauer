package store

// import "testing"
import (
	"testing"

	"github.com/mrckurz/CI-CD-MCM/internal/model"
)

func TestCreateAndGet(t *testing.T) {
	s := NewMemoryStore()
	// TODO: Add test -- create a product and verify GetByID returns it

	product := model.Product{Name: "Test Product", Price: 9.99}
	created := s.Create(product)
	if created.ID == 0 {
		t.Error("expected non-zero ID for created product")
	}

	retrieved, err := s.GetByID(created.ID)
	if err != nil {
		t.Errorf("unexpected error retrieving product: %v", err)
	}

	if retrieved.Name != product.Name || retrieved.Price != product.Price {
		t.Error("retrieved product does not match created product")
	}
}

func TestGetAllEmpty(t *testing.T) {
	s := NewMemoryStore()
	products := s.GetAll()
	if len(products) != 0 {
		t.Errorf("expected 0 products, got %d", len(products))
	}
}

func TestDeleteNonExistent(t *testing.T) {
	s := NewMemoryStore()
	err := s.Delete(999)
	if err != ErrNotFound {
		t.Error("expected ErrNotFound when deleting non-existent product")
	}
}

// TODO: Add tests for Update, Delete of existing product, and GetByID with invalid ID
func TestUpdateProduct(t *testing.T) {
	s := NewMemoryStore()
	// TODO: Add test -- Create, update, verify update was applied

	product := model.Product{Name: "Test Product", Price: 9.99}
	created := s.Create(product)

	created.Price = 19.99

	updated, err := s.Update(created.ID, created)
	if err != nil {
		t.Errorf("unexpected error updating product: %v", err)
	}

	if updated.Price != created.Price {
		t.Error("product price was not updated")
	}

	if updated.ID != created.ID {
		t.Errorf("expected ID %d, got %d", created.ID, updated.ID)
	}

	if updated.Name != created.Name || updated.Price != 19.99 {
		t.Errorf("update failed, got %+v", updated)
	}

	got, _ := s.GetByID(created.ID)
	if got.Name != created.Name || got.Price != 19.99 {
		t.Errorf("store not updated, got %+v", got)
	}
}

func TestDeleteProduct(t *testing.T) {
	s := NewMemoryStore()
	// TODO: Add test -- Create, delete, verify GetByID returns error

	product := model.Product{Name: "Test Product", Price: 9.99}
	created := s.Create(product)

	err := s.Delete(created.ID)
	if err != nil {
		t.Errorf("unexpected error deleting product: %v", err)
	}

	_, err = s.GetByID(created.ID)
	if err != ErrNotFound {
		t.Errorf("expected ErrNotFound after delete, got %v", err)
	}
}

func TestGetByIDNotFound(t *testing.T) {
	s := NewMemoryStore()
	// TODO: Add test -- Verify non-existent ID returns ErrNotFound

	tests := []struct {
		name string
		id   int
	}{
		{"zero id", 0},
		{"negative id", -1},
		{"non-existent id", 999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetByID(tt.id)
			if err != ErrNotFound {
				t.Errorf("expected ErrNotFound, got %v", err)
			}
		})
	}
}
