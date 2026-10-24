package types

import "bytes"
import "slices"
import "strconv"
import "strings"

var version_letters []string = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var version_prereleases []string = []string{"alpha", "beta", "gamma", "delta", "epsilon", "zeta", "eta", "theta", "iota", "kappa", "lambda", "mu", "nu", "xi", "omicron", "pi", "rho", "sigma", "tau", "upsilon", "phi", "chi", "psi", "omega"}
var version_max_number uint = uint(999999999)

type Version struct {
	Epoche   uint `json:"epoche"`
	Upstream struct {
		Major   uint   `json:"major"`             // 1 or 20240101
		Minor   uint   `json:"minor"`             // 2
		Patch   uint   `json:"patch"`             // 3
		Release string `json:"release,omitempty"` // alpha
		Hash    string `json:"hash,omitempty"`    // abc123abcdef
	} `json:"upstream"`
	Revision string `json:"revision,omitempty"`
}

func ParseVersion(value string) *Version {

	var result *Version = nil
	var version Version

	version.Epoche = 0
	version.Upstream.Major = 0
	version.Upstream.Minor = 0
	version.Upstream.Patch = 0
	version.Upstream.Release = ""
	version.Upstream.Hash = ""
	version.Revision = ""

	if value != "" && version.Parse(value) == true {
		result = &version
	}

	return result

}

func ToVersion(value string) Version {

	var version Version

	tmp := ParseVersion(value)

	if tmp != nil {
		version = *tmp
	}

	return version

}

func (version *Version) IsAfter(other Version) bool {

	var result bool

	if version.Epoche > other.Epoche {
		result = true
	} else if version.Epoche == other.Epoche {

		if version.Upstream.Major > other.Upstream.Major {
			result = true
		} else if version.Upstream.Major == other.Upstream.Major {

			if version.Upstream.Minor > other.Upstream.Minor {
				result = true
			} else if version.Upstream.Minor == other.Upstream.Minor {

				if version.Upstream.Patch > other.Upstream.Patch {
					result = true
				} else if version.Upstream.Patch == other.Upstream.Patch {

					compare := compareVersionRelease(version.Upstream.Release, other.Upstream.Release)

					if compare == 1 {
						result = true
					} else if compare == 0 {

						if version.Revision > other.Revision {
							result = true
						} else if version.Revision == other.Revision {
							result = false
						}

					}

				}

			}

		}

	}

	return result

}

func (version *Version) IsBefore(other Version) bool {

	var result bool

	if version.Epoche < other.Epoche {
		result = true
	} else if version.Epoche == other.Epoche {

		if version.Upstream.Major < other.Upstream.Major {
			result = true
		} else if version.Upstream.Major == other.Upstream.Major {

			if version.Upstream.Minor < other.Upstream.Minor {
				result = true
			} else if version.Upstream.Minor == other.Upstream.Minor {

				if version.Upstream.Patch < other.Upstream.Patch {
					result = true
				} else if version.Upstream.Patch == other.Upstream.Patch {

					compare := compareVersionRelease(version.Upstream.Release, other.Upstream.Release)

					if compare == -1 {
						result = true
					} else if compare == 0 {

						if version.Revision < other.Revision {
							result = true
						} else if version.Revision == other.Revision {
							result = false
						}

					}

				}

			}

		}

	}

	return result

}

func (version *Version) IsSame(other Version) bool {

	var result bool

	if version.Epoche == other.Epoche {

		if version.Upstream.Major == other.Upstream.Major {

			if version.Upstream.Minor == other.Upstream.Minor {

				if version.Upstream.Patch == other.Upstream.Patch {

					compare := compareVersionRelease(version.Upstream.Release, other.Upstream.Release)

					if compare == 0 {
						result = true
					}

				}

			}

		}

	}

	return result

}

func (version *Version) IsValid() bool {

	var result bool

	if version.Upstream.Major != 0 || version.Upstream.Minor != 0 || version.Upstream.Patch != 0 {
		result = true
	}

	return result

}

