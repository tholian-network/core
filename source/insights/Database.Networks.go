package insights

import "tholian-core/insights/templates"
import utils_tags "tholian-core/utils/tags"
import compress_gzip "tholian-core/utils/compress/gzip"
import "encoding/json"

func (database *Database) WriteNetworks() bool {

	buffer_networks, err0_networks := json.MarshalIndent(database.Networks, "", "\t")
	buffer_networks_gzip, err0_networks_gzip := compress_gzip.Compress(buffer_networks)

	if err0_networks == nil && err0_networks_gzip == nil {

		result1_networks := database.Write("/networks/Networks.json.gz", buffer_networks_gzip)
		result2_networks := database.Write("/networks/Networks.json", buffer_networks)

		for _, tag := range []string{"guard", "oversight"} {

			insight_go, insight_embed_go, err := templates.Cache(&templates.Config{
				Name:     "Networks",
				Package:  "networks",
				BuildTag: utils_tags.ToBuildTags(tag, database.GetDistributionNames()),
				Struct:   "structs.Network",
			})

			if err == nil {
				database.Write("/networks/Networks.go", insight_go)
				database.Write("/networks/Networks_embed."+tag+".go", insight_embed_go)
			}

		}

		if result1_networks && result2_networks {
			return true
		}

	}

	return false

}
