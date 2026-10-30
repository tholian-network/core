package oval

import "encoding/xml"

type Generator struct {
	XMLName        xml.Name `xml:"generator"`
	ProductName    string   `xml:"oval:product_name"`
	ProductVersion string   `xml:"oval:product_version"`
	SchemaVersion  string   `xml:"oval:schema_version"`
	Timestamp      string   `xml:"oval:timestamp"`
	ExtensionPoint string   `xml:"oval:extension_point"`
}
