package distributions_test

import (
	"strings"
	"testing"

	"tholian-core/distributions/suse"
)

func TestSUSEKeywords(t *testing.T) {

	t.Run("SLED 12", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="SLED"`,
			`VERSION="12-SP3"`,
			`VERSION_ID="12.3"`,
			`PRETTY_NAME="SUSE Linux Enterprise Desktop 12 SP3"`,
			`ID="sled"`,
			`ANSI_COLOR="0;32"`,
			`CPE_NAME="cpe:/o:suse:sled:12:sp3"`,
		}, "\n"))

		assertKeywords(t, suse.KeywordsDesktop, osRelease, map[string]string{
			"%version%": "12",
		})

	})

	t.Run("SLES 15", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="SLES"`,
			`VERSION="15-SP7"`,
			`VERSION_ID="15.7"`,
			`PRETTY_NAME="SUSE Linux Enterprise Server 15 SP7"`,
			`ID="sles"`,
			`ID_LIKE="suse"`,
			`ANSI_COLOR="0;32"`,
			`CPE_NAME="cpe:/o:suse:sles:15:sp7"`,
			`DOCUMENTATION_URL="https://documentation.suse.com/"`,
		}, "\n"))

		assertKeywords(t, suse.KeywordsServer, osRelease, map[string]string{
			"%version%": "15",
		})

	})

}
