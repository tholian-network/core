package distributions_test

import (
	"strings"
	"testing"

	"tholian-core/distributions/linuxmint"
)

func TestLinuxMintKeywords(t *testing.T) {

	t.Run("21.3", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="Linux Mint"`,
			`VERSION="21.3 (Virginia)"`,
			`ID=linuxmint`,
			`ID_LIKE="ubuntu debian"`,
			`PRETTY_NAME="Linux Mint 21.3"`,
			`VERSION_ID="21.3"`,
			`HOME_URL="https://www.linuxmint.com/"`,
			`SUPPORT_URL="https://forums.linuxmint.com/"`,
			`BUG_REPORT_URL="http://linuxmint-troubleshooting-guide.readthedocs.io/en/latest/"`,
			`PRIVACY_POLICY_URL="https://www.linuxmint.com/"`,
			`VERSION_CODENAME=virginia`,
			`UBUNTU_CODENAME=jammy`,
		}, "\n"))

		assertKeywords(t, linuxmint.Keywords, osRelease, map[string]string{
			"%release%": "virginia",
		})

	})

	t.Run("22", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="Linux Mint"`,
			`VERSION="22 (Wilma)"`,
			`ID=linuxmint`,
			`ID_LIKE="ubuntu debian"`,
			`PRETTY_NAME="Linux Mint 22"`,
			`VERSION_ID="22"`,
			`HOME_URL="https://www.linuxmint.com/"`,
			`SUPPORT_URL="https://forums.linuxmint.com/"`,
			`BUG_REPORT_URL="http://linuxmint-troubleshooting-guide.readthedocs.io/en/latest/"`,
			`PRIVACY_POLICY_URL="https://www.linuxmint.com/"`,
			`VERSION_CODENAME=wilma`,
			`UBUNTU_CODENAME=noble`,
		}, "\n"))

		assertKeywords(t, linuxmint.Keywords, osRelease, map[string]string{
			"%release%": "wilma",
		})

	})

}
