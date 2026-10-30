package oval

import "encoding/xml"

type UnameObject struct {
	XMLName    xml.Name `xml:"uname_object"`
	Identifier string   `xml:"id,attr"`
	Comment    string   `xml:"comment,attr"`
	Version    string   `xml:"version,attr"`
}
