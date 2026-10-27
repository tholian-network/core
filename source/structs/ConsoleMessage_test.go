package structs

import "slices"
import "strings"
import "testing"

func consoleMessageTestCaller() ConsoleMessage {
	return NewConsoleMessage("Log", "a\tb")
}

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

func TestConsoleMessageMatchesBranches(t *testing.T) {

	tests := []struct {
		name    string
		message ConsoleMessage
		other   ConsoleMessage
		want    bool
	}{
		{
			"progress matching counters",
			NewConsoleMessage("Progress", "Parsing 1 of 5 Incidents"),
			NewConsoleMessage("Progress", "Parsing 2 of 5 Incidents"),
			true,
		},
		{
			"progress different prefix",
			NewConsoleMessage("Progress", "Parsing 1 of 5 Incidents"),
			NewConsoleMessage("Progress", "Scanning 2 of 5 Incidents"),
			false,
		},
		{
			"progress different suffix",
			NewConsoleMessage("Progress", "Parsing 1 of 5 Incidents"),
			NewConsoleMessage("Progress", "Parsing 2 of 5 Records"),
			false,
		},
		{
			"progress with of vs without",
			NewConsoleMessage("Progress", "Parsing 1 of 5 Incidents"),
			NewConsoleMessage("Progress", "Parsing Incidents"),
			true,
		},
		{
			"progress with of vs without mismatch",
			NewConsoleMessage("Progress", "Parsing 1 of 5 Incidents"),
			NewConsoleMessage("Progress", "Parsing things"),
			false,
		},
		{
			"progress without of vs with of",
			NewConsoleMessage("Progress", "Parsing Incidents"),
			NewConsoleMessage("Progress", "Parsing 1 of 5 Incidents"),
			false,
		},
		{
			"progress equal values",
			NewConsoleMessage("Progress", "Loading"),
			NewConsoleMessage("Progress", "Loading"),
			true,
		},
		{
			"export prefix",
			NewConsoleMessage("Progress", "Export one"),
			NewConsoleMessage("Progress", "Export two"),
			true,
		},
		{
			"parse prefix",
			NewConsoleMessage("Progress", "Parse one"),
			NewConsoleMessage("Progress", "Parse two"),
			true,
		},
		{
			"update prefix",
			NewConsoleMessage("Progress", "Update one"),
			NewConsoleMessage("Progress", "Update two"),
			true,
		},
		{
			"different method",
			NewConsoleMessage("Log", "same"),
			NewConsoleMessage("Info", "same"),
			false,
		},
		{
			"different value same method",
			NewConsoleMessage("Log", "one"),
			NewConsoleMessage("Log", "two"),
			false,
		},
		{
			"same non-progress",
			NewConsoleMessage("Log", "same"),
			NewConsoleMessage("Log", "same"),
			true,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			if test.message.Matches(test.other) != test.want {
				t.Fatalf("expected Matches() to be %t", test.want)
			}

		})

	}

}

func TestConsoleMessageNewConsoleMessage(t *testing.T) {

	message := consoleMessageTestCaller()

	if message.Method != "Log" {
		t.Fatalf("expected method Log, got %q", message.Method)
	}

	if message.Value != "a    b" {
		t.Fatalf("expected tabs to be replaced with spaces, got %q", message.Value)
	}

	if strings.Contains(message.Value, "\t") == true {
		t.Fatal("expected no tabs in message value")
	}

	if message.Time.IsZero() == true {
		t.Fatal("expected a timestamp")
	}

	if message.Caller.File == "" || message.Caller.Line <= 0 {
		t.Fatalf("expected captured caller, got %#v", message.Caller)
	}

	if strings.Contains(message.Caller.File, "ConsoleMessage_test.go") == false {
		t.Fatalf("expected caller file to be this test, got %q", message.Caller.File)
	}

}

func TestConsoleMessageLines(t *testing.T) {

	tests := []struct {
		name  string
		value string
		want  []string
	}{
		{"single plain", "hello", []string{" hello"}},
		{"single greater", ">hello", []string{"->hello"}},
		{"single dash", "-hello", []string{"--hello"}},
		{"single empty", "", []string{" "}},
		{"multiline plain", "first\nsecond", []string{" first", " second"}},
		{"multiline separators", ">first\n-second\nplain", []string{"->first", "--second", " plain"}},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			message := NewConsoleMessage("Log", test.value)
			result := message.Lines()

			if slices.Equal(result, test.want) == false {
				t.Fatalf("expected %#v, got %#v", test.want, result)
			}

		})

	}

}
