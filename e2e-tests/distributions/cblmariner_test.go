package distributions_test

import (
	"strings"
	"testing"

	"tholian-core/distributions/cblmariner"
)

func TestCBLMarinerKeywords(t *testing.T) {

	osRelease := []byte(strings.Join([]string{
		`NAME="Common Base Linux Mariner"`,
		`VERSION="2.0.20230208"`,
		`ID=mariner`,
		`VERSION_ID="2.0"`,
		`PRETTY_NAME="CBL-Mariner/Linux"`,
		`ANSI_COLOR="1;34"`,
		`HOME_URL="https://aka.ms/cbl-mariner"`,
		`BUG_REPORT_URL="https://aka.ms/cbl-mariner"`,
		`SUPPORT_URL="https://aka.ms/cbl-mariner"`,
	}, "\n"))

	assertKeywords(t, cblmariner.Keywords, osRelease, map[string]string{
		"%version%": "2.0",
	})

}
