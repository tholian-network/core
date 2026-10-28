package types

import (
	"encoding/json"
	"testing"
)

func TestVersion(t *testing.T) {

	t.Run("ToVersion(alpinelinux format)", func(t *testing.T) {

		v_1_23_0 := ToVersion("1.23.0")
		v_1_23_1 := ToVersion("1.23.1")
		v_1_23_44 := ToVersion("1.23.44")
		v_1_23_45 := ToVersion("1.23.45")
		v_1_23_46 := ToVersion("1.23.46")
		v_1_23_45_r5 := ToVersion("1.23.45-r5")
		v_1_23_45_r6 := ToVersion("1.23.45-r6")
		v_1_23_45_r7 := ToVersion("1.23.45-r7")

		version1 := ToVersion("1.23.45-r6")

		if version1.IsAfter(v_1_23_44) != true {
			t.Errorf("Expected %s to be after %s", version1.String(), v_1_23_44.String())
		}

		if version1.IsAfter(v_1_23_45) != true {
			t.Errorf("Expected %s to be after %s", version1.String(), v_1_23_45.String())
		}

		if version1.IsBefore(v_1_23_46) != true {
			t.Errorf("Expected %s to be before %s", version1.String(), v_1_23_46.String())
		}

		if version1.IsAfter(v_1_23_45_r5) != true {
			t.Errorf("Expected %s to be after %s", version1.String(), v_1_23_45_r5.String())
		}

		if version1.Equal(v_1_23_45_r6) != true {
			t.Errorf("Expected %s to be the same as %s", version1.String(), v_1_23_45_r6.String())
		}

		if version1.IsBefore(v_1_23_45_r7) != true {
			t.Errorf("Expected %s to be before %s", version1.String(), v_1_23_45_r7.String())
		}

		version2 := ToVersion("1.23_pre45-r6")

		if version2.IsAfter(v_1_23_0) != true {
			t.Errorf("Expected %s to be after %s", version2.String(), v_1_23_0.String())
		}

		if version2.IsBefore(v_1_23_1) != true {
			t.Errorf("Expected %s to be before %s", version2.String(), v_1_23_1.String())
		}

		if version2.IsBefore(v_1_23_44) != true {
			t.Errorf("Expected %s to be before %s", version2.String(), v_1_23_44.String())
		}

		if version2.IsBefore(v_1_23_45) != true {
			t.Errorf("Expected %s to be before %s", version2.String(), v_1_23_45.String())
		}

		if version2.IsBefore(v_1_23_45_r5) != true {
			t.Errorf("Expected %s to be before %s", version2.String(), v_1_23_45_r5.String())
		}

		if version2.IsBefore(v_1_23_45_r6) != true {
			t.Errorf("Expected %s to be before %s", version2.String(), v_1_23_45_r6.String())
		}

		if version2.IsBefore(v_1_23_45_r7) != true {
			t.Errorf("Expected %s to be before %s", version2.String(), v_1_23_45_r7.String())
		}

	})

	t.Run("ToVersion(archlinux format)", func(t *testing.T) {

		v_1_23_44 := ToVersion("1.23.44")
		v_1_23_45 := ToVersion("1.23.45")
		v_1_23_46 := ToVersion("1.23.46")
		v_1_23_45_r5 := ToVersion("1.23.45-r5")
		v_1_23_45_r6 := ToVersion("1.23.45-r6")
		v_1_23_45_r7 := ToVersion("1.23.45-r7")

		version1 := ToVersion("1.23.45r6+g9b7d253-1")

		if version1.IsAfter(v_1_23_44) != true {
			t.Errorf("Expected %s to be after %s", version1.String(), v_1_23_44.String())
		}

		if version1.IsAfter(v_1_23_45) != true {
			t.Errorf("Expected %s to be after %s", version1.String(), v_1_23_45.String())
		}

		if version1.IsBefore(v_1_23_46) != true {
			t.Errorf("Expected %s to be before %s", version1.String(), v_1_23_46.String())
		}

		if version1.IsAfter(v_1_23_45_r5) != true {
			t.Errorf("Expected %s to be after %s", version1.String(), v_1_23_45_r5.String())
		}

		if version1.Equal(v_1_23_45_r6) != true {
			t.Errorf("Expected %s to be the same as %s", version1.String(), v_1_23_45_r6.String())
		}

		if version1.IsBefore(v_1_23_45_r7) != true {
			t.Errorf("Expected %s to be before %s", version1.String(), v_1_23_45_r7.String())
		}

		if version1.IsBefore(v_1_23_46) != true {
			t.Errorf("Expected %s to be before %s", version1.String(), v_1_23_46.String())
		}

		version2 := ToVersion("1.23.45+kde+r6-1")

		if version2.IsAfter(v_1_23_44) != true {
			t.Errorf("Expected %s to be after %s", version2.String(), v_1_23_44.String())
		}

		if version2.IsAfter(v_1_23_45) != true {
			t.Errorf("Expected %s to be after %s", version2.String(), v_1_23_45.String())
		}

		if version2.IsAfter(v_1_23_45_r5) != true {
			t.Errorf("Expected %s to be after %s", version2.String(), v_1_23_45_r5.String())
		}

		if version2.Equal(v_1_23_45_r6) != true {
			t.Errorf("Expected %s to be the same as %s", version2.String(), v_1_23_45_r6.String())
		}

		if version2.IsBefore(v_1_23_45_r7) != true {
			t.Errorf("Expected %s to be before %s", version2.String(), v_1_23_45_r7.String())
		}

		if version2.IsBefore(v_1_23_46) != true {
			t.Errorf("Expected %s to be before %s", version2.String(), v_1_23_46.String())
		}

		version3 := ToVersion("1.23.45.r6-1")

		if version3.IsAfter(v_1_23_44) != true {
			t.Errorf("Expected %s to be after %s", version3.String(), v_1_23_44.String())
		}

		if version3.IsAfter(v_1_23_45) != true {
			t.Errorf("Expected %s to be after %s", version3.String(), v_1_23_45.String())
		}

		if version3.IsAfter(v_1_23_45_r5) != true {
			t.Errorf("Expected %s to be after %s", version3.String(), v_1_23_45_r5.String())
		}

		if version3.Equal(v_1_23_45_r6) != true {
			t.Errorf("Expected %s to be the same as %s", version3.String(), v_1_23_45_r6.String())
		}

		if version3.IsBefore(v_1_23_45_r7) != true {
			t.Errorf("Expected %s to be before %s", version3.String(), v_1_23_45_r7.String())
		}

		if version3.IsBefore(v_1_23_46) != true {
			t.Errorf("Expected %s to be before %s", version3.String(), v_1_23_46.String())
		}

		version4 := ToVersion("2:1.23.45-r6-1")
		v2_1_23_44 := ToVersion("2:1.23.44")
		v2_1_23_45 := ToVersion("2:1.23.45")
		v2_1_23_46 := ToVersion("2:1.23.46")
		v2_1_23_45_r5 := ToVersion("2:1.23.45-r5")
		v2_1_23_45_r6 := ToVersion("2:1.23.45-r6")
		v2_1_23_45_r7 := ToVersion("2:1.23.45-r7")

		if version4.IsAfter(v_1_23_44) != true {
			t.Errorf("Expected %s to be after %s", version4.String(), v_1_23_44.String())
		}

		if version4.IsAfter(v_1_23_45) != true {
			t.Errorf("Expected %s to be after %s", version4.String(), v_1_23_45.String())
		}

		if version4.IsAfter(v_1_23_45_r5) != true {
			t.Errorf("Expected %s to be after %s", version4.String(), v_1_23_45_r5.String())
		}

		if version4.IsAfter(v_1_23_45_r6) != true {
			t.Errorf("Expected %s to be the same as %s", version4.String(), v_1_23_45_r6.String())
		}

		if version4.IsAfter(v_1_23_45_r7) != true {
			t.Errorf("Expected %s to be before %s", version4.String(), v_1_23_45_r7.String())
		}

		if version4.IsAfter(v_1_23_46) != true {
			t.Errorf("Expected %s to be before %s", version4.String(), v_1_23_46.String())
		}

		if version4.IsAfter(v2_1_23_44) != true {
			t.Errorf("Expected %s to be after %s", version4.String(), v2_1_23_44.String())
		}

		if version4.IsAfter(v2_1_23_45) != true {
			t.Errorf("Expected %s to be after %s", version4.String(), v2_1_23_45.String())
		}

		if version4.IsAfter(v2_1_23_45_r5) != true {
			t.Errorf("Expected %s to be after %s", version4.String(), v2_1_23_45_r5.String())
		}

		if version4.Equal(v2_1_23_45_r6) != true {
			t.Errorf("Expected %s to be the same as %s", version4.String(), v2_1_23_45_r6.String())
		}

		if version4.IsBefore(v2_1_23_45_r7) != true {
			t.Errorf("Expected %s to be before %s", version4.String(), v2_1_23_45_r7.String())
		}

		if version4.IsBefore(v2_1_23_46) != true {
			t.Errorf("Expected %s to be before %s", version4.String(), v2_1_23_46.String())
		}

	})

	t.Run("ToVersion(debian format)", func(t *testing.T) {

		v_1_23_44 := ToVersion("1.23.44")
		v_1_23_45 := ToVersion("1.23.45")
		v_1_23_46 := ToVersion("1.23.46")
		v_1_23_45_r5 := ToVersion("1.23.45~r5+git+abc123d-8")
		v_1_23_45_r6 := ToVersion("1.23.45~r6+git+fff999f-8")
		v_1_23_45_r7 := ToVersion("1.23.45~r7+git+ac3201d-8")

		v_1_23_45_deb5 := ToVersion("1.23.45+deb5u7")
		v_1_23_45_deb6 := ToVersion("1.23.45+deb6u6")
		v_1_23_45_deb7 := ToVersion("1.23.45+deb7u5")

		version1 := ToVersion("1.23.45~r6")

		if version1.IsAfter(v_1_23_44) != true {
			t.Errorf("Expected %s to be after %s", version1.String(), v_1_23_44.String())
		}

		if version1.IsAfter(v_1_23_45) != true {
			t.Errorf("Expected %s to be after %s", version1.String(), v_1_23_45.String())
		}

		if version1.IsBefore(v_1_23_46) != true {
			t.Errorf("Expected %s to be before %s", version1.String(), v_1_23_46.String())
		}

		if version1.IsAfter(v_1_23_45_r5) != true {
			t.Errorf("Expected %s to be after %s", version1.String(), v_1_23_45_r5.String())
		}

		if version1.Equal(v_1_23_45_r6) != true {
			t.Errorf("Expected %s to be the same as %s", version1.String(), v_1_23_45_r6.String())
		}

		if version1.IsBefore(v_1_23_45_r7) != true {
			t.Errorf("Expected %s to be before %s", version1.String(), v_1_23_45_r7.String())
		}

		version2 := ToVersion("1.23.45+deb6u1")

		if version2.IsAfter(v_1_23_45_deb5) != true {
			t.Errorf("Expected %s to be after %s", version2.String(), v_1_23_45_deb5.String())
		}

		if version2.IsBefore(v_1_23_45_deb6) != true {
			t.Errorf("Expected %s to be before %s", version2.String(), v_1_23_45_deb6.String())
		}

		if version2.IsBefore(v_1_23_45_deb7) != true {
			t.Errorf("Expected %s to be before %s", version2.String(), v_1_23_45_deb7.String())
		}

		version3 := ToVersion("1.23.45+deb6u6")

		if version3.IsAfter(v_1_23_45_deb5) != true {
			t.Errorf("Expected %s to be after %s", version3.String(), v_1_23_45_deb5.String())
		}

		if version3.Equal(v_1_23_45_deb6) != true {
			t.Errorf("Expected %s to be the same as %s", version3.String(), v_1_23_45_deb6.String())
		}

		if version3.IsBefore(v_1_23_45_deb7) != true {
			t.Errorf("Expected %s to be before %s", version3.String(), v_1_23_45_deb7.String())
		}

		version4 := ToVersion("0:1.23.45+dfsg-123~deb6u6")

		if version4.IsAfter(v_1_23_45_deb5) != true {
			t.Errorf("Expected %s to be after %s", version4.String(), v_1_23_45_deb5.String())
		}

		if version4.Equal(v_1_23_45_deb6) != true {
			t.Errorf("Expected %s to be the same as %s", version4.String(), v_1_23_45_deb6.String())
		}

		// XXX: This is the correct [epoche:][upstream_release][debian_release] behaviour
		// XXX: because ...+dfsg-123... is part of the upstream_release version part
		if version4.IsAfter(v_1_23_45_deb7) != true {
			t.Errorf("Expected %s to be after %s", version4.String(), v_1_23_45_deb7.String())
		}

	})

	t.Run("ToVersion(ubuntu format)", func(t *testing.T) {

		v_1_23_44 := ToVersion("1.23.44")
		v_1_23_45 := ToVersion("1.23.45")
		v_1_23_46 := ToVersion("1.23.46")

		v_1_23_45_1ubuntu65 := ToVersion("0:1.23.45-1ubuntu6.5")
		v_1_23_45_1ubuntu66 := ToVersion("0:1.23.45-1ubuntu6.6")
		v_1_23_45_1ubuntu67 := ToVersion("0:1.23.45-1ubuntu6.7")

		v_2023_12_31_1ubuntu6_20_04_03 := ToVersion("2023.12.31-1ubuntu6~20.04.3")
		v_2023_12_31_1ubuntu6_21_10_02 := ToVersion("2023.12.31-1ubuntu6~21.10.2")
		v_2023_12_31_1ubuntu6_22_04_01 := ToVersion("2023.12.31-1ubuntu6~22.04.1")

		version1 := ToVersion("0:1.23.45-1ubuntu6.6")

		if version1.IsAfter(v_1_23_44) != true {
			t.Errorf("Expected %s to be after %s", version1.String(), v_1_23_44.String())
		}

		if version1.IsAfter(v_1_23_45) != true {
			t.Errorf("Expected %s to be after %s", version1.String(), v_1_23_45.String())
		}

		if version1.IsBefore(v_1_23_46) != true {
			t.Errorf("Expected %s to be before %s", version1.String(), v_1_23_46.String())
		}

		if version1.IsAfter(v_1_23_45_1ubuntu65) != true {
			t.Errorf("Expected %s to be after %s", version1.String(), v_1_23_45_1ubuntu65.String())
		}

		if version1.Equal(v_1_23_45_1ubuntu66) != true {
			t.Errorf("Expected %s to be the same as %s", version1.String(), v_1_23_45_1ubuntu66.String())
		}

		if version1.IsBefore(v_1_23_45_1ubuntu67) != true {
			t.Errorf("Expected %s to be before %s", version1.String(), v_1_23_45_1ubuntu67.String())
		}

		version2 := ToVersion("2023.12.31-1ubuntu6~21.10.2")

		if version2.IsAfter(v_2023_12_31_1ubuntu6_20_04_03) != true {
			t.Errorf("Expected %s to be after %s", version2.String(), v_2023_12_31_1ubuntu6_20_04_03.String())
		}

		if version2.Equal(v_2023_12_31_1ubuntu6_21_10_02) != true {
			t.Errorf("Expected %s to be the same as %s", version2.String(), v_2023_12_31_1ubuntu6_21_10_02.String())
		}

		if version2.IsBefore(v_2023_12_31_1ubuntu6_22_04_01) != true {
			t.Errorf("Expected %s to be before %s", version2.String(), v_2023_12_31_1ubuntu6_22_04_01.String())
		}

		version3 := ToVersion("2023.12.31-1ubuntu6")

		if version3.Equal(v_2023_12_31_1ubuntu6_20_04_03) != true {
			t.Errorf("Expected %s to be the same as %s", version3.String(), v_2023_12_31_1ubuntu6_20_04_03.String())
		}

		if version3.Equal(v_2023_12_31_1ubuntu6_21_10_02) != true {
			t.Errorf("Expected %s to be the same as %s", version3.String(), v_2023_12_31_1ubuntu6_21_10_02.String())
		}

		if version3.Equal(v_2023_12_31_1ubuntu6_22_04_01) != true {
			t.Errorf("Expected %s to be the same as %s", version3.String(), v_2023_12_31_1ubuntu6_22_04_01.String())
		}

	})

	t.Run("ToVersion(redhat format)", func(t *testing.T) {

		v_1_23_44 := ToVersion("1.23.44.el9")
		v_1_23_45 := ToVersion("1.23.45.el8")
		v_1_23_46 := ToVersion("1.23.46.el7")

		v_1_23_45_r6 := ToVersion("1.23.45-r6.el9")
		v_1_23_45_r7 := ToVersion("1.23.45-r7.el9")
		v_1_23_45_r8 := ToVersion("1.23.46-r8.el9")

		v0_1_23_45 := ToVersion("0:1.23.45")
		v1_1_23_45 := ToVersion("1:1.23.45")
		v2_1_23_45 := ToVersion("2:1.23.45")

		version1 := ToVersion("1.23.45-r7.el9")

		if version1.IsAfter(v_1_23_44) != true {
			t.Errorf("Expected %s to be after %s", version1.String(), v_1_23_44.String())
		}

		if version1.IsAfter(v_1_23_45) != true {
			t.Errorf("Expected %s to be after %s", version1.String(), v_1_23_45.String())
		}

		if version1.IsBefore(v_1_23_46) != true {
			t.Errorf("Expected %s to be before %s", version1.String(), v_1_23_46.String())
		}

		if version1.IsAfter(v_1_23_45_r6) != true {
			t.Errorf("Expected %s to be after %s", version1.String(), v_1_23_45_r6.String())
		}

		if version1.Equal(v_1_23_45_r7) != true {
			t.Errorf("Expected %s to be the same as %s", version1.String(), v_1_23_45_r7.String())
		}

		if version1.IsBefore(v_1_23_45_r8) != true {
			t.Errorf("Expected %s to be before %s", version1.String(), v_1_23_45_r8.String())
		}

		version2 := ToVersion("0:1.23.45.6-r7.el9_1")

		if version2.IsAfter(v_1_23_44) != true {
			t.Errorf("Expected %s to be after %s", version2.String(), v_1_23_44.String())
		}

		if version2.IsAfter(v_1_23_45) != true {
			t.Errorf("Expected %s to be after %s", version2.String(), v_1_23_45.String())
		}

		if version2.IsBefore(v_1_23_46) != true {
			t.Errorf("Expected %s to be before %s", version2.String(), v_1_23_46.String())
		}

		version3 := ToVersion("1:1.23.45")

		if version3.IsAfter(v0_1_23_45) != true {
			t.Errorf("Expected %s to be after %s", version3.String(), v0_1_23_45.String())
		}

		if version3.Equal(v1_1_23_45) != true {
			t.Errorf("Expected %s to be the same as %s", version3.String(), v1_1_23_45.String())
		}

		if version3.IsBefore(v2_1_23_45) != true {
			t.Errorf("Expected %s to be before %s", version3.String(), v2_1_23_45.String())
		}

		version4 := ToVersion("1.23.45-r7-8.git2a8f9d8.el9")

		if version4.IsAfter(v_1_23_45_r6) != true {
			t.Errorf("Expected %s to be after %s", version4.String(), v_1_23_45_r6.String())
		}

		if version4.Equal(v_1_23_45_r7) != true {
			t.Errorf("Expected %s to be after %s", version4.String(), v_1_23_45_r7.String())
		}

		if version4.IsBefore(v_1_23_45_r8) != true {
			t.Errorf("Expected %s to be before %s", version4.String(), v_1_23_45_r8.String())
		}

	})

	t.Run("ToVersion(opensuse format)", func(t *testing.T) {

		v_1_23_44 := ToVersion("1.23.44")
		v_1_23_45 := ToVersion("1.23.45")
		v_1_23_46 := ToVersion("1.23.46")

		v_1_23_45_6 := ToVersion("1.23.45-6")
		v_1_23_45_7 := ToVersion("1.23.45-7")
		v_1_23_45_8 := ToVersion("1.23.46-8")

		version1 := ToVersion("1.23.45-7.5")

		if version1.IsAfter(v_1_23_44) != true {
			t.Errorf("Expected %s to be after %s", version1.String(), v_1_23_44.String())
		}

		if version1.IsAfter(v_1_23_45) != true {
			t.Errorf("Expected %s to be after %s", version1.String(), v_1_23_45.String())
		}

		if version1.IsBefore(v_1_23_46) != true {
			t.Errorf("Expected %s to be before %s", version1.String(), v_1_23_46.String())
		}

		if version1.IsAfter(v_1_23_45_6) != true {
			t.Errorf("Expected %s to be after %s", version1.String(), v_1_23_45_6.String())
		}

		if version1.Equal(v_1_23_45_7) != true {
			t.Errorf("Expected %s to be the same as %s", version1.String(), v_1_23_45_7.String())
		}

		if version1.IsBefore(v_1_23_45_8) != true {
			t.Errorf("Expected %s to be before %s", version1.String(), v_1_23_45_8.String())
		}

		version2 := ToVersion("1.23.45.7_g31486f40-8.9")

		if version2.IsAfter(v_1_23_44) != true {
			t.Errorf("Expected %s to be after %s", version2.String(), v_1_23_44.String())
		}

		if version2.Equal(v_1_23_45) != true {
			t.Errorf("Expected %s to be the same as %s", version2.String(), v_1_23_45.String())
		}
		if version2.IsBefore(v_1_23_46) != true {
			t.Errorf("Expected %s to be before %s", version2.String(), v_1_23_46.String())
		}

		if version2.IsAfter(v_1_23_45_6) != true {
			t.Errorf("Expected %s to be after %s", version2.String(), v_1_23_45_6.String())
		}

		if version2.Equal(v_1_23_45_7) != true {
			t.Errorf("Expected %s to be the same as %s", version2.String(), v_1_23_45_7.String())
		}

		if version2.IsBefore(v_1_23_45_8) != true {
			t.Errorf("Expected %s to be before %s", version2.String(), v_1_23_45_8.String())
		}

		version3 := ToVersion("1.23.45.20240305.897f2593b3-1.1")

		if version3.IsAfter(v_1_23_44) != true {
			t.Errorf("Expected %s to be after %s", version3.String(), v_1_23_44.String())
		}

		if version3.Equal(v_1_23_45) != true {
			t.Errorf("Expected %s to be the same as %s", version3.String(), v_1_23_45.String())
		}

		if version3.IsBefore(v_1_23_46) != true {
			t.Errorf("Expected %s to be before %s", version3.String(), v_1_23_46.String())
		}

	})

	t.Run("ToVersion(package manager formats)", func(t *testing.T) {

		freebsd := ToVersion("2.4.62_1")

		if freebsd.Upstream.Major != 2 || freebsd.Upstream.Minor != 4 || freebsd.Upstream.Patch != 62 {
			t.Errorf("Expected 2.4.62 from %q, got %s", "2.4.62_1", freebsd.String())
		}

		if freebsd.Revision != "1" {
			t.Errorf("Expected revision 1 from %q, got %q", "2.4.62_1", freebsd.Revision)
		}

		if freebsd.IsValid() != true {
			t.Errorf("Expected %q to be valid", "2.4.62_1")
		}

		openwrt := ToVersion("2025.07.06~7e64e8ba-r5")

		if openwrt.Upstream.Major != 2025 || openwrt.Upstream.Minor != 7 || openwrt.Upstream.Patch != 6 {
			t.Errorf("Expected 2025.7.6 from %q, got %s", "2025.07.06~7e64e8ba-r5", openwrt.String())
		}

		if openwrt.Upstream.Hash != "7e64e8ba" {
			t.Errorf("Expected hash 7e64e8ba from %q, got %q", "2025.07.06~7e64e8ba-r5", openwrt.Upstream.Hash)
		}

		if openwrt.Upstream.Release != "r5" {
			t.Errorf("Expected release r5 from %q, got %q", "2025.07.06~7e64e8ba-r5", openwrt.Upstream.Release)
		}

		if openwrt.IsValid() != true {
			t.Errorf("Expected %q to be valid", "2025.07.06~7e64e8ba-r5")
		}

		nix := ToVersion("hello-2.12.1")

		if nix.Upstream.Major != 2 || nix.Upstream.Minor != 12 || nix.Upstream.Patch != 1 {
			t.Errorf("Expected 2.12.1 from %q, got %s", "hello-2.12.1", nix.String())
		}

		if nix.IsValid() != true {
			t.Errorf("Expected %q to be valid", "hello-2.12.1")
		}

	})

}

