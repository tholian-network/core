package tags

import "sort"
import "strings"

func ToBuildTags(tag string, distributions []string) string {

	build_tag := make([]string, 0)

	switch tag {
	case "convoy":

		build_tag = append(build_tag, "convoy")

		sort.Strings(distributions)

		for _, distribution := range distributions {
			build_tag = append(build_tag, "convoy_"+distribution)
		}

	case "firewall":

		build_tag = append(build_tag, "firewall")

	case "guard":

		build_tag = append(build_tag, "guard")

		sort.Strings(distributions)

		for _, distribution := range distributions {
			build_tag = append(build_tag, "guard_"+distribution)
		}

	case "oversight":

		build_tag = append(build_tag, "oversight")

	case "scout":

		build_tag = append(build_tag, "scout")

		sort.Strings(distributions)

		for _, distribution := range distributions {
			build_tag = append(build_tag, "scout_"+distribution)
		}

	case "warps":

		build_tag = append(build_tag, "warps")

	}

	if len(build_tag) > 1 {
		return "("+strings.Join(build_tag, " || ")+")"
	} else if len(build_tag) == 1 {
		return build_tag[0]
	} else {
		return ""
	}

}
