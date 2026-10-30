package oval

import "encoding/xml"

type Definition struct {
	XMLName    xml.Name `xml:"definition"`
	Identifier string   `xml:"id,attr"`
	Version    string   `xml:"version,attr"`
	Class      string   `xml:"class,attr"`
	Deprecated bool     `xml:"deprecated,attr"`

	Metadata Metadata `xml:"metadata"`
	Criteria Criteria `xml:"criteria"`
	Notes    Notes    `xml:"notes"`
}