func TestVersionFactories(t *testing.T) {

	zeroVersion := NewVersion()

	if zeroVersion.Equal(Version{}) == false {
		t.Fatal("NewVersion() should return the zero value")
	}

	if zeroVersion.IsValid() == true {
		t.Fatal("the zero version should be invalid")
	}

	valid := []string{"0.0.0", "0.0.1", "0.1.0", "1.0.0", "1.23.45", "v1.23.45", "2:1.23.45"}
	invalid := []string{"", "not a version"}

	for _, value := range valid {
		if IsVersion(value) == false {
			t.Errorf("IsVersion(%q) = false, expected true", value)
		}
	}

	for _, value := range invalid {
		if IsVersion(value) == true {
			t.Errorf("IsVersion(%q) = true, expected false", value)
		}
	}

	for _, value := range invalid {
		if parsed := ParseVersion(value); parsed != nil {
			t.Errorf("ParseVersion(%q) = %v, expected nil", value, parsed.String())
		}
	}

	parsed := ParseVersion("1.23.45")

	if parsed == nil {
		t.Fatal("ParseVersion() returned nil for a valid value")
	}

	if parsed.Upstream.Major != 1 || parsed.Upstream.Minor != 23 || parsed.Upstream.Patch != 45 {
		t.Errorf("ParseVersion() = %+v", parsed)
	}

	emptyVersion := ToVersion("")

	if emptyVersion.Equal(Version{}) == false {
		t.Fatal("ToVersion(\"\") should return the zero value")
	}

	var version Version

	if version.Parse("1.23.45") == false {
		t.Fatal("Version.Parse() returned false for a valid value")
	}

	if version.Upstream.Major != 1 || version.Upstream.Minor != 23 || version.Upstream.Patch != 45 {
		t.Errorf("Version.Parse() = %+v", version)
	}

	if version.Parse("not a version") == true {
		t.Fatal("Version.Parse() returned true for invalid input")
	}

}

