package oval

import "encoding/xml"

type VariableObject struct {
	XMLName    xml.Name `xml:"variable_object"`
	Identifier string   `xml:"id,attr"`
	Comment    string   `xml:"comment,attr"`
	Version    string   `xml:"version,attr"`

	Variable struct {
		Reference string `xml:",chardata"`
	} `xml:"var_ref"`
}
