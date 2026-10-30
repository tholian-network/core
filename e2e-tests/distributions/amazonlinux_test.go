package distributions_test

import (
	"strings"
	"testing"

	"tholian-core/distributions/amazonlinux"
)

func TestAmazonLinuxKeywords(t *testing.T) {

	t.Run("2", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="Amazon Linux"`,
			`VERSION="2"`,
			`ID="amzn"`,
			`ID_LIKE="centos rhel fedora"`,
			`VERSION_ID="2"`,
			`PRETTY_NAME="Amazon Linux 2"`,
			`ANSI_COLOR="0;33"`,
			`CPE_NAME="cpe:2.3:o:amazon:amazon_linux:2"`,
			`HOME_URL="https://amazonlinux.com/"`,
			`SUPPORT_END="2026-06-30"`,
		}, "\n"))

		assertKeywords(t, amazonlinux.Keywords, osRelease, map[string]string{
			"%version%": "2",
		})

	})

	t.Run("2023", func(t *testing.T) {

		osRelease := []byte(strings.Join([]string{
			`NAME="Amazon Linux"`,
			`VERSION="2023"`,
			`ID="amzn"`,
			`ID_LIKE="fedora"`,
			`VERSION_ID="2023"`,
			`PLATFORM_ID="platform:al2023"`,
			`PRETTY_NAME="Amazon Linux 2023.8.20250915"`,
			`ANSI_COLOR="0;33"`,
			`CPE_NAME="cpe:2.3:o:amazon:amazon_linux:2023"`,
			`HOME_URL="https://aws.amazon.com/linux/amazon-linux-2023/"`,
			`SUPPORT_END="2029-06-30"`,
		}, "\n"))

		assertKeywords(t, amazonlinux.Keywords, osRelease, map[string]string{
			"%version%": "2023",
		})

	})

}