func (version *Version) NextEpoche() {
	version.Epoche += 1
}

func (version *Version) NextMajor() {
	version.Upstream.Major += 1
}

func (version *Version) NextMinor() {
	version.Upstream.Minor += 1
}

func (version *Version) NextPatch() {
	version.Upstream.Patch += 1
}

func (version *Version) NextRelease() {

	if version.Upstream.Release != "" {

		if slices.Contains(version_letters, version.Upstream.Release) {

			index := slices.Index(version_letters, version.Upstream.Release)

			if index < len(version_prereleases)-1 {
				version.Upstream.Release = version_letters[(index+1)%len(version_prereleases)]
			}

		} else if slices.Contains(version_prereleases, version.Upstream.Release) {

			index := slices.Index(version_prereleases, version.Upstream.Release)

			if index < len(version_prereleases)-1 {
				version.Upstream.Release = version_prereleases[(index+1)%len(version_prereleases)]
			}

		}

	}

}

func (version *Version) Parse(value string) bool {

	if strings.Contains(value, ":") {

		chunk := strings.TrimSpace(value[0:strings.Index(value, ":")])

		if isNumber(chunk) {

			num, err := strconv.ParseUint(chunk, 10, 8)

			if err == nil {
				version.Epoche = uint(num)
				value = value[strings.Index(value, ":")+1:]
			}

		}

	}

	parts := toVersionParts(strings.ToLower(value))
	major_set := false
	minor_set := false
	patch_set := false
	release_set := false
	hash_set := false

	if len(parts) >= 4 && len(parts[0]) == 4 && isNumber(parts[0]) && isNumber(parts[2]) && isNumber(parts[3]) {

		if parts[1] == "q1" {
			parts[1] = "1"
		} else if parts[1] == "q2" {
			parts[1] = "2"
		} else if parts[1] == "q3" {
			parts[1] = "3"
		} else if parts[1] == "q4" {
			parts[1] = "4"
		}

		major, err1 := strconv.ParseUint(parts[0], 10, 64)
		minor, err2 := strconv.ParseUint(parts[1], 10, 64)
		patch, err3 := strconv.ParseUint(parts[2], 10, 64)

		if err1 == nil && err2 == nil && err3 == nil {

			version.Upstream.Major = uint(major)
			version.Upstream.Minor = uint(minor)
			version.Upstream.Patch = uint(patch)

			major_set = true
			minor_set = true
			patch_set = true

		}

		parts = parts[3:]

	} else if len(parts) >= 3 && isNumber(parts[0]) && isNumber(parts[1]) && isNumber(parts[2]) {

		major, err1 := strconv.ParseUint(parts[0], 10, 64)
		minor, err2 := strconv.ParseUint(parts[1], 10, 64)
		patch, err3 := strconv.ParseUint(parts[2], 10, 64)

		if err1 == nil && err2 == nil && err3 == nil {

			version.Upstream.Major = uint(major)
			version.Upstream.Minor = uint(minor)
			version.Upstream.Patch = uint(patch)

			major_set = true
			minor_set = true
			patch_set = true

		}

		parts = parts[3:]

	} else if len(parts) >= 2 && isNumber(parts[0]) && isNumber(parts[1]) {

		major, err1 := strconv.ParseUint(parts[0], 10, 64)
		minor, err2 := strconv.ParseUint(parts[1], 10, 64)

		if err1 == nil && err2 == nil {

			version.Upstream.Major = uint(major)
			version.Upstream.Minor = uint(minor)

			major_set = true
			minor_set = true

		}

		parts = parts[2:]

	} else if len(parts) >= 1 && isNumber(parts[0]) {

		major, err1 := strconv.ParseUint(parts[0], 10, 64)

		if err1 == nil {
			version.Upstream.Major = uint(major)
			major_set = true
		}

		parts = parts[1:]

	}

	if len(parts) > 0 {

		if !major_set && !minor_set && !patch_set {

			if !isVersionHash(parts[0]) && isNumber(string(parts[0][0])) {

				_, number, suffix := splitNumber(parts[0])
				major, err1 := strconv.ParseUint(number, 10, 64)

				if err1 == nil {
					version.Upstream.Major = uint(major)
					parts[0] = suffix
					major_set = true
					minor_set = true
					patch_set = true
				}

			}

		} else if major_set && !minor_set && !patch_set {

			if !isVersionHash(parts[0]) && isNumber(string(parts[0][0])) {

				_, number, suffix := splitNumber(parts[0])
				minor, err1 := strconv.ParseUint(number, 10, 64)

				if err1 == nil {
					version.Upstream.Minor = uint(minor)
					parts[0] = suffix
					minor_set = true
					patch_set = true
				}

			}

		} else if major_set && minor_set && !patch_set {

			if !isVersionHash(parts[0]) && isNumber(string(parts[0][0])) {

				_, number, suffix := splitNumber(parts[0])
				patch, err1 := strconv.ParseUint(number, 10, 64)

				if err1 == nil {
					version.Upstream.Patch = uint(patch)
					parts[0] = suffix
					patch_set = true
				}

			}

		}

	}

	revision_parts := make([]string, 0)

	for p := 0; p < len(parts); p++ {

		part := parts[p]

		if isVersionRelease(part) {

			if release_set {
				revision_parts = append(revision_parts, part)
			} else {
				version.Upstream.Release = part
			}

			release_set = true

		} else if isVersionHash(part) {

			if hash_set {
				revision_parts = append(revision_parts, part)
			} else {
				version.Upstream.Hash = part
			}

			hash_set = true

		} else if isNumber(part) {
			revision_parts = append(revision_parts, part)
		} else {
			revision_parts = append(revision_parts, part)
		}

	}

	if len(revision_parts) > 0 {
		version.Revision = strings.Join(revision_parts, "-")
	}

	if major_set {
		return true
	}

	return false

}

