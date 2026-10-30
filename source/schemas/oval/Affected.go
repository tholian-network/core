package oval

import "encoding/xml"

type Affected struct {
	XMLName   xml.Name `xml:"affected"`
	Family    string   `xml:"family,attr"`
	Platforms []string `xml:"platform"`
	Products  []string `xml:"product"`
}
