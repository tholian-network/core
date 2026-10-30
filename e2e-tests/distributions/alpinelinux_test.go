package distributions_test

import (
	"strings"
	"testing"

	"tholian-core/distributions/alpinelinux"
)

func TestAlpineLinuxKeywords(t *testing.T) {

	t.Run("3.10", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="Alpine Linux"`,
			`ID=alpine`,
			`VERSION_ID=3.10.9`,
			`PRETTY_NAME="Alpine Linux v3.10"`,
			`HOME_URL="https://alpinelinux.org/"`,
			`BUG_REPORT_URL="https://bugs.alpinelinux.org/"`,
		}, "\n"))

		assertKeywords(t, alpinelinux.Keywords, osRelease, map[string]string{
			"%release%": "3.10",
		})

	})

	t.Run("3.18", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="Alpine Linux"`,
			`ID=alpine`,
			`VERSION_ID=3.18.12`,
			`PRETTY_NAME="Alpine Linux v3.18"`,
			`HOME_URL="https://alpinelinux.org/"`,
			`BUG_REPORT_URL="https://gitlab.alpinelinux.org/alpine/aports/-/issues"`,
		}, "\n"))

		assertKeywords(t, alpinelinux.Keywords, osRelease, map[string]string{
			"%release%": "3.18",
		})

	})

}
