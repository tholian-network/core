package distributions_test

import (
	"strings"
	"testing"

	"tholian-core/distributions/photonos"
)

func TestPhotonOSKeywords(t *testing.T) {

	t.Run("1.0", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="VMware Photon"`,
			`VERSION="1.0"`,
			`ID=photon`,
			`VERSION_ID=1.0`,
			`PRETTY_NAME="VMware Photon/Linux"`,
			`ANSI_COLOR="1;34"`,
			`HOME_URL="https://vmware.github.io/photon/"`,
			`BUG_REPORT_URL="https://github.com/vmware/photon/issues"`,
		}, "\n"))

		assertKeywords(t, photonos.KeywordsLegacy, osRelease, map[string]string{
			"%version%": "1.0",
		})

	})

	t.Run("3.0", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="VMware Photon OS"`,
			`VERSION="3.0"`,
			`ID=photon`,
			`VERSION_ID=3.0`,
			`PRETTY_NAME="VMware Photon OS/Linux"`,
			`ANSI_COLOR="1;34"`,
			`HOME_URL="https://vmware.github.io/photon/"`,
			`BUG_REPORT_URL="https://github.com/vmware/photon/issues"`,
		}, "\n"))

		assertKeywords(t, photonos.Keywords, osRelease, map[string]string{
			"%version%": "3.0",
		})

	})

}