func (version *Version) PrevEpoche() {

	if version.Epoche > 0 {
		version.Epoche -= 1
	}

}

func (version *Version) PrevMajor() {

	if version.Upstream.Major > 0 {
		version.Upstream.Major -= 1
	}

}

func (version *Version) PrevMinor() {

	if version.Upstream.Minor > 0 {
		version.Upstream.Minor -= 1
	}

}

func (version *Version) PrevPatch() {

	if version.Upstream.Patch > 0 {
		version.Upstream.Patch -= 1
	}

}

func (version *Version) PrevRelease() {

	if version.Upstream.Release != "" {

		if slices.Contains(version_letters, version.Upstream.Release) {

			index := slices.Index(version_letters, version.Upstream.Release)

			if index > 0 {
				version.Upstream.Release = version_letters[(index-1)%len(version_prereleases)]
			}

		} else if slices.Contains(version_prereleases, version.Upstream.Release) {

			index := slices.Index(version_prereleases, version.Upstream.Release)

			if index > 0 {
				version.Upstream.Release = version_prereleases[(index-1)%len(version_prereleases)]
			}

		}

	}

}

func (version *Version) SemanticString() string {

	var buffer bytes.Buffer

	buffer.WriteString(strconv.FormatUint(uint64(version.Epoche), 10))
	buffer.WriteString(":")
	buffer.WriteString(strconv.FormatUint(uint64(version.Upstream.Major), 10))
	buffer.WriteString(".")
	buffer.WriteString(strconv.FormatUint(uint64(version.Upstream.Minor), 10))
	buffer.WriteString(".")
	buffer.WriteString(strconv.FormatUint(uint64(version.Upstream.Patch), 10))

	return buffer.String()

}

