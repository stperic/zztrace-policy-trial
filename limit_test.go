package trial

import "testing"

func TestAllowed(t *testing.T) {
	if !Allowed(40) || Allowed(41) {
		t.Fatal("limit boundary is wrong")
	}
}

func TestFixedValue(t *testing.T) {
	if FixedValue() != 1 {
		t.Fatal("fixed value changed")
	}
}
