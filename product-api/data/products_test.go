package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Nix",
		Price: 1.00,
		SKU:   "abs-abs-asss",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
