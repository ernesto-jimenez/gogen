package imports

import (
	"go/types"
	"strings"
)

type Importer interface {
	AddImportsFrom(t types.Type)
	Imports() map[string]string
}

// imports contains metadata about all the imports from a given package
type imports struct {
	currentpkg string
	imp        map[string]string
}

// AddImportsFrom adds imports used in the passed type
func (imp *imports) AddImportsFrom(t types.Type) {
	switch el := t.(type) {
	case *types.Basic:
	case *types.Slice:
		imp.AddImportsFrom(el.Elem())
	case *types.Pointer:
		imp.AddImportsFrom(el.Elem())
	case *types.Named:
		pkg := el.Obj().Pkg()
		if pkg == nil {
			return
		}
		if pkg.Name() == imp.currentpkg {
			return
		}
		imp.imp[vendorlessImportPath(pkg.Path())] = pkg.Name()
	case *types.Tuple:
		for i := 0; i < el.Len(); i++ {
			imp.AddImportsFrom(el.At(i).Type())
		}
	default:
	}
}

// vendorlessImportPath returns the devendorized version of the provided import path.
// e.g. "foo/bar/vendor/a/b" => "a/b"
func vendorlessImportPath(ipath string) string {
	// Devendorize for use in import statement.
	if i := strings.LastIndex(ipath, "/vendor/"); i >= 0 {
		return ipath[i+len("/vendor/"):]
	}
	if strings.HasPrefix(ipath, "vendor/") {
		return ipath[len("vendor/"):]
	}
	return ipath
}

// AddImportsFrom adds imports used in the passed type
func (imp *imports) Imports() map[string]string {
	return imp.imp
}

// New initializes a new structure to track packages imported by the currentpkg
func New(currentpkg string) Importer {
	return &imports{
		currentpkg: currentpkg,
		imp:        make(map[string]string),
	}
}
