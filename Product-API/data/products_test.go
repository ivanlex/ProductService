package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "123",
		Price: 100,
		SKU:   "abc-abc-abc",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
