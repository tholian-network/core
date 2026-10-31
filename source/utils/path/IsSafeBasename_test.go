package path

import "testing"

func TestIsSafeBasename(t *testing.T) {

	if IsSafeBasename("") {
		t.Fatal("empty basename is not safe")
	}

	if IsSafeBasename(".") || IsSafeBasename("..") {
		t.Fatal("relative basenames are not safe")
	}

	if IsSafeBasename("a/b") || IsSafeBasename("a\\b") || IsSafeBasename("a\x00b") {
		t.Fatal("separators are not safe")
	}

	if IsSafeBasename("CVE-2026-1000") == false {
		t.Fatal("valid basename was rejected")
	}

	if IsSafeBasename("acme:widget:1.0") == false {
		t.Fatal("valid basename with colons was rejected")
	}

}
