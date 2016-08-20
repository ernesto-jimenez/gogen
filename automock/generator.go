package automock

import (
	"bufio"
	"bytes"
	"fmt"
	"go/types"
	"io"
	"text/template"

	"github.com/ernesto-jimenez/gogen/cleanimports"
	"github.com/ernesto-jimenez/gogen/importer"
	"github.com/ernesto-jimenez/gogen/imports"
)

// Generator produces code to mock an interface
type Generator struct {
	name      string
	ifaceName string
	namePkg   string
	inPkg     bool
	pkg       *types.Package
	iface     *types.Interface
	mockTmpl  *template.Template
}

// NewGenerator initializes a Generator that will mock the given interface from the specified package.
func NewGenerator(pkg, iface string) (*Generator, error) {
	p, err := importer.DefaultWithTestFiles().Import(pkg)
	if err != nil {
		return nil, err
	}
	obj := p.Scope().Lookup(iface)
	if obj == nil {
		return nil, fmt.Errorf("interface %s missing", iface)
	}
	if !types.IsInterface(obj.Type()) {
		return nil, fmt.Errorf("%s should be an interface, was %s", iface, obj.Type())
	}
	g := &Generator{
		ifaceName: iface,
		pkg:       p,
		iface:     obj.Type().Underlying().(*types.Interface).Complete(),
	}
	g.SetTemplate(defaultMockTemplate)
	return g, nil
}

// Methods returns information about all the methods required to satisfy the interface
func (g Generator) Methods() []Method {
	numMethods := g.iface.NumMethods()
	methods := make([]Method, numMethods)
	for i := 0; i < numMethods; i++ {
		methods[i] = Method{&g, g.iface.Method(i)}
	}
	return methods
}

func (g Generator) qf(pkg *types.Package) string {
	if g.inPkg && g.pkg == pkg {
		return ""
	}
	return pkg.Name()
}

// Name returns the mock type's name by default it is {interfaceName}Mock
func (g Generator) Name() string {
	if g.name != "" {
		return g.name
	}
	return g.ifaceName + "Mock"
}

// SetName changes the mock type's name
func (g *Generator) SetName(name string) {
	g.name = name
}

// Package returns the name of the package containing the mock
func (g Generator) Package() string {
	if g.namePkg != "" {
		return g.namePkg
	}
	if g.inPkg {
		return g.pkg.Name()
	}
	return "mocks"
}

// SetPackage changes the package containing the mock
func (g *Generator) SetPackage(name string) {
	g.namePkg = name
}

func (g *Generator) SetInternal(inPkg bool) {
	g.inPkg = inPkg
}

// Imports returns all the packages that have to be imported for the
func (g Generator) Imports() map[string]string {
	imports := imports.New(g.Package())
	for _, m := range g.Methods() {
		s := m.signature()
		imports.AddImportsFrom(s.Params())
		imports.AddImportsFrom(s.Results())
	}
	return imports.Imports()
}

// SetTemplate allows defining a different template to generate the mock. It will be parsed with text/template and execuded with the Generator.
func (g *Generator) SetTemplate(tmpl string) error {
	t, err := template.New("mock").Parse(tmpl)
	if err != nil {
		return err
	}
	g.mockTmpl = t
	return nil
}

// Write writes the generated code in the io.Writer
func (g Generator) Write(wr io.Writer) error {
	var buf bytes.Buffer
	if err := g.mockTmpl.Execute(&buf, g); err != nil {
		return err
	}
	err := cleanimports.Clean(wr, buf.Bytes())
	if err != nil {
		err = GenerationError{
			Err:  err,
			Code: buf.Bytes(),
		}
	}
	return err
}

// GenerationError is returned by Write when an error is encountered
type GenerationError struct {
	Err  error
	Code []byte
}

func (err GenerationError) Error() string {
	return err.Err.Error()
}

// CodeWithLineNumbers returns all the code including line numbers
func (err GenerationError) CodeWithLineNumbers() string {
	var buf bytes.Buffer
	scanner := bufio.NewScanner(bytes.NewReader(err.Code))
	var i int
	for scanner.Scan() {
		i++
		fmt.Fprintf(&buf, "%d: %s\n", i, scanner.Text())
	}
	return buf.String()
}

var (
	defaultMockTemplate = `/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/gogen/automock
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

package {{.Package}}

import (
	"fmt"
	mock "github.com/stretchr/testify/mock"
{{range $path, $name := .Imports}}
	{{$name}} "{{$path}}"{{end}}
)

// {{.Name}} mock
type {{.Name}} struct {
	mock.Mock
}

{{$gen := .}}
{{range .Methods}}
// {{.Name}} mocked method
func (m *{{$gen.Name}}) {{.Name}}({{range $index, $type := .ParamTypes}}{{if $index}}, {{end}}p{{$index}} {{$type}}{{end}}) ({{range $index, $type := .ReturnTypes}}{{if $index}}, {{end}}{{$type}}{{end}}) {
{{if .ReturnTypes}}
	ret := m.Called({{range $index, $type := .ParamTypes}}{{if $index}}, {{end}}p{{$index}}{{end}})
	{{range $index, $type := .ReturnTypes}}
	var r{{$index}} {{$type}}
	switch res := ret.Get({{$index}}).(type) {
	case nil:
	case {{$type}}:
		r{{$index}} = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}
	{{end}}
	return {{range $index, $type := .ReturnTypes}}{{if $index}}, {{end}}r{{$index}}{{end}}
{{else}}
	m.Called({{range $index, $type := .ParamTypes}}{{if $index}}, {{end}}p{{$index}}{{end}})
{{end}}
}
{{end}}
`
)
