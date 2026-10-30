package distributions_test

import (
	"strings"
	"testing"

	"tholian-core/distributions/freebsd"
)

func TestFreeBSDKeywords(t *testing.T) {

	osRelease := []byte(strings.Join([]string{
		`NAME=FreeBSD`,
		`VERSION="14.0-RELEASE-p5"`,
		`VERSION_ID="14.0"`,
		`ID=freebsd`,
		`ANSI_COLOR="0;31"`,
		`PRETTY_NAME="FreeBSD 14.0-RELEASE-p5"`,
		`CPE_NAME="cpe:/o:freebsd:freebsd:14.0"`,
		`HOME_URL="https://www.FreeBSD.org"`,
		`DOCUMENTATION_URL="https://docs.FreeBSD.org"`,
		`SUPPORT_URL="https://www.FreeBSD.org/support"`,
		`BUG_REPORT_URL="https://bugs.FreeBSD.org"`,
	}, "\n"))

	assertKeywords(t, freebsd.Keywords, osRelease, map[string]string{
		"%release%": "14.0",
	})

}
