package testunmarshalmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Debug indicates we should log debugging information
var Debug bool

type unmarshalMapper interface {
	UnmarshalMap(map[string]interface{}) error
}

// Run tests that UnmarshalMap unmarshalls all fields
func Run(t *testing.T, v unmarshalMapper) {
	// make an empty variable of the same type as v
	n := empty(t, v)

	// fill the variable with data
	fill(t, v)

	// generate a map[string]interface{}
	m := generateMap(t, v)

	// unmarshal generated map into the empty variable
	n.UnmarshalMap(m)

	require.Equal(t, v, n, "UnmarshalMap() method from %T out of date. regenerate the code", v)
}

func empty(t *testing.T, v unmarshalMapper) unmarshalMapper {
	n := makeNew(t, v)
	u, ok := n.(unmarshalMapper)
	if !ok {
		t.Fatalf("%T should implement UnmarshalMap", n)
	}
	return u
}

func generateMap(t *testing.T, v unmarshalMapper) map[string]interface{} {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(v)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]interface{}
	err = json.NewDecoder(&buf).Decode(&m)
	if err != nil {
		t.Fatal(err)
	}
	return m
}

func fill(t *testing.T, v interface{}) {
	val := reflect.ValueOf(v)
	fillReflect(t, "", val)
}

func makeNew(t *testing.T, v interface{}) interface{} {
	typ := reflect.TypeOf(v).Elem()
	val := reflect.New(typ)
	return val.Interface()
}

func fillReflect(t *testing.T, scope string, val reflect.Value) {
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()
	if !val.IsValid() {
		t.Fatalf("invalid")
	}
	if Debug {
		t.Logf("%s %s", scope, typ.Kind())
	}

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		// Skip omitted fields
		if tag := strings.Split(f.Tag.Get("json"), ","); len(tag) > 0 && tag[0] == "-" {
			continue
		}
		if Debug {
			t.Logf("%s%s %s", scope, f.Name, f.Type.Kind())
		}
		v := val.Field(i)
		fillValue(t, scope+f.Name+".", v)
	}
}

func fillValue(t *testing.T, scope string, v reflect.Value) {
	kind := v.Type().Kind()
	if kind == reflect.Ptr {
		v.Set(reflect.New(v.Type().Elem()))
		v = v.Elem()
		kind = v.Kind()
	}
	switch kind {
	case reflect.Ptr:
		t.Fatalf("%s should have been de-referenced", scope)
	case reflect.Interface, reflect.Chan, reflect.Func, reflect.Complex64, reflect.Complex128, reflect.Uintptr:
		t.Fatalf("%s cannot unmarshall %s", scope, kind)
	case reflect.Bool:
		v.SetBool(true)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(newInt())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(newUint())
	case reflect.Float32, reflect.Float64:
		v.SetFloat(newFloat())
	case reflect.Array:
		l := v.Cap()
		for i := 0; i < l; i++ {
			fillValue(t, fmt.Sprintf("%s[%d].", scope, i), v.Index(i))
		}
	case reflect.Slice:
		l := rand.Intn(5) + 1
		s := reflect.MakeSlice(v.Type(), l, l)
		v.Set(s)
		for i := 0; i < l; i++ {
			fillValue(t, fmt.Sprintf("%s[%d].", scope, i), s.Index(i))
		}
	// case reflect.Map:
	// 	t.Fatalf("%s unmarshalling maps is still pending", scope)
	case reflect.String:
		v.SetString(newStr())
	case reflect.Struct:
		fillReflect(t, scope, v)
	default:
		t.Fatalf("gounmarshalmap is missing support for %s unmarshalling", kind)
	}
}

var i uint64

func newInt() int64 {
	i++
	return int64(i)
}

func newFloat() float64 {
	i++
	return float64(i)
}

func newUint() uint64 {
	i++
	return i
}

func newStr() string {
	i++
	return strconv.FormatInt(int64(i), 10)
}
