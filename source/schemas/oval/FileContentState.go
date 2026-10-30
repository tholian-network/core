package oval

import "encoding/xml"

type FileContentState struct {
	XMLName    xml.Name `xml:"textfilecontent54_state"`
	Identifier string   `xml:"id,attr"`
	Comment    string   `xml:"comment,attr"`
	Version    string   `xml:"version,attr"`

	Comparison struct {
		Value     string `xml:",chardata"`
		DataType  string `xml:"datatype,attr"`
		Operation string `xml:"operation,attr"`
	} `xml:"subexpression"`
}
