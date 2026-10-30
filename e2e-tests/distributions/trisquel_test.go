package distributions_test

import (
	"strings"
	"testing"

	"tholian-core/distributions/trisquel"
)

func TestTrisquelKeywords(t *testing.T) {

	osRelease := []byte(strings.Join([]string{
		`NAME="Trisquel GNU/Linux"`,
		`VERSION="11.0 (Aramo)"`,
		`ID=trisquel`,
		`ID_LIKE="ubuntu debian"`,
		`VERSION_ID="11.0"`,
		`PRETTY_NAME="Trisquel GNU/Linux 11.0 (Aramo)"`,
		`VERSION_CODENAME=aramo`,
		`UBUNTU_CODENAME=jammy`,
		`HOME_URL="https://trisquel.info/"`,
		`SUPPORT_URL="https://trisquel.info/en/forum"`,
		`BUG_REPORT_URL="https://trisquel.info/en/issues"`,
	}, "\n"))

	assertKeywords(t, trisquel.Keywords, osRelease, map[string]string{
		"%release%":         "aramo",
		"%ubuntu_codename%": "jammy",
	})

}
