package distributions_test

import (
	"strings"
	"testing"

	"tholian-core/distributions/opensuse"
)

func TestOpenSUSEKeywords(t *testing.T) {

	t.Run("13.2", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME=openSUSE`,
			`VERSION="13.2 (Harlequin)"`,
			`VERSION_ID="13.2"`,
			`PRETTY_NAME="openSUSE 13.2 (Harlequin) (x86_64)"`,
			`ID=opensuse`,
			`ANSI_COLOR="0;32"`,
			`CPE_NAME="cpe:/o:opensuse:opensuse:13.2"`,
			`BUG_REPORT_URL="https://bugs.opensuse.org"`,
			`HOME_URL="https://opensuse.org/"`,
			`ID_LIKE="suse"`,
		}, "\n"))

		assertKeywords(t, opensuse.Keywords, osRelease, map[string]string{
			"%version%": "13.2",
		})

	})

	t.Run("leap-42.3", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="openSUSE Leap"`,
			`VERSION="42.3"`,
			`ID=opensuse`,
			`ID_LIKE="suse"`,
			`VERSION_ID="42.3"`,
			`PRETTY_NAME="openSUSE Leap 42.3"`,
			`ANSI_COLOR="0;32"`,
			`CPE_NAME="cpe:/o:opensuse:leap:42.3"`,
			`BUG_REPORT_URL="https://bugs.opensuse.org"`,
			`HOME_URL="https://www.opensuse.org/"`,
		}, "\n"))

		assertKeywords(t, opensuse.KeywordsLeapLegacy, osRelease, map[string]string{
			"%version%": "42.3",
		})

	})

	t.Run("leap-15.6", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="openSUSE Leap"`,
			`VERSION="15.6"`,
			`ID="opensuse-leap"`,
			`ID_LIKE="suse opensuse"`,
			`VERSION_ID="15.6"`,
			`PRETTY_NAME="openSUSE Leap 15.6"`,
			`ANSI_COLOR="0;32"`,
			`CPE_NAME="cpe:/o:opensuse:leap:15.6"`,
			`BUG_REPORT_URL="https://bugs.opensuse.org"`,
			`HOME_URL="https://www.opensuse.org/"`,
			`DOCUMENTATION_URL="https://en.opensuse.org/Portal:Leap"`,
			`LOGO="distributor-logo-Leap"`,
		}, "\n"))

		assertKeywords(t, opensuse.KeywordsLeap, osRelease, map[string]string{
			"%version%": "15",
		})

	})

	t.Run("tumbleweed", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="openSUSE Tumbleweed"`,
			`# VERSION="20240823"`,
			`ID="opensuse-tumbleweed"`,
			`ID_LIKE="opensuse suse"`,
			`VERSION_ID="20240823"`,
			`PRETTY_NAME="openSUSE Tumbleweed"`,
			`ANSI_COLOR="0;32"`,
			`# CPE 2.3 format, boo#1217921`,
			`CPE_NAME="cpe:2.3:o:opensuse:tumbleweed:20240823:*:*:*:*:*:*:*"`,
			`#CPE 2.2 format`,
			`#CPE_NAME="cpe:/o:opensuse:tumbleweed:20240823"`,
			`BUG_REPORT_URL="https://bugzilla.opensuse.org"`,
			`SUPPORT_URL="https://bugs.opensuse.org"`,
			`HOME_URL="https://www.opensuse.org"`,
			`DOCUMENTATION_URL="https://en.opensuse.org/Portal:Tumbleweed"`,
			`LOGO="distributor-logo-Tumbleweed"`,
		}, "\n"))

		assertKeywords(t, opensuse.KeywordsTumbleweed, osRelease, nil)

	})

}
