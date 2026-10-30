package oval

import "encoding/xml"

type Reference struct {
	XMLName    xml.Name `xml:"reference"`
	Source     string   `xml:"source,attr"`
	Identifier string   `xml:"ref_id,attr"`
	URL        string   `xml:"ref_url,attr"`
}