func TestVersionIsValid(t *testing.T) {

	cases := []struct {
		value string
		valid bool
	}{
		{"0.0.0", false},
		{"0.0.1", true},
		{"0.1.0", true},
		{"1.0.0", true},
		{"1.23.45", true},
	}

	for _, c := range cases {

		version := ToVersion(c.value)

		if result := version.IsValid(); result != c.valid {
			t.Errorf("ToVersion(%q).IsValid() = %v, expected %v", c.value, result, c.valid)
		}

	}

}

func TestVersionIncrementDecrement(t *testing.T) {

	base := ToVersion("2:1.23.45alpha")

	version := base
	version.NextEpoche()

	if version.Epoche != 3 {
		t.Errorf("NextEpoche() = %d, expected 3", version.Epoche)
	}

	version = base
	version.NextMajor()

	if version.Upstream.Major != 2 {
		t.Errorf("NextMajor() = %d, expected 2", version.Upstream.Major)
	}

	version = base
	version.NextMinor()

	if version.Upstream.Minor != 24 {
		t.Errorf("NextMinor() = %d, expected 24", version.Upstream.Minor)
	}

	version = base
	version.NextPatch()

	if version.Upstream.Patch != 46 {
		t.Errorf("NextPatch() = %d, expected 46", version.Upstream.Patch)
	}

	version = base
	version.PrevEpoche()

	if version.Epoche != 1 {
		t.Errorf("PrevEpoche() = %d, expected 1", version.Epoche)
	}

	version = base
	version.PrevMajor()

	if version.Upstream.Major != 0 {
		t.Errorf("PrevMajor() = %d, expected 0", version.Upstream.Major)
	}

	version = base
	version.PrevMinor()

	if version.Upstream.Minor != 22 {
		t.Errorf("PrevMinor() = %d, expected 22", version.Upstream.Minor)
	}

	version = base
	version.PrevPatch()

	if version.Upstream.Patch != 44 {
		t.Errorf("PrevPatch() = %d, expected 44", version.Upstream.Patch)
	}

	zero := NewVersion()
	zero.PrevEpoche()
	zero.PrevMajor()
	zero.PrevMinor()
	zero.PrevPatch()

	if zero.Equal(Version{}) == false {
		t.Fatal("Prev*() should not underflow the zero version")
	}

}

