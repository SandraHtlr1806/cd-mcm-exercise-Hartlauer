package model

import "testing"

func TestValidateEmptyName(t *testing.T) {
	p := Product{Name: "", Price: 10.0}
	if p.Validate() {
		t.Error("expected validation to fail for empty name")
	}
}

func TestValidateNegativePrice(t *testing.T) {
	p := Product{Name: "Widget", Price: -5.0}
	if p.Validate() {
		t.Error("expected validation to fail for negative price")
	}
}

func TestValidateValidProduct(t *testing.T) {
	p := Product{Name: "Widget", Price: 9.99}
	if !p.Validate() {
		t.Error("expected validation to pass for valid product")
	}
}

func TestProduct_Validate(t *testing.T) {
	tests := []struct {
		name    string
		product Product
		want    bool
	}{
		{
			name:    "Valid Product",
			product: Product{Name: "Widget", Price: 9.99},
			want:    true,
		},
		{
			name:    "Empty Name",
			product: Product{Name: "", Price: 10.0},
			want:    false,
		},
		{
			name:    "Negative Price",
			product: Product{Name: "Widget", Price: -5.0},
			want:    false,
		},
		{
			name:    "Price Zero",
			product: Product{Name: "Free Widget", Price: 0.0},
			want:    true, // Falls dein Model Preis > 0 verlangt
		},
		{
			name:    "Name with only spaces",
			product: Product{Name: "   ", Price: 5.0},
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.product.Validate(); got != tt.want {
				t.Errorf("Product.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
