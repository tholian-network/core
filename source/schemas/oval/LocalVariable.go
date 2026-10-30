package oval

import "encoding/xml"

type ArithmeticValue struct {
	XMLName   xml.Name `xml:"concat"`
	Operation string   `xml:"arithmetic_operation,attr"`

	Literal struct {
		Value    string `xml:",chardata"`
		DataType string `xml:"datatype,attr"`
	} `xml:"literal_component"`

	Object struct {
		Item      string `xml:"item_field,attr"`
		Reference string `xml:"object_ref,attr"`
	} `xml:"object_component"`
}

type ConcatValue struct {
	XMLName xml.Name `xml:"concat"`

	Literal struct {
		Value    string `xml:",chardata"`
		DataType string `xml:"datatype,attr"`
	} `xml:"literal_component"`

	Pattern struct {
		Value  string `xml:"pattern,attr"`
		Object struct {
			Item      string `xml:"item_field,attr"`
			Reference string `xml:"object_ref,attr"`
		} `xml:"object_component"`
	} `xml:"regexp_capture"`
}

type LocalVariable struct {
	XMLName    xml.Name `xml:"local_variable"`
	Identifier string   `xml:"id,attr"`
	Comment    string   `xml:"comment,attr"`
	Datatype   string   `xml:"datatype,attr"`
	Version    string   `xml:"version,attr"`

	Values struct {
		Arithmetic []ArithmeticValue `xml:"arithmetic"`
		Concat     []ConcatValue     `xml:"concat"`
	}
}
