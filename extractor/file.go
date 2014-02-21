package extractor

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func FileImportCalls(file string) (Imports, error) {
	v, err := newVisitor(file)
	if err != nil {
		return nil, err
	}

	v.walk()

	return v.Imports, nil
}

type visitor struct {
	Imports Imports
	fileAst *ast.File
	fset    *token.FileSet
}

func newVisitor(file string) (*visitor, error) {
	fset := token.NewFileSet()
	fileAst, err := parser.ParseFile(fset, file, nil, parser.ParseComments)

	if err != nil {
		return nil, err
	}

	v := visitor{
		Imports: Imports{},
		fileAst: fileAst,
		fset:    fset,
	}

	return &v, nil
}

func (v *visitor) walk() *visitor {
	ast.Walk(v, v.fileAst) // calls the Visit method for each ast node
	return v
}

func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {
	switch t := node.(type) {
	// adding each imported package to imports
	case *ast.GenDecl:
		if t.Tok == token.IMPORT {
			for _, spec := range t.Specs {
				if f, ok := spec.(*ast.ImportSpec); ok {
					// i'm not sure why, but f.Path.Value returns the value in "doublequotes"
					// so let's filter them out the hard way
					path := strings.Replace(f.Path.Value, "\"", "", -1)

					v.Imports = append(v.Imports, path)
				}
			}
		}
	}

	return v
}
