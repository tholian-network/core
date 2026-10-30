package oval

import "encoding/xml"

type VariableState struct {
	XMLName    xml.Name `xml:"variable_state"`
	Identifier string   `xml:"id,attr"`
	Comment    string   `xml:"comment,attr"`
	Version    string   `xml:"version,attr"`

	Comparison struct {
		Variable struct {
			Check     string `xml:"var_check,attr"`
			Reference string `xml:"var_ref,attr"`
		}
		DataType  string `xml:"datatype,attr"`
		Operation string `xml:"operation,attr"`
	} `xml:"value"`
}
