package net

import "testing"

func TestSupportsIPv4(t *testing.T) {

	first := SupportsIPv4()

	var _ bool = first

	for i := 0; i < 3; i++ {

		if result := SupportsIPv4(); result != first {
			t.Fatalf("SupportsIPv4() is not stable: %t != %t", result, first)
		}

	}

}
