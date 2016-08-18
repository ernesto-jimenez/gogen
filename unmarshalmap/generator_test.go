package unmarshalmap

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"github.com/ernesto-jimenez/gogen/unmarshalmap/testpkg"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestNewGenerator(t *testing.T) {
	_, err := NewGenerator("github.com/ernesto-jimenez/gogen/unmarshalmap/testpkg", "SimpleStruct")
	assert.NoError(t, err)
}

func TestNewGeneratorErrors(t *testing.T) {
	_, err := NewGenerator("someNonsense", "Writer")
	assert.Error(t, err)

	_, err = NewGenerator("io", "SomeWriter")
	assert.Error(t, err)
}

func TestFields(t *testing.T) {
	g, err := NewGenerator("./testpkg", "SimpleStruct")
	assert.NoError(t, err)
	assert.Len(t, g.Fields(), 7)
}

func TestImports(t *testing.T) {
	//g, err := NewGenerator("", "SimpleStruct")
	//assert.NoError(t, err)
	//assert.Equal(t, map[string]string{}, g.Imports())

	//g, err = NewGenerator("", "SimpleStruct")
	//assert.NoError(t, err)
	//assert.Equal(t, map[string]string{
	//"net/http": "http",
	//"net/url":  "url",
	//}, g.Imports())
}

func TestWritesProperly(t *testing.T) {
	tests := []struct {
		pkg   string
		iface string
	}{
		{"./testpkg", "SimpleStruct"},
	}
	for _, test := range tests {
		var out bytes.Buffer
		g, err := NewGenerator(test.pkg, test.iface)
		if err != nil {
			t.Error(err)
			continue
		}
		err = g.Write(&out)
		if !assert.NoError(t, err) {
			fmt.Println(test)
			fmt.Println(err)
			printWithLines(bytes.NewBuffer(out.Bytes()))
		}
	}
}

func printWithLines(txt io.Reader) {
	line := 0
	scanner := bufio.NewScanner(txt)
	for scanner.Scan() {
		line++
		fmt.Printf("%-4d| %s\n", line, scanner.Text())
	}
}

func TestSimpleStruct(t *testing.T) {
	var s testpkg.SimpleStruct
	expected := testpkg.SimpleStruct{
		SimpleField:             "hello",
		SimpleJSONTagged:        "second field",
		SimpleJSONTaggedOmitted: "third field",
		SimpleOmitEmptyNoName:   "noname",
	}
	m := map[string]interface{}{
		"SimpleField":   "hello",
		"field2":        "second field",
		"field3":        "third field",
		"SimpleSkipped": "skipped",
		"Ignored":       "ignore",
		"-":             "ignore",
		"SimpleOmitEmptyNoName": "noname",
	}

	err := s.UnmarshalMap(m)
	assert.NoError(t, err)
	assert.Equal(t, expected, s)

	equalJSONs(t, expected, m)
}

func equalJSONs(t assert.TestingT, exp, act interface{}) bool {
	e, err := json.Marshal(exp)
	if assert.NoError(t, err) {
		return false
	}
	a, err := json.Marshal(act)
	if assert.NoError(t, err) {
		return false
	}
	return assert.JSONEq(t, string(e), string(a))
}

func TestArrayStruct(t *testing.T) {
	var s testpkg.Array
	expected := testpkg.Array{
		List: []string{"1", "2", "3"},
	}
	m := map[string]interface{}{
		"List": []string{"1", "2", "3"},
	}

	err := s.UnmarshalMap(m)
	assert.NoError(t, err)
	assert.Equal(t, expected, s)
	equalJSONs(t, expected, m)

	s = testpkg.Array{}
	m = map[string]interface{}{}
	data, err := json.Marshal(expected)
	assert.NoError(t, err)
	json.Unmarshal(data, &m)

	err = s.UnmarshalMap(m)
	assert.NoError(t, err)
	assert.Equal(t, expected, s)
	equalJSONs(t, expected, m)

	s = testpkg.Array{}
	expected = testpkg.Array{}
	m = map[string]interface{}{}
	data, err = json.Marshal(expected)
	assert.NoError(t, err)
	json.Unmarshal(data, &m)

	err = s.UnmarshalMap(m)
	assert.NoError(t, err)
	assert.Equal(t, expected, s)
	equalJSONs(t, expected, m)
}

func TestFailWithInvalidType(t *testing.T) {
	tests := []struct {
		s unmarshalmapper
		m map[string]interface{}
	}{
		{
			&testpkg.Array{},
			map[string]interface{}{
				"List": []int{1, 2, 3},
			},
		},
		{
			&testpkg.SimpleStruct{},
			map[string]interface{}{
				"SimpleField": 12,
			},
		},
	}

	for _, test := range tests {
		err := test.s.UnmarshalMap(test.m)
		assert.Error(t, err)
	}
}

type unmarshalmapper interface {
	UnmarshalMap(map[string]interface{}) error
}

