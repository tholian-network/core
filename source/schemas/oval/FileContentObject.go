package oval

import "encoding/xml"

type FileContentObject struct {
	XMLName    xml.Name `xml:"textfilecontent54_object"`
	Identifier string   `xml:"id,attr"`
	Comment    string   `xml:"comment,attr"`
	Version    string   `xml:"version,attr"`

	Behaviours string `xml:"behaviours"`
	Filename   string `xml:"filename"`

	Folder struct {
		Reference string `xml:",chardata"`
	} `xml:"path"`

	File struct {
		Reference string `xml:",chardata"`
	} `xml:"filepath"`

	Pattern struct {
		Value     string `xml:",chardata"`
		Operation string `xml:"operation,attr"`
	} `xml:"pattern"`

	Instance struct {
		Value    string `xml:",chardata"`
		DataType string `xml:"datatype,attr"`
	} `xml:"instance"`
}
