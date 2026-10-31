package insights

import "tholian-core/insights/templates"
import utils_tags "tholian-core/utils/tags"
import compress_gzip "tholian-core/utils/compress/gzip"
import "encoding/json"
import "slices"
import "sort"
import "strings"

func (database *Database) GetDistributionNames() []string {

	result := make([]string, 0)

	for _, distribution := range database.Distributions.Distributions {

		if strings.Contains(distribution.Name, "-") == false && slices.Contains(result, distribution.Name) == false {
			result = append(result, distribution.Name)
		}

	}

	sort.Strings(result)

	return result

}

func (database *Database) WriteDistributions() bool {

	buffer_distributions, err0_distributions := json.MarshalIndent(database.Distributions, "", "\t")
	buffer_distributions_gzip, err0_distributions_gzip := compress_gzip.Compress(buffer_distributions)

	if err0_distributions == nil && err0_distributions_gzip == nil {

		result1_distributions := database.Write("/distributions/Distributions.json.gz", buffer_distributions_gzip)
		result2_distributions := database.Write("/distributions/Distributions.json", buffer_distributions)

		for _, tag := range []string{"scout", "guard", "oversight"} {

			insight_go, insight_embed_go, err := templates.Cache(&templates.Config{
				Name:     "Distributions",
				Package:  "distributions",
				BuildTag: utils_tags.ToBuildTags(tag, database.GetDistributionNames()),
				Struct:   "structs.Distribution",
			})

			if err == nil {
				database.Write("/distributions/Distributions.go", insight_go)
				database.Write("/distributions/Distributions_embed."+tag+".go", insight_embed_go)
			}

		}

		if result1_distributions && result2_distributions {
			return true
		}

	}

	return false

}
