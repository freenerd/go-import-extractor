package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/freenerd/go-import-extractor/extractor"
)

var (
	file = flag.String("f", "", "a go source file to extract imports from")
	pkg  = flag.String("p", "", "a go package to extract imports from")
	format = flag.String("format", "xml", "format of output. valid values: xml (default)")
)

func main() {
	flag.Parse()

	if *file == "" && *pkg == "" {
		log.Printf("need go source file path or go package to continue\n\n")
		flag.Usage()
		return
	}

	if *file != "" && *pkg != "" {
		log.Printf("please only specify either a source file or a package, not both\n\n")
		flag.Usage()
		return
	}

	imports := extractor.Imports{}
	var err error

	if *file != "" {
		imports, err = extractor.FileImportCalls(*file)
	} else if *pkg != "" {
		imports, err = extractor.PackageImportCalls(*pkg)
	}
	if err != nil {
		log.Fatal(err)
		return
	}

	format := strings.ToLower(*format)
	if format == "xml" {
		extractor.PrintXml(imports, os.Stdout)
	} else {
		log.Printf("unknown format %s\n\n", format)
		flag.Usage()
		return
	}
}
