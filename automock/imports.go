package automock

import "go/types"

// Imports contains metadata about all the imports from a given package
type Imports struct {
	gen *generator
	imp map[string]string
}

func (imports *Imports) fillImports(t types.Type) {
	switch el := t.(type) {
	case *types.Basic:
	case *types.Slice:
		imports.fillImports(el.Elem())
	case *types.Pointer:
		imports.fillImports(el.Elem())
	case *types.Named:
		pkg := el.Obj().Pkg()
		if pkg == nil {
			return
		}
		if imports.gen.inPkg && pkg == imports.gen.pkg {
			return
		}
		imports.imp[pkg.Path()] = pkg.Name()
	default:
	}
}

func (imports *Imports) importsFromParams(t *types.Tuple) {
	for i := 0; i < t.Len(); i++ {
		imports.fillImports(t.At(i).Type())
	}
}

func (imports *Imports) init(gen *generator) {
	imports.gen = gen
	imports.imp = make(map[string]string)
}
