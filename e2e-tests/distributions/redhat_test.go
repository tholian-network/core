package distributions_test

import (
	"strings"
	"testing"

	"tholian-core/distributions/redhat"
)

func TestRedHatKeywords(t *testing.T) {

	t.Run("7", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="Red Hat Enterprise Linux Server"`,
			`VERSION="7.5 (Maipo)"`,
			`ID="rhel"`,
			`ID_LIKE="fedora"`,
			`VARIANT="Server"`,
			`VARIANT_ID="server"`,
			`VERSION_ID="7.5"`,
			`PRETTY_NAME="Red Hat Enterprise Linux Server 7.5 (Maipo)"`,
			`ANSI_COLOR="0;31"`,
			`CPE_NAME="cpe:/o:redhat:enterprise_linux:7.5:GA:server"`,
			`HOME_URL="https://www.redhat.com/"`,
			`BUG_REPORT_URL="https://bugzilla.redhat.com/"`,
		}, "\n"))

		assertKeywords(t, redhat.KeywordsLegacy, osRelease, map[string]string{
			"%version%": "7",
		})

	})

	t.Run("9", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="Red Hat Enterprise Linux"`,
			`VERSION="9.6 (Plow)"`,
			`ID="rhel"`,
			`ID_LIKE="fedora"`,
			`VERSION_ID="9.6"`,
			`PLATFORM_ID="platform:el9"`,
			`PRETTY_NAME="Red Hat Enterprise Linux 9.6 (Plow)"`,
			`ANSI_COLOR="0;31"`,
			`LOGO="fedora-logo-icon"`,
			`CPE_NAME="cpe:/o:redhat:enterprise_linux:9::baseos"`,
			`HOME_URL="https://www.redhat.com/"`,
			`BUG_REPORT_URL="https://issues.redhat.com/"`,
		}, "\n"))

		assertKeywords(t, redhat.Keywords, osRelease, map[string]string{
			"%version%": "9",
		})

	})

}
