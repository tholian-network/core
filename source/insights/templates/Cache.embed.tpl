{{if .BuildTag}}//go:build {{.BuildTag}}

{{end}}package {{.Package}}

import "bytes"
import "compress/gzip"
import "encoding/json"
import _ "embed"
import "io"
import "log"

//go:embed {{.Name}}.json.gz
var embedded_{{.Name}} []byte

func init() {

	decompress := func(compressed []byte) ([]byte, error) {

		buffer := bytes.NewBuffer(compressed)
		reader, err1 := gzip.NewReader(buffer)

		if err1 == nil {

			decompressed, err2 := io.ReadAll(reader)
			defer reader.Close()

			if err2 == nil {
				return decompressed, nil
			} else {
				return []byte{}, err2
			}

		} else {
			return []byte{}, err1
		}

	}


	buffer, err1 := decompress(embedded_{{.Name}})

	if err1 == nil {

		err2 := json.Unmarshal(buffer, &{{.Name}})

		if err2 != nil {
			log.Printf("Cannot unmarshal embedded {{.Name}}.json.gz: %s", err2.Error())
		}

	} else {
		log.Printf("Cannot decompress embedded {{.Name}}.json.gz: %s", err1.Error())
	}

}
