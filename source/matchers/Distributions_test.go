package matchers

import "strings"
import "testing"

func TestDistributionNew(t *testing.T) {

	distribution := NewDistribution()

	if distribution.Name != Any {
		t.Errorf("Expected Name to be %q, got %q", Any, distribution.Name)
	}

	if distribution.Version != Any {
		t.Errorf("Expected Version to be %q, got %q", Any, distribution.Version)
	}

	if distribution.Manager != Any {
		t.Errorf("Expected Manager to be %q, got %q", Any, distribution.Manager)
	}

	if distribution.Vendor != Any {
		t.Errorf("Expected Vendor to be %q, got %q", Any, distribution.Vendor)
	}

}

func TestDistributionToDistribution(t *testing.T) {

	distribution := ToDistribution("debian-bookworm")

	if distribution.Name != "debian-bookworm" {
		t.Errorf("Expected Name to be %q, got %q", "debian-bookworm", distribution.Name)
	}

	if distribution.Version != Any {
		t.Errorf("Expected Version to be %q, got %q", Any, distribution.Version)
	}

	if distribution.Manager != Any {
		t.Errorf("Expected Manager to be %q, got %q", Any, distribution.Manager)
	}

	if distribution.Vendor != Any {
		t.Errorf("Expected Vendor to be %q, got %q", Any, distribution.Vendor)
	}

}

func TestDistributionMatchesName(t *testing.T) {

	tests := []struct {
		name     string
		matcher  string
		value    string
		expected bool
	}{
		{name: "exact release", matcher: "debian-bookworm", value: "debian-bookworm", expected: true},
		{name: "exact release mismatch", matcher: "debian-bookworm", value: "debian-bullseye", expected: false},
		{name: "exact release other distribution", matcher: "debian-bookworm", value: "fedora-42", expected: false},
		{name: "any release matches release", matcher: "debian-any", value: "debian-bookworm", expected: true},
		{name: "any release matches distribution", matcher: "debian-any", value: "debian", expected: true},
		{name: "any release other distribution", matcher: "debian-any", value: "fedora-42", expected: false},
		{name: "star release matches release", matcher: "debian-*", value: "debian-bookworm", expected: true},
		{name: "star release matches distribution", matcher: "debian-*", value: "debian", expected: true},
		{name: "star release requires separator", matcher: "debian-*", value: "debianbookworm", expected: false},
		{name: "arch requires separator", matcher: "arch-any", value: "archlinux", expected: false},
		{name: "long release exact", matcher: "opensuse-leap-15.5", value: "opensuse-leap-15.5", expected: true},
		{name: "long release any", matcher: "opensuse-leap-any", value: "opensuse-leap-15.5", expected: true},
		{name: "long release star", matcher: "opensuse-leap-*", value: "opensuse-leap-15.5", expected: true},
		{name: "long release mismatch", matcher: "opensuse-leap-15.5", value: "opensuse-leap-15.4", expected: false},
		{name: "long release any distribution", matcher: "opensuse-any", value: "opensuse-leap-15.5", expected: true},
		{name: "suse desktop any", matcher: "suse-desktop-any", value: "suse-desktop-15", expected: true},
		{name: "suse desktop star", matcher: "suse-desktop-*", value: "suse-desktop-15", expected: true},
		{name: "suse desktop mismatch", matcher: "suse-desktop-15", value: "suse-server-15", expected: false},
		{name: "provider any matches all", matcher: Any, value: "debian-bookworm", expected: true},
		{name: "empty matcher", matcher: "", value: "debian-bookworm", expected: false},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			distribution := Distribution{Name: test.matcher}

			if result := distribution.MatchesName(test.value); result != test.expected {
				t.Errorf("Expected %q to match %q to be %t, got %t", test.matcher, test.value, test.expected, result)
			}

		})

	}

}

func TestDistributionMatchesNameFixtures(t *testing.T) {

	for _, fixture := range distributionFixtures() {

		for _, name := range fixture.Names {

			matchers := []string{name, name + "-any", name + "-*"}

			for _, matcher := range matchers {

				distribution := NewDistribution()
				distribution.SetName(matcher)

				if distribution.MatchesName(name) == false {
					t.Errorf("Expected %q to match %q", matcher, name)
				}

			}

			for _, release := range fixture.Releases {

				value := name + "-" + release

				matchers := []string{value, name + "-any", name + "-*"}

				index := strings.LastIndex(value, "-")

				if index > -1 {

					prefix := value[0:index]

					matchers = append(matchers, prefix+"-any", prefix+"-*")

				}

				for _, matcher := range matchers {

					distribution := NewDistribution()
					distribution.SetName(matcher)

					if distribution.MatchesName(value) == false {
						t.Errorf("Expected %q to match %q", matcher, value)
					}

				}

				for _, other := range fixture.Releases {

					if other == release {
						continue
					}

					distribution := NewDistribution()
					distribution.SetName(value)

					if distribution.MatchesName(name+"-"+other) == true {
						t.Errorf("Expected %q not to match %q", value, name+"-"+other)
					}

				}

			}

		}

	}

}

func TestDistributionMatches(t *testing.T) {

	distribution := NewDistribution()
	distribution.SetName("debian-*")
	distribution.Manager = "apt"
	distribution.SetVendor("debian")

	if distribution.Matches("debian-bookworm", "12", "apt", "debian") == false {
		t.Error("Expected distribution to match")
	}

	if distribution.Matches("debian-bookworm", "12", "pacman", "debian") == true {
		t.Error("Expected distribution not to match manager")
	}

	if distribution.Matches("fedora-42", "42", "apt", "debian") == true {
		t.Error("Expected distribution not to match name")
	}

}

func TestDistributionEqual(t *testing.T) {

	base := NewDistribution()
	value := NewDistribution()

	if base.Equal(value) == false {
		t.Error("Expected distributions to be equal")
	}

	value.SetName("debian-bookworm")

	if base.Equal(value) == true {
		t.Error("Expected distributions not to be equal")
	}

}

func TestDistributionSetName(t *testing.T) {

	tests := []struct {
		value    string
		expected string
	}{
		{value: "all", expected: Any},
		{value: "any", expected: Any},
		{value: "*", expected: Any},
		{value: "debian-any", expected: "debian-any"},
		{value: "debian-*", expected: "debian-*"},
		{value: "debian-bookworm", expected: "debian-bookworm"},
		{value: "  debian-bookworm  ", expected: "debian-bookworm"},
		{value: "", expected: Any},
	}

	for _, test := range tests {

		distribution := NewDistribution()
		distribution.SetName(test.value)

		if distribution.Name != test.expected {
			t.Errorf("Expected SetName(%q) to be %q, got %q", test.value, test.expected, distribution.Name)
		}

	}

}
