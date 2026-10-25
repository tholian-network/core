package structs

import "testing"

func TestConsoleMessageMatches(t *testing.T) {

	progress1 := NewConsoleMessage("Progress", "Parsing 1 of 5 Incidents")
	progress2 := NewConsoleMessage("Progress", "Parsing 2 of 5 Incidents")
	log := NewConsoleMessage("Log", "hello")
	other := NewConsoleMessage("Log", "world")

	if progress1.Matches(progress2) == false {
		t.Fatal("progress messages should match")
	}

	if log.Matches(log) == false {
		t.Fatal("identical messages should match")
	}

	if log.Matches(other) == true {
		t.Fatal("different log messages should not match")
	}

}
