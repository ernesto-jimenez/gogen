package testpkg

var (
	Str                                           = "something"
	MapKey          map[interface{}]string        = make(map[interface{}]string)
	MapValue        map[string]interface{}        = make(map[string]interface{})
	Array           []interface{}                 = make([]interface{}, 0)
	Channel         chan interface{}              = make(chan interface{})
	ROChannel       <-chan interface{}            = make(chan interface{})
	SOChannel       chan<- interface{}            = make(chan interface{})
	Var             interface{}                   = Str
	AnonymousStruct struct{ Field interface{} }   = struct{ Field interface{} }{"value"}
	AnonymousFn     func(interface{}) interface{} = func(param interface{}) interface{} { return param }
)

func Fn(param interface{}) interface{} {
	return param
}
