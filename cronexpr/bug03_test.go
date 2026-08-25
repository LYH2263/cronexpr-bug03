package cronexpr

import (
	"errors"
	"testing"
)

func TestBug03_ValidateKeepsInvalidChain(t *testing.T) {
	err := ValidateFields("0 0 *")
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, ErrInvalid) {
		t.Fatalf("want ErrInvalid, got %v", err)
	}
}
