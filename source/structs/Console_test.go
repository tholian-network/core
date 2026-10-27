package structs

import (
	"os"
	"strings"
	"sync"
	"testing"
)

func consoleTestFiles(t *testing.T) (*os.File, *os.File) {

	t.Helper()

	directory := t.TempDir()

	stdout, err := os.CreateTemp(directory, "stdout-*")

	if err != nil {
		t.Fatalf("could not create stdout file: %v", err)
	}

	stderr, err := os.CreateTemp(directory, "stderr-*")

	if err != nil {
		_ = stdout.Close()
		t.Fatalf("could not create stderr file: %v", err)
	}

	t.Cleanup(func() {
		_ = stdout.Close()
		_ = stderr.Close()
	})

	return stdout, stderr

}

func consoleTestOutput(t *testing.T, file *os.File) string {

	t.Helper()

	if err := file.Sync(); err != nil {
		t.Fatalf("could not sync output file: %v", err)
	}

	data, err := os.ReadFile(file.Name())

	if err != nil {
		t.Fatalf("could not read output file: %v", err)
	}

	return string(data)

}

func TestConsoleConcurrentAccess(t *testing.T) {

	console := NewConsole(nil, nil, 0)

	var writers sync.WaitGroup
	var readers sync.WaitGroup

	for w := 0; w < 16; w++ {

		writers.Add(1)

		go func(id int) {

			defer writers.Done()

			for i := 0; i < 200; i++ {

				console.Log("log " + string(rune('a'+id)))
				console.Info("info")
				console.Warn("warn")
				console.Error("error")
				console.Group("group")
				console.GroupEnd("group")
				console.Progress("progress")
				console.SetOffset(1)
				console.DisableMethod("Info")
				console.EnableMethod("Info")
				console.Clear("group")

			}

		}(w)

	}

	for r := 0; r < 8; r++ {

		readers.Add(1)

		go func() {

			defer readers.Done()

			for i := 0; i < 500; i++ {
				_ = console.GetOffset()
				_ = console.Length()
				_ = console.HasErrors()
				_ = console.UseColors()
				_ = console.UseMethod("Log")
			}

		}()

	}

	writers.Wait()
	readers.Wait()

	if console.Length() == 0 {
		t.Fatal("expected messages to be recorded")
	}

}

func TestConsoleRenderInspect(t *testing.T) {

	source := NewConsole(nil, nil, 0)
	target := NewConsole(nil, nil, 0)

	source.Inspect(map[string]string{"key": "value"})
	source.Render(target)

	if target.Length() != 1 {
		t.Fatalf("expected 1 rendered message, got %d", target.Length())
	}

	message := target.Messages[0]

	if message.Method != "Inspect" {
		t.Fatalf("expected Inspect message, got %q", message.Method)
	}

	if message.Value == "" {
		t.Fatal("Inspect message was empty")
	}

}

func TestConsoleRenderConcurrent(t *testing.T) {

	source := NewConsole(nil, nil, 0)
	target := NewConsole(nil, nil, 0)

	source.Log("one")
	source.Info("two")
	source.Warn("three")

	var waitgroup sync.WaitGroup

	for r := 0; r < 8; r++ {

		waitgroup.Add(1)

		go func() {
			defer waitgroup.Done()
			source.Render(target)
		}()

	}

	waitgroup.Wait()

	if target.Length() == 0 {
		t.Fatal("expected rendered messages")
	}

}

func TestConsoleNewConsoleDefaults(t *testing.T) {

	console := NewConsole(nil, nil, -5)

	if console.GetOffset() != 0 {
		t.Fatalf("expected negative offset to be clamped, got %d", console.GetOffset())
	}

	if console.Messages == nil {
		t.Fatal("expected initialized messages")
	}

	if console.Stdout != nil || console.Stderr != nil {
		t.Fatal("expected nil output files to remain nil")
	}

	if console.UseColors() == true {
		t.Fatal("expected colors to be disabled without output files")
	}

	for _, method := range []string{"Group", "Log", "Info", "Warn", "Error", "Progress"} {

		if console.UseMethod(method) == false {
			t.Fatalf("expected method %q to be enabled by default", method)
		}

	}

	for _, method := range []string{"Inspect", "Clear", "GroupEnd", "GroupEndResult"} {

		if console.UseMethod(method) == true {
			t.Fatalf("expected method %q to be unknown by default", method)
		}

	}

	offset := NewConsole(nil, nil, 2)

	if offset.GetOffset() != 2 {
		t.Fatalf("expected offset 2, got %d", offset.GetOffset())
	}

}