func TestVersionNextPrevRelease(t *testing.T) {

	version := ToVersion("1.23.45alpha")
	version.NextRelease()

	if version.Upstream.Release != "beta" {
		t.Errorf("NextRelease() = %q, expected beta", version.Upstream.Release)
	}

	version.PrevRelease()

	if version.Upstream.Release != "alpha" {
		t.Errorf("PrevRelease() = %q, expected alpha", version.Upstream.Release)
	}

	version.PrevRelease()

	if version.Upstream.Release != "alpha" {
		t.Errorf("PrevRelease() at the first release = %q, expected alpha", version.Upstream.Release)
	}

	letter := ToVersion("1.23.45a")
	letter.NextRelease()

	if letter.Upstream.Release != "b" {
		t.Errorf("NextRelease() = %q, expected b", letter.Upstream.Release)
	}

	letter.PrevRelease()

	if letter.Upstream.Release != "a" {
		t.Errorf("PrevRelease() = %q, expected a", letter.Upstream.Release)
	}

	empty := ToVersion("1.23.45")
	empty.NextRelease()

	if empty.Upstream.Release != "" {
		t.Errorf("NextRelease() on an empty release = %q, expected empty", empty.Upstream.Release)
	}

	last := ToVersion("1.23.45omega")
	last.NextRelease()

	if last.Upstream.Release != "omega" {
		t.Errorf("NextRelease() at the last release = %q, expected omega", last.Upstream.Release)
	}

}