func (version *Version) String() string {

	var buffer bytes.Buffer

	buffer.WriteString(strconv.FormatUint(uint64(version.Epoche), 10))
	buffer.WriteString(":")
	buffer.WriteString(strconv.FormatUint(uint64(version.Upstream.Major), 10))
	buffer.WriteString(".")
	buffer.WriteString(strconv.FormatUint(uint64(version.Upstream.Minor), 10))
	buffer.WriteString(".")
	buffer.WriteString(strconv.FormatUint(uint64(version.Upstream.Patch), 10))

	if version.Upstream.Release != "" {
		buffer.WriteString(version.Upstream.Release)
	}

	if version.Upstream.Hash != "" {
		buffer.WriteString("+")
		buffer.WriteString(version.Upstream.Hash)
	}

	if version.Revision != "" {
		buffer.WriteString("~")
		buffer.WriteString(version.Revision)
	}

	return buffer.String()

}

func (version *Version) ToEarlier(minor bool, patch bool) Version {

	var result Version

	result.Epoche = version.Epoche
	result.Upstream.Major = version.Upstream.Major

	if minor == true {

		if version.Upstream.Minor > 0 {
			result.Upstream.Minor = uint(version.Upstream.Minor - 1)
		} else {
			result.Upstream.Minor = 0
		}

	} else {
		result.Upstream.Minor = version.Upstream.Minor
	}

	if patch == true {

		if version.Upstream.Patch > 0 {
			result.Upstream.Patch = uint(version.Upstream.Patch - 1)
		} else {
			result.Upstream.Patch = 0
		}

	} else {
		result.Upstream.Patch = version.Upstream.Patch
	}

	return result

}

func (version *Version) ToEarliest(minor bool, patch bool) Version {

	var result Version

	result.Epoche = version.Epoche
	result.Upstream.Major = version.Upstream.Major

	if minor == true {
		result.Upstream.Minor = uint(0)
	} else {
		result.Upstream.Minor = version.Upstream.Minor
	}

	if patch == true {
		result.Upstream.Patch = uint(0)
	} else {
		result.Upstream.Patch = version.Upstream.Patch
	}

	result.Upstream.Release = ""
	result.Upstream.Hash = ""
	result.Revision = ""

	return result

}

func (version *Version) ToLater(minor bool, patch bool) Version {

	var result Version

	result.Epoche = version.Epoche

	if minor == true {

		if version.Upstream.Minor < version_max_number {
			result.Upstream.Minor = uint(version.Upstream.Minor + 1)
		} else {
			result.Upstream.Minor = version_max_number
		}

	} else {
		result.Upstream.Minor = version.Upstream.Minor
	}

	if patch == true {

		if version.Upstream.Patch < version_max_number {
			result.Upstream.Patch = uint(version.Upstream.Patch + 1)
		} else {
			result.Upstream.Patch = version_max_number
		}

	} else {
		result.Upstream.Patch = version.Upstream.Patch
	}

	result.Upstream.Release = ""
	result.Upstream.Hash = ""
	result.Revision = ""

	return result

}

func (version *Version) ToLatest(minor bool, patch bool) Version {

	var result Version

	result.Epoche = version.Epoche
	result.Upstream.Major = version.Upstream.Major

	if minor == true {
		result.Upstream.Minor = version_max_number
	} else {
		result.Upstream.Minor = version.Upstream.Minor
	}

	if patch == true {
		result.Upstream.Patch = version_max_number
	} else {
		result.Upstream.Patch = version.Upstream.Patch
	}

	result.Upstream.Release = ""
	result.Upstream.Hash = ""
	result.Revision = ""

	return result

}

func (version Version) MarshalJSON() ([]byte, error) {

	quoted := strconv.Quote(version.String())

	return []byte(quoted), nil

}

func (version *Version) UnmarshalJSON(data []byte) error {

	unquoted, err := strconv.Unquote(string(data))

	if err != nil {
		return err
	}

	check := ToVersion(unquoted)

	if check.IsValid() {

		version.Epoche = check.Epoche
		version.Upstream.Major = check.Upstream.Major
		version.Upstream.Minor = check.Upstream.Minor
		version.Upstream.Patch = check.Upstream.Patch
		version.Upstream.Release = check.Upstream.Release
		version.Upstream.Hash = check.Upstream.Hash
		version.Revision = check.Revision

	}

	return nil

}
