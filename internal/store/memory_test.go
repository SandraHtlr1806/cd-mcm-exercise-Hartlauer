package store

import (
	"testing"

	"github.com/mrckurz/CI-CD-MCM/internal/model"
)

func TestMemoryStore_CoverageBoost(t *testing.T) {
	s := NewMemoryStore()

	p := s.Create(model.Product{Name: "Test", Price: 10})
	if len(s.GetAll()) != 1 {
		t.Error("Expected 1")
	}

	_, err := s.GetByID(p.ID)
	if err != nil {
		t.Error("Should find it")
	}
	_, err = s.GetByID(999)
	if err != ErrNotFound {
		t.Error("Should be ErrNotFound")
	}

	_, err = s.Update(p.ID, model.Product{Name: "New"})
	if err != nil {
		t.Error("Update should work")
	}
	_, err = s.Update(999, model.Product{})
	if err != ErrNotFound {
		t.Error("Update fail expected")
	}

	err = s.Delete(p.ID)
	if err != nil {
		t.Error("Delete should work")
	}
	err = s.Delete(p.ID)
	if err != ErrNotFound {
		t.Error("Expected ErrNotFound")
	}
}

func TestMemoryStore_AllPathCoverage(t *testing.T) {
	s := NewMemoryStore()

	p := s.Create(model.Product{Name: "Test", Price: 10})
	if len(s.GetAll()) != 1 {
		t.Error("Should have 1 product")
	}

	if _, err := s.GetByID(p.ID); err != nil {
		t.Error("Product should exist")
	}
	if _, err := s.GetByID(999); err != ErrNotFound {
		t.Error("Should return ErrNotFound")
	}

	if _, err := s.Update(p.ID, model.Product{Name: "New"}); err != nil {
		t.Error("Update should work")
	}
	if _, err := s.Update(999, model.Product{}); err != ErrNotFound {
		t.Error("Update should fail")
	}

	if err := s.Delete(p.ID); err != nil {
		t.Error("Delete should work")
	}
	if err := s.Delete(p.ID); err != ErrNotFound {
		t.Error("Delete should fail now")
	}
}
