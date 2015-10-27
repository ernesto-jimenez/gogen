/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/gogen/specific
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

package testpkg

var (
	Str					= "something"
	MapKey		map[string]string	= make(map[string]string)
	MapValue	map[string]string	= make(map[string]string)
	Array		[]string		= make([]string, 0)
	Channel		chan string		= make(chan string)
	ROChannel	<-chan string		= make(chan string)
	SOChannel	chan<- string		= make(chan string)
	Var		interface{}		= Str
	AnonymousStruct	struct{ Field string }	= struct{ Field string }{"value"}
	AnonymousFn	func(string) string	= func(param string) string { return param }
)

func Fn(param string) string {
	return param
}
