package insights

import compress_gzip "tholian-core/utils/compress/gzip"
import "strings"

func readBuffer(database *Database, file string) []byte {

	buffer := database.Read(file)

	if len(buffer) == 0 && strings.HasSuffix(file, ".json") {

		compressed := database.Read(file + ".gz")

		if len(compressed) > 0 {

			decompressed, err := compress_gzip.Decompress(compressed)

			if err == nil {
				buffer = decompressed
			}

		}

	}

	return buffer

}
