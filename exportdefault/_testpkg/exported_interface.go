package testpkg

import (
	"io"
)

//go:generate go run ../../cmd/goexportdefault/main.go -prefix=EDEI ExportedDefaultExportedInterface

// ExportedDefaultExportedInterface to be generated
var ExportedDefaultExportedInterface ExportedInterface = impl{}

//go:generate go run ../../cmd/goexportdefault/main.go -prefix=UDEI unexportedDefaultExportedInterface
var unexportedDefaultExportedInterface ExportedInterface = impl{}

// ExportedInterface for tests
type ExportedInterface interface {
	embeddedInterface
	// Wrapped documentation goes here
	Wrapped(something string) (io.Writer, error)
	WrappedVariadric(something ...string) error
}

type impl struct {
}

func (impl) Wrapped(string) (io.Writer, error) { return nil, nil }
func (impl) WrappedVariadric(...string) error  { return nil }
func (impl) Embedded()                         {}
