package trial

import "testing"

func TestAllowed(t *testing.T) {
	if !Allowed(40) || Allowed(41) {
		t.Fatal("limit boundary is wrong")
	}
	if Allowed(50) {
		t.Fatal("the previous limit must be rejected")
	}
}

func TestFixedValue(t *testing.T) {
	if FixedValue() != 1 {
		t.Fatal("fixed value changed")
	}
}
