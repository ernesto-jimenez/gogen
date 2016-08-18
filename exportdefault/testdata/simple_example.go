/*
* CODE GENERATED AUTOMATICALLY WITH goexportdefault
* THIS FILE MUST NOT BE EDITED BY HAND
*
* Install goexportdefault with:
* go get github.com/ernesto-jimenez/gogen/cmd/goexportdefault
 */

package testpkg

import (
	io "io"
)

// Embedded is a wrapper around ExportedDefaultExportedInterface.Embedded
func Embedded() {
	ExportedDefaultExportedInterface.Embedded()
}

// Wrapped is a wrapper around ExportedDefaultExportedInterface.Wrapped
func Wrapped(something string) (io.Writer, error) {
	return ExportedDefaultExportedInterface.Wrapped(something)
}

// WrappedVariadric is a wrapper around ExportedDefaultExportedInterface.WrappedVariadric
func WrappedVariadric(something ...string) error {
	return ExportedDefaultExportedInterface.WrappedVariadric(something...)
}
