package matchers

import "sort"
import "strings"
import "testing"

import "tholian-core/distributions/almalinux"
import "tholian-core/distributions/alpinelinux"
import "tholian-core/distributions/amazonlinux"
import "tholian-core/distributions/cblmariner"
import "tholian-core/distributions/centos"
import "tholian-core/distributions/debian"
import "tholian-core/distributions/fedora"
import "tholian-core/distributions/freebsd"
import "tholian-core/distributions/linuxmint"
import "tholian-core/distributions/netbsd"
import "tholian-core/distributions/opensuse"
import "tholian-core/distributions/openwrt"
import "tholian-core/distributions/oraclelinux"
import "tholian-core/distributions/photonos"
import "tholian-core/distributions/redhat"
import "tholian-core/distributions/rockylinux"
import "tholian-core/distributions/suse"
import "tholian-core/distributions/trisquel"
import "tholian-core/distributions/ubuntu"

type distributionFixture struct {
	Names    []string
	Releases []string
}

func intReleaseNames(releases map[string]int) []string {

	names := make([]string, 0, len(releases))

	for name := range releases {
		names = append(names, name)
	}

	sort.Strings(names)

	return names

}

func floatReleaseNames(releases map[string]float32) []string {

	names := make([]string, 0, len(releases))

	for name := range releases {
		names = append(names, name)
	}

	sort.Strings(names)

	return names

}

func distributionFixtures() []distributionFixture {

	return []distributionFixture{
		{Names: []string{"almalinux"}, Releases: intReleaseNames(almalinux.Releases)},
		{Names: []string{"alpinelinux", "alpine"}, Releases: floatReleaseNames(alpinelinux.Releases)},
		{Names: []string{"amazonlinux", "amzn"}, Releases: intReleaseNames(amazonlinux.Releases)},
		{Names: []string{"antergos"}, Releases: []string{}},
		{Names: []string{"archlinux", "arch"}, Releases: []string{}},
		{Names: []string{"cblmariner", "mariner"}, Releases: floatReleaseNames(cblmariner.Releases)},
		{Names: []string{"centos"}, Releases: intReleaseNames(centos.Releases)},
		{Names: []string{"debian"}, Releases: intReleaseNames(debian.Releases)},
		{Names: []string{"fedora"}, Releases: intReleaseNames(fedora.Releases)},
		{Names: []string{"freebsd"}, Releases: floatReleaseNames(freebsd.Releases)},
		{Names: []string{"gentoolinux", "gentoo"}, Releases: []string{}},
		{Names: []string{"linuxmint"}, Releases: floatReleaseNames(linuxmint.Releases)},
		{Names: []string{"manjaro", "manjaro-arm"}, Releases: []string{}},
		{Names: []string{"netbsd"}, Releases: floatReleaseNames(netbsd.Releases)},
		{Names: []string{"openbsd"}, Releases: []string{}},
		{Names: []string{"opensuse"}, Releases: floatReleaseNames(opensuse.Releases)},
		{Names: []string{"opensuse-leap", "opensuse-tumbleweed"}, Releases: []string{}},
		{Names: []string{"openwrt"}, Releases: floatReleaseNames(openwrt.Releases)},
		{Names: []string{"oraclelinux"}, Releases: intReleaseNames(oraclelinux.Releases)},
		{Names: []string{"photonos"}, Releases: floatReleaseNames(photonos.Releases)},
		{Names: []string{"redhat"}, Releases: intReleaseNames(redhat.Releases)},
		{Names: []string{"rockylinux"}, Releases: intReleaseNames(rockylinux.Releases)},
		{Names: []string{"suse"}, Releases: intReleaseNames(suse.Releases)},
		{Names: []string{"trisquel"}, Releases: floatReleaseNames(trisquel.Releases)},
		{Names: []string{"ubuntu"}, Releases: floatReleaseNames(ubuntu.Releases)},
	}

}

func TestVulnerabilityMatchesDistribution(t *testing.T) {

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

			vulnerability := NewVulnerability()
			vulnerability.Distribution = test.matcher

			if result := vulnerability.MatchesDistribution(test.value); result != test.expected {
				t.Errorf("Expected %q to match %q to be %t, got %t", test.matcher, test.value, test.expected, result)
			}

		})

	}

}

func TestVulnerabilityMatchesDistributionFixtures(t *testing.T) {

	for _, fixture := range distributionFixtures() {

		for _, name := range fixture.Names {

			matchers := []string{name, name + "-any", name + "-*"}

			for _, matcher := range matchers {

				vulnerability := NewVulnerability()
				vulnerability.Distribution = matcher

				if vulnerability.MatchesDistribution(name) == false {
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

					vulnerability := NewVulnerability()
					vulnerability.Distribution = matcher

					if vulnerability.MatchesDistribution(value) == false {
						t.Errorf("Expected %q to match %q", matcher, value)
					}

				}

				for _, other := range fixture.Releases {

					if other == release {
						continue
					}

					vulnerability := NewVulnerability()
					vulnerability.Distribution = value

					if vulnerability.MatchesDistribution(name+"-"+other) == true {
						t.Errorf("Expected %q not to match %q", value, name+"-"+other)
					}

				}

			}

		}

	}

}
