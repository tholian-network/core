package distributions_test

import (
	"strings"
	"testing"

	"tholian-core/distributions/netbsd"
)

func TestNetBSDKeywords(t *testing.T) {

	// NetBSD does not ship /etc/os-release upstream; this fixture represents
	// the expected os-release format used by our matchers.
	osRelease := []byte(strings.Join([]string{
		`NAME="NetBSD"`,
		`ID=netbsd`,
		`VERSION_ID="10.0"`,
		`PRETTY_NAME="NetBSD 10.0"`,
		`HOME_URL="https://www.netbsd.org/"`,
		`SUPPORT_URL="https://www.netbsd.org/support/"`,
		`BUG_REPORT_URL="https://www.netbsd.org/support/send-pr.html"`,
	}, "\n"))

	assertKeywords(t, netbsd.Keywords, osRelease, map[string]string{
		"%release%": "10.0",
	})

}
