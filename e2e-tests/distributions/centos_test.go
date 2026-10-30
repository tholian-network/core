package distributions_test

import (
	"strings"
	"testing"

	"tholian-core/distributions/centos"
)

func TestCentOSKeywords(t *testing.T) {

	t.Run("7", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="CentOS Linux"`,
			`VERSION="7 (Core)"`,
			`ID="centos"`,
			`ID_LIKE="rhel fedora"`,
			`VERSION_ID="7"`,
			`PRETTY_NAME="CentOS Linux 7 (Core)"`,
			`ANSI_COLOR="0;31"`,
			`CPE_NAME="cpe:/o:centos:centos:7"`,
			`HOME_URL="https://www.centos.org/"`,
			`BUG_REPORT_URL="https://bugs.centos.org/"`,
		}, "\n"))

		assertKeywords(t, centos.KeywordsLegacy, osRelease, map[string]string{
			"%version%": "7",
		})

	})

	t.Run("9", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="CentOS Stream"`,
			`VERSION="9"`,
			`ID="centos"`,
			`ID_LIKE="rhel fedora"`,
			`VERSION_ID="9"`,
			`PLATFORM_ID="platform:el9"`,
			`PRETTY_NAME="CentOS Stream 9"`,
			`ANSI_COLOR="0;31"`,
			`LOGO="fedora-logo-icon"`,
			`CPE_NAME="cpe:/o:centos:centos:9"`,
			`HOME_URL="https://centos.org/"`,
			`BUG_REPORT_URL="https://issues.redhat.com/"`,
		}, "\n"))

		assertKeywords(t, centos.Keywords, osRelease, map[string]string{
			"%version%": "9",
		})

	})

}
