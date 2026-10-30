package oval

import "encoding/xml"

type Metadata struct {
	XMLName     xml.Name    `xml:"metadata"`
	Title       string      `xml:"title"`
	Affected    Affected    `xml:"affected"`
	References  []Reference `xml:"reference"`
	Description string      `xml:"description"`
}
