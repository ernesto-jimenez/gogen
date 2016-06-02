// Package cleanimports provides functionality to clean unused imports from the
// given source code
package cleanimports

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
)

// Clean writes the clean source to io.Writer. The source can be a io.Reader,
// string or []bytes
func Clean(w io.Writer, src interface{}) error {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "clean.go", src, parser.ParseComments)
	if err != nil {
		return err
	}
	// Clean unused imports
	imports := astutil.Imports(fset, file)
	for _, group := range imports {
		for _, imp := range group {
			path := strings.Trim(imp.Path.Value, `"`)
			if !astutil.UsesImport(file, path) {
				astutil.DeleteImport(fset, file, path)
			}
		}
	}
	ast.SortImports(fset, file)
	// Write formatted code without unused imports
	return format.Node(w, fset, file)
}
