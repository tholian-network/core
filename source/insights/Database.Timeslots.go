package insights

import "tholian-core/insights/templates"
import utils_tags "tholian-core/utils/tags"
import compress_gzip "tholian-core/utils/compress/gzip"
import "encoding/json"

func (database *Database) WriteTimeslots() bool {

	buffer_timeslots, err0_timeslots:= json.MarshalIndent(database.Timeslots, "", "\t")
	buffer_timeslots_gzip, err0_timeslots_gzip := compress_gzip.Compress(buffer_timeslots)

	if err0_timeslots == nil && err0_timeslots_gzip == nil {

		result1_timeslots := database.Write("/timeslots/Timeslots.json.gz", buffer_timeslots_gzip)
		result2_timeslots := database.Write("/timeslots/Timeslots.json", buffer_timeslots)

		for _, tag := range []string{"scout", "guard", "oversight"} {

			insight_go, insight_embed_go, err := templates.Cache(&templates.Config{
				Name:     "Timeslots",
				Package:  "timeslots",
				BuildTag: utils_tags.ToBuildTags(tag, database.GetDistributionNames()),
				Struct:   "structs.Timeslot",
			})

			if err == nil {
				database.Write("/timeslots/Timeslots.go", insight_go)
				database.Write("/timeslots/Timeslots_embed." + tag + ".go", insight_embed_go)
			}

		}

		if result1_timeslots && result2_timeslots {
			return true
		}

	}

	return false

}
