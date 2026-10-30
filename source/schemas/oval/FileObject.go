package oval

import "encoding/xml"

type FileObject struct {
	XMLName    xml.Name `xml:"file_object"`
	Identifier string   `xml:"id,attr"`
	Comment    string   `xml:"comment,attr"`
	Version    string   `xml:"version,attr"`
	Filepath   string   `xml:"filepath"`
}
