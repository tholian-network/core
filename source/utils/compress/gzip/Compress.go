package gzip

import "bytes"
import "compress/gzip"

func Compress(source []byte) ([]byte, error) {

	var result_buffer []byte
	var result_error error = nil

	buffer := bytes.NewBuffer([]byte{})
	writer := gzip.NewWriter(buffer)
	_, err1 := writer.Write(source)

	if err1 == nil {

		err2 := writer.Flush()
		err3 := writer.Close()

		if err2 == nil && err3 == nil {
			result_buffer = buffer.Bytes()
		} else if err2 != nil {
			result_error = err2
		} else if err3 != nil {
			result_error = err3
		}

	} else {
		result_error = err1
	}

	return result_buffer, result_error

}
