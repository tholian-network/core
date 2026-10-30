package photonos

import "strconv"

func ToRelease(platform string) string {

	for name, val := range Releases {

		version := strconv.FormatFloat(float64(val), 'f', 1, 32)

		if platform == "VMWare Photon OS "+version ||
			platform == "Photon OS "+version ||
			platform == "PhotonOS "+version {

			return name

		}

	}

	return ""

}
