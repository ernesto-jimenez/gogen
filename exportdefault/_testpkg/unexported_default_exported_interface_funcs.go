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

// UDEIEmbedded is a wrapper around unexportedDefaultExportedInterface.Embedded
func UDEIEmbedded() {
	unexportedDefaultExportedInterface.Embedded()
}

// UDEIWrapped is a wrapper around unexportedDefaultExportedInterface.Wrapped
func UDEIWrapped(something string) (io.Writer, error) {
	return unexportedDefaultExportedInterface.Wrapped(something)
}

// UDEIWrappedVariadric is a wrapper around unexportedDefaultExportedInterface.WrappedVariadric
func UDEIWrappedVariadric(something ...string) error {
	return unexportedDefaultExportedInterface.WrappedVariadric(something...)
}