func TestConsoleOffset(t *testing.T) {

	console := NewConsole(nil, nil, 3)

	if console.SetOffset(-1) == true {
		t.Fatal("expected negative offset to be rejected")
	}

	if console.GetOffset() != 3 {
		t.Fatalf("expected offset to remain 3, got %d", console.GetOffset())
	}

	if console.SetOffset(0) == false {
		t.Fatal("expected zero offset to be accepted")
	}

	if console.GetOffset() != 0 {
		t.Fatalf("expected offset 0, got %d", console.GetOffset())
	}

}

func TestConsoleMethods(t *testing.T) {

	console := NewConsole(nil, nil, 0)

	if console.DisableMethod("Info") == false {
		t.Fatal("expected Info to be disabled")
	}

	if console.UseMethod("Info") == true {
		t.Fatal("expected Info to be disabled")
	}

	if console.EnableMethod("Info") == false {
		t.Fatal("expected Info to be enabled")
	}

	if console.UseMethod("Info") == false {
		t.Fatal("expected Info to be enabled")
	}

	if console.DisableMethod("Unknown") == true {
		t.Fatal("expected unknown method to be rejected")
	}

	if console.EnableMethod("Unknown") == true {
		t.Fatal("expected unknown method to be rejected")
	}

}

func TestConsoleHasErrorsAndReset(t *testing.T) {

	build := func() *Console {

		console := NewConsole(nil, nil, 4)
		console.Log("log")
		console.Error("error")
		console.Warn("warn")

		return console

	}

	tests := []struct {
		name       string
		offset     bool
		errors     bool
		warnings   bool
		wantOffset int
		want       []string
		wantErrors bool
	}{
		{"reset all", true, true, true, 0, nil, false},
		{"keep errors and warnings", false, false, false, 4, []string{"Error", "Warn"}, true},
		{"keep warnings only", false, true, false, 4, []string{"Warn"}, false},
		{"keep errors only", false, false, true, 4, []string{"Error"}, true},
		{"keep warnings reset offset", true, true, false, 0, []string{"Warn"}, false},
		{"keep errors reset offset", true, false, true, 0, []string{"Error"}, true},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			console := build()

			if console.HasErrors() == false {
				t.Fatal("expected errors before reset")
			}

			console.Reset(test.offset, test.errors, test.warnings)

			if console.GetOffset() != test.wantOffset {
				t.Fatalf("expected offset %d, got %d", test.wantOffset, console.GetOffset())
			}

			if console.Length() != len(test.want) {
				t.Fatalf("expected %d messages, got %d", len(test.want), console.Length())
			}

			for index, method := range test.want {

				if console.Messages[index].Method != method {
					t.Fatalf("expected message %d to be %q, got %q", index, method, console.Messages[index].Method)
				}

			}

			if console.HasErrors() != test.wantErrors {
				t.Fatalf("expected HasErrors() to be %t, got %t", test.wantErrors, console.HasErrors())
			}

		})

	}

}

func TestConsoleLengthAndClearScreen(t *testing.T) {

	stdout, stderr := consoleTestFiles(t)
	console := NewConsole(stdout, stderr, 0)

	if console.Length() != 0 {
		t.Fatalf("expected empty console, got %d messages", console.Length())
	}

	console.ClearScreen()

	if console.Length() != 1 {
		t.Fatalf("expected 1 message, got %d", console.Length())
	}

	if console.Messages[0].Method != "Clear" {
		t.Fatalf("expected Clear message, got %q", console.Messages[0].Method)
	}

	output := consoleTestOutput(t, stdout)

	if strings.Contains(output, "\u001b[2J") == false ||
		strings.Contains(output, "\u001b[0f") == false ||
		strings.Contains(output, "\u001b[3J") == false {
		t.Fatalf("expected clear screen escape sequences, got %q", output)
	}

	console.Error("boom")

	if console.HasErrors() == false {
		t.Fatal("expected errors to be recorded")
	}

	beforeLength := console.Length()
	console.ClearScreen()

	if console.Length() != beforeLength+1 {
		t.Fatalf("expected %d messages, got %d", beforeLength+1, console.Length())
	}

	if after := consoleTestOutput(t, stdout); after != output {
		t.Fatal("expected ClearScreen to skip terminal output when errors exist")
	}

}

