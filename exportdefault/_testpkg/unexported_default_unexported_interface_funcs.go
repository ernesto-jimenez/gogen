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

// UDUIEmbedded is a wrapper around unexportedDefaultUnexportedInterface.Embedded
func UDUIEmbedded() {
	unexportedDefaultUnexportedInterface.Embedded()
}

// UDUIWrapped is a wrapper around unexportedDefaultUnexportedInterface.Wrapped
func UDUIWrapped(something string) (io.Writer, error) {
	return unexportedDefaultUnexportedInterface.Wrapped(something)
}

// UDUIWrappedVariadric is a wrapper around unexportedDefaultUnexportedInterface.WrappedVariadric
func UDUIWrappedVariadric(something ...string) error {
	return unexportedDefaultUnexportedInterface.WrappedVariadric(something...)
}
