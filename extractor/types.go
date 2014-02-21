package extractor

import (
	"encoding/xml"
)

type Imports []string

type xmlImports struct {
	XMLName xml.Name `xml:"go-import-extractor"`
	Imports Imports  `xml:"imports>import"`
}

func uniqueAppend(slice []string, element string) []string {
	for _, v := range slice {
		if v == element {
			return slice
		}
	}

	return append(slice, element)
}
