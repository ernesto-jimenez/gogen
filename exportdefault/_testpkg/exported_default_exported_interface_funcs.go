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

// EDEIEmbedded is a wrapper around ExportedDefaultExportedInterface.Embedded
func EDEIEmbedded() {
	ExportedDefaultExportedInterface.Embedded()
}

// EDEIWrapped is a wrapper around ExportedDefaultExportedInterface.Wrapped
func EDEIWrapped(something string) (io.Writer, error) {
	return ExportedDefaultExportedInterface.Wrapped(something)
}

// EDEIWrappedVariadric is a wrapper around ExportedDefaultExportedInterface.WrappedVariadric
func EDEIWrappedVariadric(something ...string) error {
	return ExportedDefaultExportedInterface.WrappedVariadric(something...)
}
