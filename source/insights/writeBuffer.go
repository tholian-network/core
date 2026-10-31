package insights

import compress_gzip "tholian-core/utils/compress/gzip"

func writeBuffer(database *Database, file string, buffer []byte) bool {

	buffer_gzip, err := compress_gzip.Compress(buffer)

	if err == nil {

		result1 := database.Write(file+".gz", buffer_gzip)
		result2 := database.Write(file, buffer)

		if result1 && result2 {
			return true
		}

	}

	return false

}
