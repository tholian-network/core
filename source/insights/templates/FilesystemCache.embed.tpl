{{if .BuildTag}}//go:build {{.BuildTag}}

{{end}}package {{.Package}}

import "embed"
{{if .Files}}import "tholian-core/structs"
import "bytes"
import "compress/gzip"
import "encoding/json"
import "io"
import "log"
import "strings"
{{end}}
{{if .Files}}//go:embed {{StringsJoin .Files " "}}
{{end}}var embedded_{{.Name}} embed.FS

func init() {
{{if .Files}}
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


	files, err1 := embedded_{{.Name}}.ReadDir("{{.Folder}}")

	if err1 == nil {

		for _, file := range files {

			var err error = nil

			if file.IsDir() {

				subfiles, err2 := embedded_{{.Name}}.ReadDir("{{.Folder}}/"+file.Name())

				if err2 == nil {

					for _, subfile := range subfiles {

						var suberr error = nil

						if !subfile.IsDir() {

							tmp, err3 := embedded_{{.Name}}.ReadFile("{{.Folder}}/"+file.Name()+"/"+subfile.Name())

							if err3 == nil {

								if strings.HasSuffix(subfile.Name(), ".json.gz") {

									buffer, err4 := decompress(tmp)

									if err4 == nil {

										var instance {{.Struct}}

										err5 := json.Unmarshal(buffer, &instance)

										if err5 == nil {
											{{.Name}}.Add(instance)
										} else {
											log.Printf("Cannot unmarshal embedded file {{.Folder}}/"+file.Name()+"/"+subfile.Name()+" for {{.Name}}: %s", err5.Error())
										}

									} else {
										log.Printf("Cannot decompress embedded file {{.Folder}}/"+file.Name()+"/"+subfile.Name()+" for {{.Name}}: %s", err4.Error())
									}

								} else if strings.HasSuffix(subfile.Name(), ".json") {

									var instance {{.Struct}}

									err4 := json.Unmarshal(tmp, &instance)

									if err4 == nil {
										{{.Name}}.Add(instance)
									} else {
										log.Printf("Cannot unmarshal embedded file {{.Folder}}/"+file.Name()+"/"+subfile.Name()+" for {{.Name}}: %s", err4.Error())
									}

								}

							} else {
								log.Printf("Cannot read embedded file {{.Folder}}/"+file.Name()+"/"+subfile.Name()+" for {{.Name}}: %s", err3.Error())
							}

						}

					}

				} else {
					log.Printf("Cannot read embedded filesystem at {{.Folder}}/"+file.Name()+" for {{.Name}}: %s", err2.Error())
				}

			} else {

				tmp, err2 := embedded_{{.Name}}.ReadFile("{{.Folder}}/"+file.Name())

				if err2 == nil {

					if strings.HasSuffix(file.Name(), ".json.gz") {

						buffer, err3 := decompress(tmp)

						if err3 == nil {

							var instance {{.Struct}}

							err4 := json.Unmarshal(buffer, &instance)

							if err4 == nil {
								{{.Name}}.Add(instance)
							} else {
								log.Printf("Cannot unmarshal embedded file {{.Folder}}/"+file.Name()+" for {{.Name}}: %s", err4.Error())
							}

						} else {
							log.Printf("Cannot decompress embedded file {{.Folder}}/"+file.Name()+" for {{.Name}}: %s", err3.Error())
						}

					} else if strings.HasSuffix(file.Name(), ".json") {

						var instance {{.Struct}}

						err3 := json.Unmarshal(tmp, &instance)

						if err3 == nil {
							{{.Name}}.Add(instance)
						} else {
							log.Printf("Cannot unmarshal embedded file {{.Folder}}/"+file.Name()+" for {{.Name}}: %s", err3.Error())
						}

					}

				} else {
					log.Printf("Cannot read embedded file {{.Folder}}/"+file.Name()+" for {{.Name}}: %s", err2.Error())
				}

			}

		}

	} else {
		log.Printf("Cannot read embedded filesystem at {{.Folder}} for {{.Name}}: %s", err1.Error())
	}
{{end}}
}
