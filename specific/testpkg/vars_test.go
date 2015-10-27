package testpkg

import (
	"testing"
)

//go:generate go run ../../cmd/gospecific/main.go -specific-type=string -pkg=github.com/ernesto-jimenez/gogen/specific/_testpkg -out-dir=./

func TestProperTypes(t *testing.T) {
	var (
		_ map[string]string      = MapKey
		_ map[string]string      = MapValue
		_ []string               = Array
		_ chan string            = Channel
		_ <-chan string          = ROChannel
		_ chan<- string          = SOChannel
		_ struct{ Field string } = AnonymousStruct
		_ func(string) string    = Fn
	)
}
