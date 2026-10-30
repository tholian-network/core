package distributions_test

import (
	"strings"
	"testing"

	"tholian-core/distributions/ubuntu"
)

func TestUbuntuKeywords(t *testing.T) {

	t.Run("14.04", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="Ubuntu"`,
			`VERSION="14.04.6 LTS, Trusty Tahr"`,
			`ID=ubuntu`,
			`ID_LIKE=debian`,
			`PRETTY_NAME="Ubuntu 14.04.6 LTS"`,
			`VERSION_ID="14.04"`,
			`HOME_URL="http://www.ubuntu.com/"`,
			`SUPPORT_URL="http://help.ubuntu.com/"`,
			`BUG_REPORT_URL="http://bugs.launchpad.net/ubuntu/"`,
		}, "\n"))

		assertKeywords(t, ubuntu.KeywordsLegacy, osRelease, map[string]string{
			"%version%": "14.04",
		})

	})

	t.Run("24.04", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`PRETTY_NAME="Ubuntu 24.04.3 LTS"`,
			`NAME="Ubuntu"`,
			`VERSION_ID="24.04"`,
			`VERSION="24.04.3 LTS (Noble Numbat)"`,
			`VERSION_CODENAME=noble`,
			`ID=ubuntu`,
			`ID_LIKE=debian`,
			`HOME_URL="https://www.ubuntu.com/"`,
			`SUPPORT_URL="https://help.ubuntu.com/"`,
			`BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"`,
			`PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"`,
			`UBUNTU_CODENAME=noble`,
			`LOGO=ubuntu-logo`,
		}, "\n"))

		assertKeywords(t, ubuntu.Keywords, osRelease, map[string]string{
			"%release%": "noble",
		})

	})

}
