package distributions_test

import (
	"strings"
	"testing"

	"tholian-core/distributions/rockylinux"
)

func TestRockyLinuxKeywords(t *testing.T) {

	t.Run("8", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="Rocky Linux"`,
			`VERSION="8.9 (Green Obsidian)"`,
			`ID="rocky"`,
			`ID_LIKE="rhel centos fedora"`,
			`VERSION_ID="8.9"`,
			`PLATFORM_ID="platform:el8"`,
			`PRETTY_NAME="Rocky Linux 8.9 (Green Obsidian)"`,
			`ANSI_COLOR="0;32"`,
			`LOGO="fedora-logo-icon"`,
			`CPE_NAME="cpe:/o:rocky:rocky:8:GA"`,
			`HOME_URL="https://rockylinux.org/"`,
			`BUG_REPORT_URL="https://bugs.rockylinux.org/"`,
		}, "\n"))

		assertKeywords(t, rockylinux.Keywords, osRelease, map[string]string{
			"%version%": "8",
		})

	})

	t.Run("9", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="Rocky Linux"`,
			`VERSION="9.3 (Blue Onyx)"`,
			`ID="rocky"`,
			`ID_LIKE="rhel centos fedora"`,
			`VERSION_ID="9.3"`,
			`PLATFORM_ID="platform:el9"`,
			`PRETTY_NAME="Rocky Linux 9.3 (Blue Onyx)"`,
			`ANSI_COLOR="0;32"`,
			`LOGO="fedora-logo-icon"`,
			`CPE_NAME="cpe:/o:rocky:rocky:9::baseos"`,
			`HOME_URL="https://rockylinux.org/"`,
			`BUG_REPORT_URL="https://bugs.rockylinux.org/"`,
		}, "\n"))

		assertKeywords(t, rockylinux.Keywords, osRelease, map[string]string{
			"%version%": "9",
		})

	})

}
