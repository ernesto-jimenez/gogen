package testpkg

//go:generate go run ../../cmd/goexportdefault/main.go -prefix=EDES ExportedDefaultExportedStruct

// ExportedDefaultExportedStruct to be generated
var ExportedDefaultExportedStruct = ExportedStruct{}

//go:generate go run ../../cmd/goexportdefault/main.go -prefix=UDES unexportedDefaultExportedStruct
var unexportedDefaultExportedStruct = ExportedStruct{}

//go:generate go run ../../cmd/goexportdefault/main.go -prefix=EDESP ExportedDefaultExportedStructPtr

// ExportedDefaultExportedStructPtr to be generated
var ExportedDefaultExportedStructPtr = &ExportedStruct{}

//go:generate go run ../../cmd/goexportdefault/main.go -prefix=UDESP unexportedDefaultExportedStructPtr
var unexportedDefaultExportedStructPtr = &ExportedStruct{}

// ExportedStruct is a random test struct
type ExportedStruct struct {
	embeddedStruct
}

// MethodVal docs
func (ExportedStruct) MethodVal() {}

// MethodPtr docs
func (*ExportedStruct) MethodPtr()           {}
func (ExportedStruct) unexportedMethodVal()  {}
func (*ExportedStruct) uenxportedMethodPtr() {}
