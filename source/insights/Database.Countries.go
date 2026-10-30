package insights

import "tholian-core/insights/templates"
import utils_tags "tholian-core/utils/tags"
import compress_gzip "tholian-core/utils/compress/gzip"
import "encoding/json"

func (database *Database) WriteCountries() bool {

	buffer_countries, err0_countries := json.MarshalIndent(database.Countries, "", "\t")
	buffer_countries_gzip, err0_countries_gzip := compress_gzip.Compress(buffer_countries)

	if err0_countries == nil && err0_countries_gzip == nil {

		result1_countries := database.Write("/countries/Countries.json.gz", buffer_countries_gzip)
		result2_countries := database.Write("/countries/Countries.json", buffer_countries)

		for _, tag := range []string{"scout", "guard", "oversight"} {

			insight_go, insight_embed_go, err := templates.Cache(&templates.Config{
				Name:     "Countries",
				Package:  "countries",
				BuildTag: utils_tags.ToBuildTags(tag, database.GetDistributionNames()),
				Struct:   "structs.Country",
			})

			if err == nil {
				database.Write("/countries/Countries.go", insight_go)
				database.Write("/countries/Countries_embed." + tag + ".go", insight_embed_go)
			}

		}

		if result1_countries && result2_countries {
			return true
		}

	}

	return false

}
