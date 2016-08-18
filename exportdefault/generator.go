// Package exportdefault provides the functionality to automatically generate
// package-level exported functions wrapping calls to a package-level default
// instance of a type.
//
// This helps auto-generating code for the common use case where a package
// implements certain information as methods within a stub and, for
// convenience, exports functions that wrap calls to those methods on a default
// variable.
//
// Some examples of that behaviour in the stdlib:
//
//  - `net/http` has `http.DefaultClient` and functions like `http.Get` just
//     call the default `http.DefaultClient.Get`
//  - `log` has `log.Logger` and functions like `log.Print` just call the
//     default `log.std.Print`
package exportdefault

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/doc"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io"
	"io/ioutil"
	"path"
	"regexp"
	"text/template"

	"github.com/ernesto-jimenez/gogen/cleanimports"
	"github.com/ernesto-jimenez/gogen/imports"
)

// Generator contains the metadata needed to generate all the function wrappers
// arround methods from a package variable
type Generator struct {
	Name           string
	Imports        map[string]string
	funcs          []fn
	FuncNamePrefix string
	Include        *regexp.Regexp
	Exclude        *regexp.Regexp
}

// New initialises a new Generator for the corresponding package's variable
//
// Returns an error if the package or variable are invalid
func New(pkg string, variable string) (*Generator, error) {
	scope, docs, err := parsePackageSource(pkg)
	if err != nil {
		return nil, err
	}

	importer, funcs, err := analyzeCode(scope, docs, variable)
	if err != nil {
		return nil, err
	}

	return &Generator{
		Name:    docs.Name,
		Imports: importer.Imports(),
		funcs:   funcs,
	}, nil
}

// Write the generated code into the given io.Writer
//
// Returns an error if there is a problem generating the code
func (g *Generator) Write(w io.Writer) error {
	buff := bytes.NewBuffer(nil)

	// Generate header
	if err := headerTpl.Execute(buff, g); err != nil {
		return err
	}

	// Generate funcs
	for _, fn := range g.funcs {
		if g.Include != nil && !g.Include.MatchString(fn.Name) {
			continue
		}
		if g.Exclude != nil && g.Exclude.MatchString(fn.Name) {
			continue
		}
		fn.FuncNamePrefix = g.FuncNamePrefix
		buff.Write([]byte("\n\n"))
		if err := funcTpl.Execute(buff, &fn); err != nil {
			return err
		}
	}

	return cleanimports.Clean(w, buff.Bytes())
}

type fn struct {
	FuncNamePrefix string
	WrappedVar     string
	Name           string
	CurrentPkg     string
	TypeInfo       *types.Func
}

func (f *fn) Qualifier(p *types.Package) string {
	if p == nil || p.Name() == f.CurrentPkg {
		return ""
	}
	return p.Name()
}

func (f *fn) Params() string {
	sig := f.TypeInfo.Type().(*types.Signature)
	params := sig.Params()
	p := ""
	comma := ""
	to := params.Len()
	var i int

	if sig.Variadic() {
		to--
	}
	for i = 0; i < to; i++ {
		param := params.At(i)
		name := param.Name()
		if name == "" {
			name = fmt.Sprintf("p%d", i)
		}
		p += fmt.Sprintf("%s%s %s", comma, name, types.TypeString(param.Type(), f.Qualifier))
		comma = ", "
	}
	if sig.Variadic() {
		param := params.At(params.Len() - 1)
		name := param.Name()
		if name == "" {
			name = fmt.Sprintf("p%d", to)
		}
		p += fmt.Sprintf("%s%s ...%s", comma, name, types.TypeString(param.Type().(*types.Slice).Elem(), f.Qualifier))
	}
	return p
}

func (f *fn) ReturnsAnything() bool {
	sig := f.TypeInfo.Type().(*types.Signature)
	params := sig.Results()
	return params.Len() > 0
}

func (f *fn) ReturnTypes() string {
	sig := f.TypeInfo.Type().(*types.Signature)
	params := sig.Results()
	p := ""
	comma := ""
	to := params.Len()
	var i int

	for i = 0; i < to; i++ {
		param := params.At(i)
		p += fmt.Sprintf("%s %s", comma, types.TypeString(param.Type(), f.Qualifier))
		comma = ", "
	}
	if to > 1 {
		p = fmt.Sprintf("(%s)", p)
	}
	return p
}