func TestConsoleClearGroup(t *testing.T) {

	t.Run("without group end", func(t *testing.T) {

		console := NewConsole(nil, nil, 0)
		console.Group("Alpha")
		console.Log("inside")
		console.Clear("Alpha\nignored")

		if console.Length() != 1 {
			t.Fatalf("expected 1 message, got %d", console.Length())
		}

		if console.Messages[0].Method != "Group" || console.Messages[0].Value != "Alpha" {
			t.Fatalf("expected preserved group, got %#v", console.Messages[0])
		}

		if console.GetOffset() != 1 {
			t.Fatalf("expected offset 1 after re-rendering group, got %d", console.GetOffset())
		}

	})

	t.Run("with group end", func(t *testing.T) {

		console := NewConsole(nil, nil, 0)
		console.Group("Alpha")
		console.Log("inside")
		console.GroupEnd("Alpha")

		before := console.Length()
		console.Clear("Alpha")

		if console.Length() != before {
			t.Fatalf("expected %d messages, got %d", before, console.Length())
		}

	})

	t.Run("unknown group", func(t *testing.T) {

		console := NewConsole(nil, nil, 0)
		console.Log("inside")

		before := console.Length()
		console.Clear("Missing")

		if console.Length() != before {
			t.Fatalf("expected %d messages, got %d", before, console.Length())
		}

	})

}

func TestConsoleGroupOffsets(t *testing.T) {

	console := NewConsole(nil, nil, 0)

	console.Group("outer")

	if console.GetOffset() != 1 {
		t.Fatalf("expected offset 1, got %d", console.GetOffset())
	}

	console.Group("inner")

	if console.GetOffset() != 2 {
		t.Fatalf("expected offset 2, got %d", console.GetOffset())
	}

	console.GroupEnd("inner")

	if console.GetOffset() != 1 {
		t.Fatalf("expected offset 1, got %d", console.GetOffset())
	}

	console.GroupEndResult(true, "outer")

	if console.GetOffset() != 0 {
		t.Fatalf("expected offset 0, got %d", console.GetOffset())
	}

	last := console.Messages[console.Length()-1]

	if last.Method != "GroupEnd" || last.Value != "outer succeeded" {
		t.Fatalf("expected successful group end, got %#v", last)
	}

	console.GroupEndResult(false, "nothing")

	last = console.Messages[console.Length()-1]

	if last.Method != "GroupEnd" || last.Value != "nothing failed" {
		t.Fatalf("expected failed group end, got %#v", last)
	}

	if console.GetOffset() != 0 {
		t.Fatalf("expected offset to remain 0, got %d", console.GetOffset())
	}

	console.GroupEnd("extra")

	if console.GetOffset() != 0 {
		t.Fatalf("expected offset to remain 0, got %d", console.GetOffset())
	}

	console.DisableMethod("Group")

	beforeLength := console.Length()
	beforeOffset := console.GetOffset()

	console.Group("disabled")
	console.GroupEnd("disabled")
	console.GroupEndResult(true, "disabled")

	if console.Length() != beforeLength || console.GetOffset() != beforeOffset {
		t.Fatal("expected disabled Group method to be ignored")
	}

}

func TestConsoleInspect(t *testing.T) {

	console := NewConsole(nil, nil, 0)

	console.Inspect("plain")

	if console.Length() != 1 {
		t.Fatalf("expected 1 message, got %d", console.Length())
	}

	if console.Messages[0].Method != "Inspect" || console.Messages[0].Value != "plain" {
		t.Fatalf("expected string inspect message, got %#v", console.Messages[0])
	}

	console.Inspect(map[string]string{"key": "value"})

	if console.Length() != 2 {
		t.Fatalf("expected 2 messages, got %d", console.Length())
	}

	message := console.Messages[1]

	if message.Method != "Inspect" {
		t.Fatalf("expected Inspect message, got %q", message.Method)
	}

	if strings.Contains(message.Value, "\"key\": \"value\"") == false {
		t.Fatalf("expected JSON inspect value, got %q", message.Value)
	}

	console.Inspect(make(chan int))

	if console.Length() != 3 {
		t.Fatalf("expected 3 messages, got %d", console.Length())
	}

	if console.Messages[2].Method != "Error" {
		t.Fatalf("expected Error message for unsupported type, got %q", console.Messages[2].Method)
	}

	console.DisableMethod("Log")
	console.Inspect("hidden")

	if console.Length() != 3 {
		t.Fatalf("expected Inspect to respect disabled Log method, got %d messages", console.Length())
	}

}

func TestConsoleLogLevels(t *testing.T) {

	call := func(console *Console, method string, value string) {

		switch method {
		case "Log":
			console.Log(value)
		case "Error":
			console.Error(value)
		case "Info":
			console.Info(value)
		case "Progress":
			console.Progress(value)
		case "Warn":
			console.Warn(value)
		}

	}

	for _, method := range []string{"Log", "Error", "Info", "Progress", "Warn"} {

		t.Run(method, func(t *testing.T) {

			console := NewConsole(nil, nil, 0)

			call(console, method, "baseline")

			if console.Length() != 1 {
				t.Fatalf("expected 1 message, got %d", console.Length())
			}

			if console.Messages[0].Method != method {
				t.Fatalf("expected %q, got %q", method, console.Messages[0].Method)
			}

			console.DisableMethod(method)
			call(console, method, "hidden")

			if console.Length() != 1 {
				t.Fatalf("expected disabled %q to be ignored, got %d messages", method, console.Length())
			}

			console.EnableMethod(method)
			call(console, method, "shown")

			if console.Length() != 2 {
				t.Fatalf("expected enabled %q to be recorded, got %d messages", method, console.Length())
			}

		})

	}

}

