package distributions_test

import (
	"strings"
	"testing"

	"tholian-core/distributions/fedora"
)

func TestFedoraKeywords(t *testing.T) {

	t.Run("38", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="Fedora Linux"`,
			`VERSION="38 (Workstation Edition)"`,
			`ID=fedora`,
			`VERSION_ID=38`,
			`VERSION_CODENAME=""`,
			`PLATFORM_ID="platform:f38"`,
			`PRETTY_NAME="Fedora Linux 38 (Workstation Edition)"`,
			`ANSI_COLOR="0;38;2;60;110;180"`,
			`LOGO=fedora-logo-icon`,
			`CPE_NAME="cpe:/o:fedoraproject:fedora:38"`,
			`DEFAULT_HOSTNAME="fedora"`,
			`HOME_URL="https://fedoraproject.org/"`,
			`VARIANT="Workstation Edition"`,
			`VARIANT_ID=workstation`,
		}, "\n"))

		assertKeywords(t, fedora.Keywords, osRelease, map[string]string{
			"%version%": "38",
		})

	})

	t.Run("41", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="Fedora Linux"`,
			`VERSION="41 (Container Image)"`,
			`RELEASE_TYPE=stable`,
			`ID=fedora`,
			`VERSION_ID=41`,
			`VERSION_CODENAME=""`,
			`PLATFORM_ID="platform:f41"`,
			`PRETTY_NAME="Fedora Linux 41 (Container Image)"`,
			`ANSI_COLOR="0;38;2;60;110;180"`,
			`LOGO=fedora-logo-icon`,
			`CPE_NAME="cpe:/o:fedoraproject:fedora:41"`,
			`DEFAULT_HOSTNAME="fedora"`,
			`VARIANT="Container Image"`,
			`VARIANT_ID=container`,
		}, "\n"))

		assertKeywords(t, fedora.Keywords, osRelease, map[string]string{
			"%version%": "41",
		})

	})

}
