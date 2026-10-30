package distributions_test

import (
	"strings"
	"testing"

	"tholian-core/distributions/debian"
)

func TestDebianKeywords(t *testing.T) {

	t.Run("jessie", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`PRETTY_NAME="Debian GNU/Linux 8 (jessie)"`,
			`NAME="Debian GNU/Linux"`,
			`VERSION_ID="8"`,
			`VERSION="8 (jessie)"`,
			`ID=debian`,
			`HOME_URL="http://www.debian.org/"`,
			`SUPPORT_URL="http://www.debian.org/support"`,
			`BUG_REPORT_URL="https://bugs.debian.org/"`,
		}, "\n"))

		assertKeywords(t, debian.KeywordsLegacy, osRelease, map[string]string{
			"%version%": "8",
		})

	})

	t.Run("bookworm", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`PRETTY_NAME="Debian GNU/Linux 12 (bookworm)"`,
			`NAME="Debian GNU/Linux"`,
			`VERSION_ID="12"`,
			`VERSION="12 (bookworm)"`,
			`VERSION_CODENAME=bookworm`,
			`ID=debian`,
			`HOME_URL="https://www.debian.org/"`,
			`SUPPORT_URL="https://www.debian.org/support"`,
			`BUG_REPORT_URL="https://bugs.debian.org/"`,
		}, "\n"))

		assertKeywords(t, debian.Keywords, osRelease, map[string]string{
			"%release%": "bookworm",
		})

	})

}
