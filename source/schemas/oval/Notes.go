package oval

import "encoding/xml"

type Notes struct {
	XMLName xml.Name `xml:"notes"`
	Notes   []string `xml:"note"`
}
