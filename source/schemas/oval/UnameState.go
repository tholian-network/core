package oval

import "encoding/xml"

type UnameState struct {
	XMLName    xml.Name `xml:"uname_state"`
	Identifier string   `xml:"id,attr"`
	Comment    string   `xml:"comment,attr"`
	Version    string   `xml:"version,attr"`

	Comparison struct {
		Value     string `xml:",chardata"`
		Operation string `xml:"operation,attr"`
	} `xml:"os_release"`
}
