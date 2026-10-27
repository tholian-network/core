package net

import "testing"

func TestSupportsIPv6(t *testing.T) {

	first := SupportsIPv6()

	var _ bool = first

	for i := 0; i < 3; i++ {

		if result := SupportsIPv6(); result != first {
			t.Fatalf("SupportsIPv6() is not stable: %t != %t", result, first)
		}

	}

}