func TestVersionComparisonExtra(t *testing.T) {

	version1 := ToVersion("1.23.45-r6")
	version2 := ToVersion("1.23.45-r6")
	version3 := ToVersion("1.23.45-r7")

	if version1.Equal(version2) == false {
		t.Fatal("identical versions should be equal")
	}

	if version1.IsAfter(version2) == true || version1.IsBefore(version2) == true {
		t.Fatal("identical versions should be neither after nor before each other")
	}

	if version1.IsBefore(version3) == false || version3.IsAfter(version1) == false {
		t.Fatal("release ordering failed")
	}

	if version1.Equal(version3) == true {
		t.Fatal("different releases should not be equal")
	}

}

func TestVersionHash(t *testing.T) {

	version1 := ToVersion("1.23.45-r6")
	version2 := ToVersion("1.23.45-r6")
	version3 := ToVersion("1.23.45-r7")

	if version1.Hash() == "" || version1.Hash() != version2.Hash() {
		t.Fatal("Hash() should be deterministic for equal versions")
	}

	if version1.Hash() == version3.Hash() {
		t.Fatal("different versions should not share a hash")
	}

}

func TestVersionSemanticStringString(t *testing.T) {

	version := ToVersion("2:1.23.45alpha+abc123~rev1")

	if version.SemanticString() != "2:1.23.45" {
		t.Errorf("SemanticString() = %q, expected \"2:1.23.45\"", version.SemanticString())
	}

	if version.String() != "2:1.23.45alpha+abc123~rev1" {
		t.Errorf("String() = %q", version.String())
	}

	zero := NewVersion()

	if zero.SemanticString() != "0:0.0.0" {
		t.Errorf("zero SemanticString() = %q", zero.SemanticString())
	}

	if zero.String() != "0:0.0.0" {
		t.Errorf("zero String() = %q", zero.String())
	}

}