func TestConsoleProgressReplacement(t *testing.T) {

	console := NewConsole(nil, nil, 0)

	console.Progress("Parsing 1 of 5 Incidents")

	if console.Length() != 1 {
		t.Fatalf("expected 1 message, got %d", console.Length())
	}

	console.Progress("Parsing 2 of 5 Incidents")

	if console.Length() != 1 {
		t.Fatalf("expected matching progress to replace the last message, got %d", console.Length())
	}

	if console.Messages[0].Value != "Parsing 2 of 5 Incidents" {
		t.Fatalf("expected replaced progress value, got %q", console.Messages[0].Value)
	}

	console.Progress("Parsing 3 of 6 Incidents")

	if console.Length() != 1 {
		t.Fatalf("expected matching prefix and suffix to replace the last message, got %d", console.Length())
	}

	console.Progress("Other task")

	if console.Length() != 2 {
		t.Fatalf("expected non-matching progress to be appended, got %d", console.Length())
	}

	console.Progress("Other task")

	if console.Length() != 2 {
		t.Fatalf("expected equal progress to replace the last message, got %d", console.Length())
	}

	console.Progress("Export foo")

	if console.Length() != 3 {
		t.Fatalf("expected non-matching progress to be appended, got %d", console.Length())
	}

	console.Progress("Export bar")

	if console.Length() != 3 {
		t.Fatalf("expected matching Export progress to replace the last message, got %d", console.Length())
	}

	if console.Messages[2].Value != "Export bar" {
		t.Fatalf("expected replaced Export progress value, got %q", console.Messages[2].Value)
	}

}

func TestConsoleRenderDispatch(t *testing.T) {

	source := NewConsole(nil, nil, 0)

	source.Log("log")
	source.Info("info")
	source.Warn("warn")
	source.Error("error")
	source.Group("group")
	source.GroupEnd("group")
	source.Progress("progress")
	source.Inspect("inspect")

	target := NewConsole(nil, nil, 0)
	source.Render(target)

	if target.Length() != source.Length() {
		t.Fatalf("expected %d rendered messages, got %d", source.Length(), target.Length())
	}

	for index := 0; index < source.Length(); index++ {

		if target.Messages[index].Method != source.Messages[index].Method {
			t.Fatalf("expected method %q at %d, got %q", source.Messages[index].Method, index, target.Messages[index].Method)
		}

		if target.Messages[index].Value != source.Messages[index].Value {
			t.Fatalf("expected value %q at %d, got %q", source.Messages[index].Value, index, target.Messages[index].Value)
		}

	}

}

func TestConsoleRenderClearDispatch(t *testing.T) {

	source := NewConsole(nil, nil, 0)
	source.Group("group")
	source.Log("inside")
	source.Messages = append(source.Messages, NewConsoleMessage("Clear", "group"))

	target := NewConsole(nil, nil, 0)
	source.Render(target)

	if target.Length() != 1 {
		t.Fatalf("expected 1 rendered message, got %d", target.Length())
	}

	if target.Messages[0].Method != "Group" || target.Messages[0].Value != "group" {
		t.Fatalf("expected preserved group, got %#v", target.Messages[0])
	}

}

func TestConsoleColors(t *testing.T) {

	tests := []struct {
		name    string
		term    string
		noColor string
		want    bool
	}{
		{"xterm", "xterm", "", true},
		{"xterm 16 color", "xterm-16color", "", true},
		{"xterm 88 color", "xterm-88color", "", true},
		{"xterm 256 color", "xterm-256color", "", true},
		{"xterm kitty", "xterm-kitty", "", true},
		{"unsupported term", "dumb", "", false},
		{"no color 1", "xterm-256color", "1", false},
		{"no color yes", "xterm", "yes", false},
		{"no color true", "xterm", "true", false},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			t.Setenv("TERM", test.term)
			t.Setenv("NO_COLOR", test.noColor)

			stdout, stderr := consoleTestFiles(t)
			console := NewConsole(stdout, stderr, 0)

			if console.UseColors() != test.want {
				t.Fatalf("expected UseColors() to be %t, got %t", test.want, console.UseColors())
			}

		})

	}

}
