package insights

import "tholian-core/insights/templates"
import utils_tags "tholian-core/utils/tags"
import compress_gzip "tholian-core/utils/compress/gzip"
import "encoding/json"

func (database *Database) WriteProxies() bool {

	buffer_proxies, err0_proxies := json.MarshalIndent(database.Proxies, "", "\t")
	buffer_proxies_gzip, err0_proxies_gzip := compress_gzip.Compress(buffer_proxies)

	if err0_proxies == nil && err0_proxies_gzip == nil {

		result1_proxies := database.Write("/proxies/Proxies.json.gz", buffer_proxies_gzip)
		result2_proxies := database.Write("/proxies/Proxies.json", buffer_proxies)

		for _, tag := range []string{"guard", "oversight"} {

			insight_go, insight_embed_go, err := templates.Cache(&templates.Config{
				Name:     "Proxies",
				Package:  "proxies",
				BuildTag: utils_tags.ToBuildTags(tag, database.GetDistributionNames()),
				Struct:   "structs.Proxy",
			})

			if err == nil {
				database.Write("/proxies/Proxies.go", insight_go)
				database.Write("/proxies/Proxies_embed." + tag + ".go", insight_embed_go)
			}

		}

		if result1_proxies && result2_proxies {
			return true
		}

	}

	return false

}
