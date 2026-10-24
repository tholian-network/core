package types

import "testing"

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

		if version1.IsSame(v_1_23_45_r6) != true {
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

		if version1.IsSame(v_1_23_45_r6) != true {
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

		if version2.IsSame(v_1_23_45_r6) != true {
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

		if version3.IsSame(v_1_23_45_r6) != true {
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

		if version4.IsSame(v2_1_23_45_r6) != true {
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

		if version1.IsSame(v_1_23_45_r6) != true {
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

		if version3.IsSame(v_1_23_45_deb6) != true {
			t.Errorf("Expected %s to be the same as %s", version3.String(), v_1_23_45_deb6.String())
		}

		if version3.IsBefore(v_1_23_45_deb7) != true {
			t.Errorf("Expected %s to be before %s", version3.String(), v_1_23_45_deb7.String())
		}

		version4 := ToVersion("0:1.23.45+dfsg-123~deb6u6")

		if version4.IsAfter(v_1_23_45_deb5) != true {
			t.Errorf("Expected %s to be after %s", version4.String(), v_1_23_45_deb5.String())
		}

		if version4.IsSame(v_1_23_45_deb6) != true {
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

		if version1.IsSame(v_1_23_45_1ubuntu66) != true {
			t.Errorf("Expected %s to be the same as %s", version1.String(), v_1_23_45_1ubuntu66.String())
		}

		if version1.IsBefore(v_1_23_45_1ubuntu67) != true {
			t.Errorf("Expected %s to be before %s", version1.String(), v_1_23_45_1ubuntu67.String())
		}

		version2 := ToVersion("2023.12.31-1ubuntu6~21.10.2")

		if version2.IsAfter(v_2023_12_31_1ubuntu6_20_04_03) != true {
			t.Errorf("Expected %s to be after %s", version2.String(), v_2023_12_31_1ubuntu6_20_04_03.String())
		}

		if version2.IsSame(v_2023_12_31_1ubuntu6_21_10_02) != true {
			t.Errorf("Expected %s to be the same as %s", version2.String(), v_2023_12_31_1ubuntu6_21_10_02.String())
		}

		if version2.IsBefore(v_2023_12_31_1ubuntu6_22_04_01) != true {
			t.Errorf("Expected %s to be before %s", version2.String(), v_2023_12_31_1ubuntu6_22_04_01.String())
		}

		version3 := ToVersion("2023.12.31-1ubuntu6")

		if version3.IsSame(v_2023_12_31_1ubuntu6_20_04_03) != true {
			t.Errorf("Expected %s to be the same as %s", version3.String(), v_2023_12_31_1ubuntu6_20_04_03.String())
		}

		if version3.IsSame(v_2023_12_31_1ubuntu6_21_10_02) != true {
			t.Errorf("Expected %s to be the same as %s", version3.String(), v_2023_12_31_1ubuntu6_21_10_02.String())
		}

		if version3.IsSame(v_2023_12_31_1ubuntu6_22_04_01) != true {
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

		if version1.IsSame(v_1_23_45_r7) != true {
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

		if version3.IsSame(v1_1_23_45) != true {
			t.Errorf("Expected %s to be the same as %s", version3.String(), v1_1_23_45.String())
		}

		if version3.IsBefore(v2_1_23_45) != true {
			t.Errorf("Expected %s to be before %s", version3.String(), v2_1_23_45.String())
		}

		version4 := ToVersion("1.23.45-r7-8.git2a8f9d8.el9")

		if version4.IsAfter(v_1_23_45_r6) != true {
			t.Errorf("Expected %s to be after %s", version4.String(), v_1_23_45_r6.String())
		}

		if version4.IsSame(v_1_23_45_r7) != true {
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

		if version1.IsSame(v_1_23_45_7) != true {
			t.Errorf("Expected %s to be the same as %s", version1.String(), v_1_23_45_7.String())
		}

		if version1.IsBefore(v_1_23_45_8) != true {
			t.Errorf("Expected %s to be before %s", version1.String(), v_1_23_45_8.String())
		}

		version2 := ToVersion("1.23.45.7_g31486f40-8.9")

		if version2.IsAfter(v_1_23_44) != true {
			t.Errorf("Expected %s to be after %s", version2.String(), v_1_23_44.String())
		}

		if version2.IsSame(v_1_23_45) != true {
			t.Errorf("Expected %s to be the same as %s", version2.String(), v_1_23_45.String())
		}
		if version2.IsBefore(v_1_23_46) != true {
			t.Errorf("Expected %s to be before %s", version2.String(), v_1_23_46.String())
		}

		if version2.IsAfter(v_1_23_45_6) != true {
			t.Errorf("Expected %s to be after %s", version2.String(), v_1_23_45_6.String())
		}

		if version2.IsSame(v_1_23_45_7) != true {
			t.Errorf("Expected %s to be the same as %s", version2.String(), v_1_23_45_7.String())
		}

		if version2.IsBefore(v_1_23_45_8) != true {
			t.Errorf("Expected %s to be before %s", version2.String(), v_1_23_45_8.String())
		}

		version3 := ToVersion("1.23.45.20240305.897f2593b3-1.1")

		if version3.IsAfter(v_1_23_44) != true {
			t.Errorf("Expected %s to be after %s", version3.String(), v_1_23_44.String())
		}

		if version3.IsSame(v_1_23_45) != true {
			t.Errorf("Expected %s to be the same as %s", version3.String(), v_1_23_45.String())
		}

		if version3.IsBefore(v_1_23_46) != true {
			t.Errorf("Expected %s to be before %s", version3.String(), v_1_23_46.String())
		}

	})

}
