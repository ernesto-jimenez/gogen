package testpkg

import (
	"io"
)

//go:generate go run ../../cmd/goexportdefault/main.go -prefix=EDUI ExportedDefaultUnexportedInterface

// ExportedDefaultunexportedInterface to be generated
var ExportedDefaultUnexportedInterface unexportedInterface = impl{}

//go:generate go run ../../cmd/goexportdefault/main.go -prefix=UDUI unexportedDefaultUnexportedInterface
var unexportedDefaultUnexportedInterface unexportedInterface = impl{}

type unexportedInterface interface {
	embeddedInterface
	Wrapped(something string) (io.Writer, error)
	WrappedVariadric(something ...string) error
}
