package oval

import "encoding/xml"
import "bytes"
import "strings"

type Criteria struct {
	XMLName            xml.Name           `xml:"criteria"`
	Operator           string             `xml:"operator,attr"`
	Negate             bool               `xml:"negate,attr"`
	Criteria           []Criteria         `xml:"criteria"`
	Criterion          []Criterion        `xml:"criterion"`
	ExtendDefinitionn  []ExtendDefinition `xml:"extend_definition"`
	ApplicabilityCheck bool               `xml:"applicability_check"`
}

func (criteria *Criteria) HasCriteria() bool {

	var result bool = false

	if len(criteria.Criteria) > 0 {
		result = true
	}

	return result

}

func (criteria *Criteria) HasCriterion() bool {

	var result bool = false

	if len(criteria.Criterion) > 0 {
		result = true
	}

	return result

}

func (criteria *Criteria) IsAND() bool {

	var result bool = false

	if criteria.Operator == "AND" {
		result = true
	}

	return result

}

func (criteria *Criteria) IsOR() bool {

	var result bool = false

	if criteria.Operator == "OR" {
		result = true
	}

	return result

}

func (criteria *Criteria) ToCriterions() []Criterion {

	var filtered []Criterion

	if criteria.HasCriteria() {

		for c1 := 0; c1 < len(criteria.Criteria); c1++ {

			criterions := criteria.Criteria[c1].ToCriterions()

			for c2 := 0; c2 < len(criterions); c2++ {
				filtered = append(filtered, criterions[c2])
			}

		}

	}

	if criteria.HasCriterion() {

		for c := 0; c < len(criteria.Criterion); c++ {
			filtered = append(filtered, criteria.Criterion[c])
		}

	}

	return filtered

}

func (criteria *Criteria) ToLog() string {

	var buffer bytes.Buffer

	buffer.WriteString("Criteria({\n")
	buffer.WriteString("    Operator: \"" + criteria.Operator + "\",\n")
	buffer.WriteString("    Criteria: [\n")

	for c := 0; c < len(criteria.Criteria); c++ {

		var chunk = criteria.Criteria[c].ToLog()

		if c < len(criteria.Criteria)-1 {
			chunk += ","
		}

		var lines = strings.Split(chunk, "\n")

		for l := 0; l < len(lines); l++ {
			buffer.WriteString("        " + lines[l] + "\n")
		}

	}

	buffer.WriteString("    ],\n")
	buffer.WriteString("    Criterion: [\n")

	for c := 0; c < len(criteria.Criterion); c++ {

		var chunk = criteria.Criterion[c].ToLog()

		if c < len(criteria.Criterion)-1 {
			chunk += ","
		}

		var lines = strings.Split(chunk, "\n")

		for l := 0; l < len(lines); l++ {
			buffer.WriteString("        " + lines[l] + "\n")
		}

	}

	buffer.WriteString("    ]\n")
	buffer.WriteString("})")

	return buffer.String()

}
