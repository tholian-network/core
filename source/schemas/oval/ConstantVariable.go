package oval

import "encoding/xml"

type ConstantVariable struct {
	XMLName    xml.Name `xml:"constant_variable"`
	Identifier string   `xml:"id,attr"`
	Comment    string   `xml:"comment,attr"`
	Datatype   string   `xml:"datatype,attr"`
	Version    string   `xml:"version,attr"`
	Values     []string `xml:"value"`
}
