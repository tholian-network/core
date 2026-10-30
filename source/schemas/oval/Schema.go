package oval

import "encoding/xml"

type Schema struct {
	XMLName     xml.Name     `xml:"oval_definitions"`
	Generator   Generator    `xml:"generator"`
	Definitions []Definition `xml:"definitions>definition"`

	Objects struct {
		File        []FileObject        `xml:"file_object"`
		FileContent []FileContentObject `xml:"textfilecontent54_object"`
		Uname       []UnameObject       `xml:"uname_object"`
		Variable    []VariableObject    `xml:"variable_object"`
	} `xml:"objects"`

	States struct {
		File        []FileState        `xml:"file_state"`
		FileContent []FileContentState `xml:"textfilecontent54_state"`
		Uname       []UnameState       `xml:"uname_state"`
		Variable    []VariableState    `xml:"variable_state"`
	} `xml:"states"`

	Tests struct {
		File        []FileTest        `xml:"file_test"`
		FileContent []FileContentTest `xml:"textfilecontent54_test"`
		Uname       []UnameTest       `xml:"uname_test"`
		Variable    []VariableTest    `xml:"variable_test"`
	} `xml:"tests"`

	Variables struct {
		Constant []ConstantVariable `xml:"constant_variable"`
		Local    []LocalVariable    `xml:"local_variable"`
	} `xml:"variables"`
}
