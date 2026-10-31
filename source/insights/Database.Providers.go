package insights

import "tholian-core/insights/templates"
import utils_tags "tholian-core/utils/tags"
import compress_gzip "tholian-core/utils/compress/gzip"
import "encoding/json"

func (database *Database) WriteProviders() bool {

	buffer_providers, err0_providers := json.MarshalIndent(database.Providers, "", "\t")
	buffer_providers_gzip, err0_providers_gzip := compress_gzip.Compress(buffer_providers)

	if err0_providers == nil && err0_providers_gzip == nil {

		result1_providers := database.Write("/providers/Providers.json.gz", buffer_providers_gzip)
		result2_providers := database.Write("/providers/Providers.json", buffer_providers)

		for _, tag := range []string{"guard", "oversight"} {

			insight_go, insight_embed_go, err := templates.Cache(&templates.Config{
				Name:     "Providers",
				Package:  "providers",
				BuildTag: utils_tags.ToBuildTags(tag, database.GetDistributionNames()),
				Struct:   "structs.Provider",
			})

			if err == nil {
				database.Write("/providers/Providers.go", insight_go)
				database.Write("/providers/Providers_embed."+tag+".go", insight_embed_go)
			}

		}

		if result1_providers && result2_providers {
			return true
		}

	}

	return false

}
