package oval

import "encoding/xml"

type ExtendDefinition struct {
	XMLName            xml.Name `xml:"extend_definition"`
	Definition         string   `xml:"definition_ref,attr"`
	Negate             bool     `xml:"negate,attr"`
	Comment            string   `xml:"comment,attr"`
	ApplicabilityCheck bool     `xml:"applicability_check"`
}
