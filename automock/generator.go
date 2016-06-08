package automock

import (
	"bytes"
	"fmt"
	"go/types"
	"io"
	"text/template"

	"github.com/ernesto-jimenez/gogen/cleanimports"
	"github.com/ernesto-jimenez/gogen/importer"
	"github.com/ernesto-jimenez/gogen/imports"
)

// Generator interface
type Generator interface {
	Name() string
	SetName(string)
	Package() string
	SetPackage(string)
	SetInternal(bool)
	Methods() []Method
	Imports() map[string]string
	Write(io.Writer) error
}

type generator struct {
	name      string
	ifaceName string
	namePkg   string
	inPkg     bool
	pkg       *types.Package
	iface     *types.Interface
}

// NewGenerator initializes a generator
func NewGenerator(pkg, iface string) (Generator, error) {
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
	return &generator{
		ifaceName: iface,
		pkg:       p,
		iface:     obj.Type().Underlying().(*types.Interface).Complete(),
	}, nil
}

func (g generator) Methods() []Method {
	numMethods := g.iface.NumMethods()
	methods := make([]Method, numMethods)
	for i := 0; i < numMethods; i++ {
		methods[i] = Method{&g, g.iface.Method(i)}
	}
	return methods
}

func (g generator) qf(pkg *types.Package) string {
	if g.inPkg && g.pkg == pkg {
		return ""
	}
	return pkg.Name()
}

func (g generator) Name() string {
	if g.name != "" {
		return g.name
	}
	return g.ifaceName + "Mock"
}

func (g *generator) SetName(name string) {
	g.name = name
}

func (g generator) Package() string {
	if g.namePkg != "" {
		return g.namePkg
	}
	if g.inPkg {
		return g.pkg.Name()
	}
	return "mocks"
}

func (g *generator) SetPackage(name string) {
	g.namePkg = name
}

func (g *generator) SetInternal(inPkg bool) {
	g.inPkg = inPkg
}

func (g generator) Imports() map[string]string {
	imports := imports.New(g.Package())
	for _, m := range g.Methods() {
		s := m.signature()
		imports.AddImportsFrom(s.Params())
		imports.AddImportsFrom(s.Results())
	}
	return imports.Imports()
}

func (g generator) Write(wr io.Writer) error {
	var buf bytes.Buffer
	if err := mockTmpl.Execute(&buf, g); err != nil {
		return err
	}
	return cleanimports.Clean(wr, buf.Bytes())
}

var (
	mockTmpl = template.Must(template.New("mock").Parse(`/*
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
`))
)
