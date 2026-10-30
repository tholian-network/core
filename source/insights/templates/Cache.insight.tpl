package {{.Package}}

import "tholian-endpoint/caches"

var {{.Name}} *caches.{{.Name}} = caches.New{{.Name}}()
