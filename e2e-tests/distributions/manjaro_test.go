package distributions_test

import (
	"strings"
	"testing"

	"tholian-core/distributions/manjaro"
)

func TestManjaroKeywords(t *testing.T) {

	t.Run("manjaro", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="Manjaro Linux"`,
			`PRETTY_NAME="Manjaro Linux"`,
			`ID=manjaro`,
			`ID_LIKE=arch`,
			`BUILD_ID=rolling`,
			`ANSI_COLOR="32;1;24;144;200"`,
			`HOME_URL="https://manjaro.org/"`,
			`DOCUMENTATION_URL="https://wiki.manjaro.org/"`,
			`SUPPORT_URL="https://forum.manjaro.org/"`,
			`BUG_REPORT_URL="https://docs.manjaro.org/reporting-bugs/"`,
			`PRIVACY_POLICY_URL="https://manjaro.org/privacy-policy/"`,
			`LOGO=manjarolinux`,
		}, "\n"))

		assertKeywords(t, manjaro.Keywords, osRelease, nil)

	})

	t.Run("manjaro-arm", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="Manjaro ARM"`,
			`ID="manjaro-arm"`,
			`ID_LIKE="manjaro arch"`,
			`PRETTY_NAME="Manjaro ARM"`,
			`ANSI_COLOR="1;32"`,
			`HOME_URL="https://www.manjaro.org/"`,
			`SUPPORT_URL="https://forum.manjaro.org/c/arm/"`,
			`LOGO=manjarolinux`,
		}, "\n"))

		assertKeywords(t, manjaro.KeywordsARM, osRelease, nil)

	})

}
