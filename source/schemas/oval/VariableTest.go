package oval

import "encoding/xml"

type VariableTest struct {
	XMLName        xml.Name `xml:"variable_test"`
	Identifier     string   `xml:"id,attr"`
	Version        string   `xml:"version,attr"`
	CheckExistence string   `xml:"check_existence,attr"`
	Check          string   `xml:"check,attr"`
	StateOperator  string   `xml:"state_operator,attr"`
	Comment        string   `xml:"comment,attr"`
	Deprecated     bool     `xml:"deprecated,attr"`
	Notes          Notes    `xml:"notes"`
	Object         struct {
		Reference string `xml:"object_ref,attr"`
	} `xml:"object"`
	State struct {
		Reference string `xml:"state_ref,attr"`
	} `xml:"state"`
}
