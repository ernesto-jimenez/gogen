package testpkg

//go:generate go run ../../cmd/goexportdefault/main.go -prefix=EDUS ExportedDefaultUnexportedStruct

// ExportedDefaultUnexportedStruct to be generated
var ExportedDefaultUnexportedStruct = unexportedStruct{}

//go:generate go run ../../cmd/goexportdefault/main.go -prefix=UDUS unexportedDefaultUnexportedStruct
var unexportedDefaultUnexportedStruct = unexportedStruct{}

//go:generate go run ../../cmd/goexportdefault/main.go -prefix=EDUSP ExportedDefaultUnexportedStructPtr

// ExportedDefaultUnexportedStructPtr to be generated
var ExportedDefaultUnexportedStructPtr = &unexportedStruct{}

//go:generate go run ../../cmd/goexportdefault/main.go -prefix=UDUSP unexportedDefaultUnexportedStructPtr
var unexportedDefaultUnexportedStructPtr = &unexportedStruct{}

type unexportedStruct struct {
	embeddedStruct
}

func (unexportedStruct) MethodVal()            {}
func (*unexportedStruct) MethodPtr()           {}
func (unexportedStruct) unexportedMethodVal()  {}
func (*unexportedStruct) uenxportedMethodPtr() {}
