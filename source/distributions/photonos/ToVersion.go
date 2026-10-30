package photonos

import "strconv"

func ToVersion(platform string) string {

	for release, val := range Releases {

		version := strconv.FormatFloat(float64(val), 'f', 1, 32)

		if platform == "VMWare Photon OS "+version ||
			platform == "Photon OS "+version ||
			platform == "PhotonOS "+version {

			return version

		} else if platform == release {

			return version

		}

	}

	return ""

}
