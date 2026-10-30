package distributions_test

import (
	"strings"
	"testing"

	"tholian-core/distributions/oraclelinux"
)

func TestOracleLinuxKeywords(t *testing.T) {

	t.Run("7", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="Oracle Linux Server"`,
			`VERSION="7.9"`,
			`ID="ol"`,
			`ID_LIKE="fedora"`,
			`VARIANT="Server"`,
			`VARIANT_ID="server"`,
			`VERSION_ID="7.9"`,
			`PRETTY_NAME="Oracle Linux Server 7.9"`,
			`ANSI_COLOR="0;31"`,
			`CPE_NAME="cpe:/o:oracle:linux:7:9:server"`,
			`HOME_URL="https://linux.oracle.com/"`,
			`BUG_REPORT_URL="https://bugzilla.oracle.com/"`,
		}, "\n"))

		assertKeywords(t, oraclelinux.KeywordsLegacy, osRelease, map[string]string{
			"%version%": "7",
		})

	})

	t.Run("9", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="Oracle Linux Server"`,
			`VERSION="9.1"`,
			`ID="ol"`,
			`ID_LIKE="fedora"`,
			`VARIANT="Server"`,
			`VARIANT_ID="server"`,
			`VERSION_ID="9.1"`,
			`PLATFORM_ID="platform:el9"`,
			`PRETTY_NAME="Oracle Linux Server 9.1"`,
			`ANSI_COLOR="0;31"`,
			`CPE_NAME="cpe:/o:oracle:linux:9:1:server"`,
			`HOME_URL="https://linux.oracle.com/"`,
			`BUG_REPORT_URL="https://github.com/oracle/oracle-linux"`,
		}, "\n"))

		assertKeywords(t, oraclelinux.Keywords, osRelease, map[string]string{
			"%version%": "9",
		})

	})

}
