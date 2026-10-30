package templates

import "bytes"
import _ "embed"
import "errors"
import "sort"
import "strings"
import "text/template"

//go:embed FilesystemCache.embed.tpl
var filesystem_cache_embed_tpl []byte

func FilesystemCache(config *Config) ([]byte, []byte, error) {

	if config.IsValid() {

		var buffer_insight bytes.Buffer
		var buffer_embed bytes.Buffer

		sort.Strings(config.Files)

		template_insight := template.Must(template.New("insight.go").Funcs(template.FuncMap{
			"StringsJoin": strings.Join,
		}).Parse(string(cache_insight_tpl)))

		template_embed := template.Must(template.New("embed.go").Funcs(template.FuncMap{
			"StringsJoin": strings.Join,
		}).Parse(string(filesystem_cache_embed_tpl)))

		err1 := template_insight.Execute(&buffer_insight, config)

		if err1 != nil {
			return []byte{}, []byte{}, err1
		}

		err2 := template_embed.Execute(&buffer_embed, config)

		if err2 != nil {
			return []byte{}, []byte{}, err2
		}

		return buffer_insight.Bytes(), buffer_embed.Bytes(), nil

	} else {
		return []byte{}, []byte{}, errors.New("Template: Invalid Config")
	}

}
