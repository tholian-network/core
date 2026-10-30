package insights

import "tholian-core/insights/templates"
import utils_tags "tholian-core/utils/tags"
import compress_gzip "tholian-core/utils/compress/gzip"
import "encoding/json"

func (database *Database) WriteDevices() bool {

	buffer_devices, err0_devices := json.MarshalIndent(database.Devices, "", "\t")
	buffer_devices_gzip, err0_devices_gzip := compress_gzip.Compress(buffer_devices)

	if err0_devices == nil && err0_devices_gzip == nil {

		result1_devices := database.Write("/devices/Devices.json.gz", buffer_devices_gzip)
		result2_devices := database.Write("/devices/Devices.json", buffer_devices)

		for _, tag := range []string{"scout", "guard", "oversight"} {

			insight_go, insight_embed_go, err := templates.Cache(&templates.Config{
				Name:     "Devices",
				Package:  "devices",
				BuildTag: utils_tags.ToBuildTags(tag, database.GetDistributionNames()),
				Struct:   "structs.Device",
			})

			if err == nil {
				database.Write("/devices/Devices.go", insight_go)
				database.Write("/devices/Devices_embed." + tag + ".go", insight_embed_go)
			}

		}

		if result1_devices && result2_devices {
			return true
		}

	}

	return false

}
