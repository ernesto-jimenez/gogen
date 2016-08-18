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

// Wrapped is a wrapper around ExportedDefaultExportedInterface.Wrapped
func Wrapped(something string) (io.Writer, error) {
	return ExportedDefaultExportedInterface.Wrapped(something)
}
