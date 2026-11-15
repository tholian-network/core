package methods

import "regexp"
import "testing"

var methodPattern = regexp.MustCompile(`^tholian-[a-z0-9-]+\.[A-Z][A-Za-z0-9]*$`)

func TestAllMethodsAreUniqueAndWellFormed(t *testing.T) {

	seen := make(map[Method]bool)

	for _, method := range AllMethods() {

		if methodPattern.MatchString(method.String()) == false {
			t.Fatalf("method %q does not match the project-qualified pattern", method)
		}

		if seen[method] == true {
			t.Fatalf("method %q is defined more than once", method)
		}

		seen[method] = true

		if method.Project() == "" || method.Action() == "" {
			t.Fatalf("method %q cannot be split into project and action", method)
		}

		if method.IsValid() == false {
			t.Fatalf("method %q is not valid", method)
		}

	}

	if len(seen) == 0 {
		t.Fatal("no methods were registered")
	}

}
