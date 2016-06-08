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

// EDUIEmbedded is a wrapper around ExportedDefaultUnexportedInterface.Embedded
func EDUIEmbedded() {
	ExportedDefaultUnexportedInterface.Embedded()
}

// EDUIWrapped is a wrapper around ExportedDefaultUnexportedInterface.Wrapped
func EDUIWrapped(something string) (io.Writer, error) {
	return ExportedDefaultUnexportedInterface.Wrapped(something)
}

// EDUIWrappedVariadric is a wrapper around ExportedDefaultUnexportedInterface.WrappedVariadric
func EDUIWrappedVariadric(something ...string) error {
	return ExportedDefaultUnexportedInterface.WrappedVariadric(something...)
}
