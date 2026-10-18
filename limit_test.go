package trial

import "testing"

func TestAllowed(t *testing.T) {
	if !Allowed(50) || Allowed(51) {
		t.Fatal("limit boundary is wrong")
	}
}

func TestFixedValue(t *testing.T) {
	if FixedValue() != 2 {
		t.Fatal("fixed value changed")
	}
}
