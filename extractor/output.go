package extractor

import (
	"encoding/xml"
	"io"
)

func PrintXml(imports Imports, output io.Writer) error {
	enc := xml.NewEncoder(output)
	enc.Indent("  ", "    ")

	err := enc.Encode(xmlImports{Imports: imports})
	if err != nil {
		return err
	}

	return nil
}
