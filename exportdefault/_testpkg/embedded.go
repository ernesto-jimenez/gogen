package testpkg

type embeddedInterface interface {
	Embedded()
}

type embeddedStruct struct {
}

func (_ embeddedStruct) EmbeddedVal()            {}
func (_ *embeddedStruct) EmbeddedPtr()           {}
func (_ embeddedStruct) unexportedEmbeddedVal()  {}
func (_ *embeddedStruct) uenxportedEmbeddedPtr() {}
