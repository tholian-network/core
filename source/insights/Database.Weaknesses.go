package insights

import "tholian-core/insights/templates"
import utils_tags "tholian-core/utils/tags"
import compress_gzip "tholian-core/utils/compress/gzip"
import "encoding/json"

func (database *Database) WriteWeaknesses() bool {

	buffer_weaknesses, err0_weaknesses := json.MarshalIndent(database.Weaknesses, "", "\t")
	buffer_weaknesses_gzip, err0_weaknesses_gzip := compress_gzip.Compress(buffer_weaknesses)

	if err0_weaknesses == nil && err0_weaknesses_gzip == nil {

		result1_weaknesses := database.Write("/weaknesses/Weaknesses.json.gz", buffer_weaknesses_gzip)
		result2_weaknesses := database.Write("/weaknesses/Weaknesses.json", buffer_weaknesses)

		for _, tag := range []string{"scout", "guard", "oversight"} {

			insight_go, insight_embed_go, err := templates.Cache(&templates.Config{
				Name:     "Weaknesses",
				Package:  "weaknesses",
				BuildTag: utils_tags.ToBuildTags(tag, database.GetDistributionNames()),
				Struct:   "structs.Weakness",
			})

			if err == nil {
				database.Write("/weaknesses/Weaknesses.go", insight_go)
				database.Write("/weaknesses/Weaknesses_embed." + tag + ".go", insight_embed_go)
			}

		}

		if result1_weaknesses && result2_weaknesses {
			return true
		}

	}

	return false

}
