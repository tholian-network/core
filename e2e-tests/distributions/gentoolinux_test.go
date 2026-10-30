package distributions_test

import (
	"strings"
	"testing"

	"tholian-core/distributions/gentoolinux"
)

func TestGentooLinuxKeywords(t *testing.T) {

	osRelease := []byte(strings.Join([]string{
		`NAME=Gentoo`,
		`ID=gentoo`,
		`PRETTY_NAME="Gentoo Linux"`,
		`ANSI_COLOR="1;32"`,
		`HOME_URL="https://www.gentoo.org/"`,
		`SUPPORT_URL="https://www.gentoo.org/support/"`,
		`BUG_REPORT_URL="https://bugs.gentoo.org/"`,
		`VERSION_ID="2.18"`,
	}, "\n"))

	assertKeywords(t, gentoolinux.Keywords, osRelease, nil)

}
