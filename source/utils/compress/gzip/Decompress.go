package gzip

import "bytes"
import "compress/gzip"
import "io"

func Decompress(source []byte) ([]byte, error) {

	var result_buffer []byte
	var result_error error = nil

	buffer := bytes.NewBuffer(source)
	reader, err1 := gzip.NewReader(buffer)

	if err1 == nil {

		decompressed, err2 := io.ReadAll(reader)

		defer reader.Close()

		if err2 == nil {
			result_buffer = decompressed
		} else {
			result_error = err2
		}

	} else {
		result_error = err1
	}

	return result_buffer, result_error

}
