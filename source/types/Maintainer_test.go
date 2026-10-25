package types

import "testing"

func TestMaintainer(t *testing.T) {

	t.Run("IsMaintainer()", func(t *testing.T) {

		result1 := IsMaintainer("Harold Fynch")
		result2 := IsMaintainer("harold@fynch.io")
		result3 := IsMaintainer("Harold Fynch <harold@fynch.io>")
		result4 := IsMaintainer("<harold@fynch.io>")
		result5 := IsMaintainer("<harold@fynch.io> Harold Fynch")

		if result1 != true {
			t.Errorf("Expected %t to be true", result1)
		}

		if result2 != true {
			t.Errorf("Expected %t to be true", result2)
		}

		if result3 != true {
			t.Errorf("Expected %t to be true", result3)
		}

		if result3 != true {
			t.Errorf("Expected %t to be true", result4)
		}

		if result5 != false {
			t.Errorf("Expected %t to be false", result5)
		}

	})

	t.Run("ParseMaintainer()", func(t *testing.T) {

		maintainer1 := ParseMaintainer("Harold Fynch")
		maintainer2 := ParseMaintainer("harold@fynch.io")
		maintainer3 := ParseMaintainer("Harold Fynch <harold@fynch.io>")
		maintainer4 := ParseMaintainer("<harold@fynch.io>")
		maintainer5 := ParseMaintainer("<harold@fynch.io> Harold Fynch")

		if maintainer1 == nil {
			t.Errorf("Expected %s to be valid", "Harold Fynch")
		} else if maintainer1.String() != "Harold Fynch" {
			t.Errorf("Expected %s to be %s", maintainer1.String(), "Harold Fynch")
		}

		if maintainer2 == nil {
			t.Errorf("Expected %s to be valid", "harold@fynch.io")
		} else if maintainer2.String() != "harold@fynch.io" {
			t.Errorf("Expected %s to be %s", maintainer2.String(), "harold@fynch.io")
		}

		if maintainer3 == nil {
			t.Errorf("Expected %s to be valid", "Harold Fynch <harold@fynch.io>")
		} else if maintainer3.String() != "Harold Fynch <harold@fynch.io>" {
			t.Errorf("Expected %s to be %s", maintainer3.String(), "Harold Fynch <harold@fynch.io>")
		}

		if maintainer4 == nil {
			t.Errorf("Expected %s to be valid", "<harold@fynch.io>")
		} else if maintainer4.String() != "harold@fynch.io" {
			t.Errorf("Expected %s to be %s", maintainer4.String(), "harold@fynch.io")
		}

		if maintainer5 != nil {
			t.Errorf("Expected %s to be invalid", maintainer5.String())
		}

	})

}
