package oval

import "bytes"
import "encoding/xml"

type Criterion struct {
	XMLName            xml.Name `xml:"criterion"`
	Negate             bool     `xml:"negate,attr"`
	Comment            string   `xml:"comment,attr"`
	Reference          string   `xml:"test_ref,attr"`
	ApplicabilityCheck bool     `xml:"applicability_check"`
}

func (criterion *Criterion) ToLog() string {

	var buffer bytes.Buffer

	buffer.WriteString("Criterion({\n")
	buffer.WriteString("    Comment:   \"" + criterion.Comment + "\",\n")
	buffer.WriteString("    Reference: \"" + criterion.Reference + "\"\n")
	buffer.WriteString("})")

	return buffer.String()

}