func TestVersionNavigation(t *testing.T) {

	version := ToVersion("2:1.23.45alpha+abc123~rev1")

	earlier := version.ToEarlier(true, true)

	if earlier.Epoche != 2 || earlier.Upstream.Major != 1 {
		t.Errorf("ToEarlier() epoche/major = %+v", earlier)
	}

	if earlier.Upstream.Minor != 22 || earlier.Upstream.Patch != 44 {
		t.Errorf("ToEarlier() minor/patch = %d/%d, expected 22/44", earlier.Upstream.Minor, earlier.Upstream.Patch)
	}

	if earlier.Upstream.Release != "" || earlier.Upstream.Hash != "" || earlier.Revision != "" {
		t.Errorf("ToEarlier() should clear release metadata: %+v", earlier)
	}

	earlierDefault := version.ToEarlier(false, false)

	if result := earlierDefault.String(); result != "2:1.23.45" {
		t.Errorf("ToEarlier(false, false) = %q", result)
	}

	earliest := version.ToEarliest(true, true)

	if earliest.String() != "2:1.0.0" {
		t.Errorf("ToEarliest(true, true) = %q", earliest.String())
	}

	earliestDefault := version.ToEarliest(false, false)

	if result := earliestDefault.String(); result != "2:1.23.45" {
		t.Errorf("ToEarliest(false, false) = %q", result)
	}

	later := version.ToLater(true, true)

	if later.Epoche != 2 || later.Upstream.Minor != 24 || later.Upstream.Patch != 46 {
		t.Errorf("ToLater(true, true) = %+v", later)
	}

	if later.Upstream.Release != "" || later.Upstream.Hash != "" || later.Revision != "" {
		t.Errorf("ToLater() should clear release metadata: %+v", later)
	}

	laterDefault := version.ToLater(false, false)

	if result := laterDefault.String(); result != "2:0.23.45" {
		t.Errorf("ToLater(false, false) = %q", result)
	}

	latest := version.ToLatest(true, true)

	if latest.String() != "2:1.999999999.999999999" {
		t.Errorf("ToLatest(true, true) = %q", latest.String())
	}

	latestDefault := version.ToLatest(false, false)

	if result := latestDefault.String(); result != "2:1.23.45" {
		t.Errorf("ToLatest(false, false) = %q", result)
	}

	zeroMinor := ToVersion("1.0.0")
	zeroMinorEarlier := zeroMinor.ToEarlier(true, true)

	if result := zeroMinorEarlier.String(); result != "0:1.0.0" {
		t.Errorf("ToEarlier() at zero minor/patch = %q", result)
	}

	maxVersion := ToVersion("1.999999999.999999999")
	maxLater := maxVersion.ToLater(true, true)

	if maxLater.Upstream.Minor != 999999999 || maxLater.Upstream.Patch != 999999999 {
		t.Errorf("ToLater() should clamp at the maximum: %d/%d", maxLater.Upstream.Minor, maxLater.Upstream.Patch)
	}

}

