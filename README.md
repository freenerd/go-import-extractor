# go-import-extractor

go-import-extractor helps to analyze go programs. It extracts a list of imported packages from either a .go file or a go package. On a go package, it may recursively analyze the imported packages.

## Install

    go get github.com/freenerd/go-import-extractor

## Use

For basic usage give the absolute go source file path as first argument. In this example, we analyze the code of this package:

    go-import-extractor $GOPATH/src/github.com/freenerd/go-import-extractor/main.go

The output goes to STDOUT and is formated as xml. It looks like this:

```xml
<imports>
  <import>fmt</import>
  <import>github.com/freenerd/go-import-extractor/extractor</import>
</imports>
```

To analyze a whole package, call like this:

    go-import-extractor -p github.com/freenerd/go-import-extractor

Output can be filtered by specific suspect package calls (TODO: make customizable)

    go-import-extractor -p github.com/freenerd/go-import-extractor -s

## Output formats

XML

## Limitations

- does not expose package renames, but only uses the original package import name

## Todo
