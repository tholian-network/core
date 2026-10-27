package structs

import "tholian-core/matchers"

func equalPackageUnresolved(value_a matchers.Unresolved, value_b matchers.Unresolved) bool {

	if len(value_a.Candidates) != len(value_b.Candidates) {
		return false
	}

	for c := 0; c < len(value_a.Candidates); c++ {

		if value_a.Candidates[c].Equal(value_b.Candidates[c]) == false {
			return false
		}

	}

	return true

}