func TestVersionJSON(t *testing.T) {

	original := ToVersion("2:1.23.45alpha+abc123~rev1")

	data, err := json.Marshal(original)

	if err != nil {
		t.Fatalf("Marshal() returned error: %v", err)
	}

	if string(data) != `"2:1.23.45alpha+abc123~rev1"` {
		t.Errorf("Marshal() = %s", string(data))
	}

	var decoded Version

	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal() returned error: %v", err)
	}

	if decoded.String() != original.String() {
		t.Errorf("Unmarshal() = %q, expected %q", decoded.String(), original.String())
	}

	simple := ToVersion("1.23.45")
	simpleData, err := json.Marshal(simple)

	if err != nil {
		t.Fatalf("Marshal() returned error: %v", err)
	}

	if string(simpleData) != `"0:1.23.45"` {
		t.Errorf("Marshal() = %s, expected \"0:1.23.45\"", string(simpleData))
	}

	var malformed Version

	if err := json.Unmarshal([]byte("123"), &malformed); err == nil {
		t.Fatal("Unmarshal() should fail for malformed JSON")
	}

	unchanged := original

	if err := json.Unmarshal([]byte(`"not a version"`), &unchanged); err != nil {
		t.Fatalf("Unmarshal() of an invalid string returned error: %v", err)
	}

	if unchanged.String() != original.String() {
		t.Fatal("Unmarshal() of an invalid string should leave the version unchanged")
	}

}