func TestNestedStruct(t *testing.T) {
	var s testpkg.Nested
	expected := testpkg.Nested{
		First:  testpkg.Embedded{Field: "first embedded"},
		Second: &testpkg.Embedded{Field: "second embedded"},
		Third:  []testpkg.Embedded{{Field: "third embedded"}},
		Fourth: []*testpkg.Embedded{&testpkg.Embedded{Field: "fourth embedded"}},
		Fifth:  [3]testpkg.Embedded{{Field: "fifth embedded"}},
		Sixth:  [3]*testpkg.Embedded{{Field: "sixth embedded"}},
	}
	m := map[string]interface{}{
		"First":  map[string]interface{}{"Field": "first embedded"},
		"Second": map[string]interface{}{"Field": "second embedded"},
		"Third":  []interface{}{map[string]interface{}{"Field": "third embedded"}},
		"Fourth": []interface{}{map[string]interface{}{"Field": "fourth embedded"}},
		"Fifth":  []interface{}{map[string]interface{}{"Field": "fifth embedded"}},
		"Sixth":  []interface{}{map[string]interface{}{"Field": "sixth embedded"}},
	}

	err := s.UnmarshalMap(m)
	assert.NoError(t, err)
	assert.Equal(t, expected, s)
	equalJSONs(t, expected, m)

	s = testpkg.Nested{}
	m = map[string]interface{}{}
	data, err := json.Marshal(expected)
	assert.NoError(t, err)
	json.Unmarshal(data, &m)

	err = s.UnmarshalMap(m)
	assert.NoError(t, err)
	t.Logf("map: %#v", m)
	t.Logf("obj: %#v", s)
	assert.Equal(t, expected, s)
	equalJSONs(t, expected, m)

	s = testpkg.Nested{}
	expected = testpkg.Nested{}
	m = map[string]interface{}{}
	data, err = json.Marshal(expected)
	assert.NoError(t, err)
	json.Unmarshal(data, &m)

	err = s.UnmarshalMap(m)
	assert.NoError(t, err)
	t.Logf("map: %#v", m)
	t.Logf("obj: %#v", s)
	assert.Equal(t, expected, s)
	equalJSONs(t, expected, m)

	s = testpkg.Nested{}
	expected = testpkg.Nested{}
	m = map[string]interface{}{}

	err = s.UnmarshalMap(m)
	assert.NoError(t, err)
	assert.Equal(t, expected, s)
	equalJSONs(t, expected, m)
}

func TestEmbeddedStruct(t *testing.T) {
	var s testpkg.Composed
	expected := testpkg.Composed{
		Embedded: testpkg.Embedded{Field: "first embedded"},
		Base:     "hello",
	}
	m := map[string]interface{}{
		"Field": "first embedded",
		"Base":  "hello",
	}

	err := s.UnmarshalMap(m)
	assert.NoError(t, err)
	assert.Equal(t, expected, s)
	equalJSONs(t, expected, m)

	s = testpkg.Composed{}
	m = map[string]interface{}{}
	data, err := json.Marshal(expected)
	assert.NoError(t, err)
	json.Unmarshal(data, &m)

	err = s.UnmarshalMap(m)
	assert.NoError(t, err)
	assert.Equal(t, expected, s)
	equalJSONs(t, expected, m)

	s = testpkg.Composed{}
	expected = testpkg.Composed{}
	m = map[string]interface{}{}
	data, err = json.Marshal(expected)
	assert.NoError(t, err)
	json.Unmarshal(data, &m)

	err = s.UnmarshalMap(m)
	assert.NoError(t, err)
	assert.Equal(t, expected, s)
	equalJSONs(t, expected, m)

	s = testpkg.Composed{}
	expected = testpkg.Composed{}
	m = map[string]interface{}{}

	err = s.UnmarshalMap(m)
	assert.NoError(t, err)
	assert.Equal(t, expected, s)
	equalJSONs(t, expected, m)
}

func BenchmarkUnmarshalmap(b *testing.B) {
	to := testpkg.SimpleStruct{}
	m := map[string]interface{}{
		"SimpleField":   "hello",
		"field2":        "second field",
		"field3":        "third field",
		"SimpleSkipped": "skipped",
		"Ignored":       "ignore",
		"-":             "ignore",
		"SimpleOmitEmptyNoName": "noname",
	}
	for i := 0; i < b.N; i++ {
		to.UnmarshalMap(m)
	}
}

func BenchmarkMapstructure(b *testing.B) {
	to := testpkg.SimpleStruct{}
	m := map[string]interface{}{
		"SimpleField":   "hello",
		"field2":        "second field",
		"field3":        "third field",
		"SimpleSkipped": "skipped",
		"Ignored":       "ignore",
		"-":             "ignore",
		"SimpleOmitEmptyNoName": "noname",
	}
	for i := 0; i < b.N; i++ {
		mapstructure.Decode(m, &to)
	}
}
