package {{.Package}}

import "tholian-core/caches"

var {{.Name}} *caches.{{.Name}} = caches.New{{.Name}}()
