package extractor

import (
	"fmt"
	"go/build"
	"os"
	"path"
)

func PackageImportCalls(pkgPath string) (Imports, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working dir: %s", err)
	}

	return processPackage(cwd, pkgPath, Imports{})
}

func processPackage(root, pkgPath string, imports Imports) (Imports, error) {
	// read package
	pkg, err := build.Import(pkgPath, root, 0)
	if err != nil {
		return nil, err
	}

	// Don't worry about dependencies for stdlib packages
	if pkg.Goroot {
		return imports, nil
	}

	// analyze each file in package, merge results
	for _, file := range pkg.GoFiles {
		fileImports, err := FileImportCalls(path.Join(pkg.Dir, file))
		if err != nil {
			return nil, fmt.Errorf("failed in file %s: %s", file, err)
		}

		for _, v := range fileImports {
			imports = uniqueAppend(imports, v)
		}
	}

	// recursively extract from each imported package
	for _, imp := range pkg.Imports {
		// TODO: Don't do already analyzed packages
		imports, err = processPackage(root, imp, imports)
		if err != nil {
			return nil, fmt.Errorf("failed to process package %s:", imp)
		}
	}

	return imports, nil
}