func (f *fn) ForwardedParams() string {
	sig := f.TypeInfo.Type().(*types.Signature)
	params := sig.Params()
	p := ""
	comma := ""
	to := params.Len()
	var i int

	if sig.Variadic() {
		to--
	}
	for i = 0; i < to; i++ {
		param := params.At(i)
		name := param.Name()
		if name == "" {
			name = fmt.Sprintf("p%d", i)
		}
		p += fmt.Sprintf("%s%s", comma, name)
		comma = ", "
	}
	if sig.Variadic() {
		param := params.At(params.Len() - 1)
		name := param.Name()
		if name == "" {
			name = fmt.Sprintf("p%d", to)
		}
		p += fmt.Sprintf("%s%s...", comma, name)
	}
	return p
}

// parsePackageSource returns the types scope and the package documentation from the specified package
func parsePackageSource(pkg string) (*types.Scope, *doc.Package, error) {
	pd, err := build.Import(pkg, ".", 0)
	if err != nil {
		return nil, nil, err
	}

	fset := token.NewFileSet()
	files := make(map[string]*ast.File)
	fileList := make([]*ast.File, len(pd.GoFiles))
	for i, fname := range pd.GoFiles {
		src, err := ioutil.ReadFile(path.Join(pd.SrcRoot, pd.ImportPath, fname))
		if err != nil {
			return nil, nil, err
		}
		f, err := parser.ParseFile(fset, fname, src, parser.ParseComments|parser.AllErrors)
		if err != nil {
			return nil, nil, err
		}
		files[fname] = f
		fileList[i] = f
	}

	cfg := types.Config{
		Importer: importer.Default(),
	}
	info := types.Info{
		Defs: make(map[*ast.Ident]types.Object),
	}
	tp, err := cfg.Check(pkg, fset, fileList, &info)
	if err != nil {
		return nil, nil, err
	}

	scope := tp.Scope()

	ap, _ := ast.NewPackage(fset, files, nil, nil)
	docs := doc.New(ap, pkg, doc.AllDecls|doc.AllMethods)

	return scope, docs, nil
}

func analyzeCode(scope *types.Scope, docs *doc.Package, variable string) (imports.Importer, []fn, error) {
	pkg := docs.Name
	v, ok := scope.Lookup(variable).(*types.Var)
	if v == nil {
		return nil, nil, fmt.Errorf("impossible to find variable %s", variable)
	}
	if !ok {
		return nil, nil, fmt.Errorf("%s must be a variable", variable)
	}
	var vType interface {
		NumMethods() int
		Method(int) *types.Func
	}
	switch t := v.Type().(type) {
	case *types.Interface:
		vType = t
	case *types.Pointer:
		vType = t.Elem().(*types.Named)
	case *types.Named:
		vType = t
		if t, ok := t.Underlying().(*types.Interface); ok {
			vType = t
		}
	default:
		return nil, nil, fmt.Errorf("variable is of an invalid type: %T", v.Type().Underlying())
	}

	importer := imports.New(pkg)
	var funcs []fn
	for i := 0; i < vType.NumMethods(); i++ {
		f := vType.Method(i)

		if !f.Exported() {
			continue
		}

		sig := f.Type().(*types.Signature)

		funcs = append(funcs, fn{
			WrappedVar: variable,
			Name:       f.Name(),
			CurrentPkg: pkg,
			TypeInfo:   f,
		})
		importer.AddImportsFrom(sig.Params())
		importer.AddImportsFrom(sig.Results())
	}
	return importer, funcs, nil
}

var headerTpl = template.Must(template.New("header").Parse(`/*
* CODE GENERATED AUTOMATICALLY WITH goexportdefault
* THIS FILE MUST NOT BE EDITED BY HAND
*
* Install goexportdefault with:
* go get github.com/ernesto-jimenez/gogen/cmd/goexportdefault
*/

package {{.Name}}

import (
{{range $path, $name := .Imports}}
	{{$name}} "{{$path}}"{{end}}
)
`))

var funcTpl = template.Must(template.New("func").Parse(`// {{.FuncNamePrefix}}{{.Name}} is a wrapper around {{.WrappedVar}}.{{.Name}}
func {{.FuncNamePrefix}}{{.Name}}({{.Params}}) {{.ReturnTypes}} {
	{{if .ReturnsAnything}}return {{end}}{{.WrappedVar}}.{{.Name}}({{.ForwardedParams}})
}`))
