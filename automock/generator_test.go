package automock

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewGenerator(t *testing.T) {
	_, err := NewGenerator("io", "Writer")
	assert.NoError(t, err)
}

func TestNewGeneratorErrors(t *testing.T) {
	_, err := NewGenerator("someNonsense", "Writer")
	assert.Error(t, err)

	_, err = NewGenerator("io", "SomeWriter")
	assert.Error(t, err)
}

func TestMethods(t *testing.T) {
	g, err := NewGenerator("io", "Writer")
	assert.NoError(t, err)
	assert.Len(t, g.Methods(), 1)
}

func TestImports(t *testing.T) {
	g, err := NewGenerator("io", "Writer")
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{}, g.Imports())

	g, err = NewGenerator("net/http", "CookieJar")
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{
		"net/http": "http",
		"net/url":  "url",
	}, g.Imports())
}

func TestWritesProperly(t *testing.T) {
	tests := []struct {
		pkg   string
		iface string
	}{
		{"net/http", "CookieJar"},
		{"io", "Writer"},
		{"io", "ByteScanner"},
		{"github.com/ernesto-jimenez/gogen/automock", "unexported"},
		{".", "unexported"},
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

//go:generate go run ../cmd/goautomock/main.go io.Writer
//go:generate go run ../cmd/goautomock/main.go -pkg=io ByteScanner
//go:generate go run ../cmd/goautomock/main.go net/http.CookieJar

func TestMockedIOWriter(t *testing.T) {
	m := &WriterMock{}
	expected := []byte("hello")
	m.On("Write", expected).Return(5, nil)
	n, err := m.Write(expected)
	assert.Equal(t, 5, n)
	assert.Equal(t, nil, err)
}

func TestMockedCookieJar(t *testing.T) {
	jar := &CookieJarMock{}
	cookie := http.Cookie{Name: "hello", Value: "World"}
	jar.On("Cookies", mock.AnythingOfType("*url.URL")).Return([]*http.Cookie{&cookie}).Once()
	c := http.Client{Jar: jar}
	c.Get("http://localhost")

	jar.On("Cookies", mock.AnythingOfType("*url.URL")).Return(nil).Once()
	c.Get("http://localhost")
}

func TestMockByteScanner(t *testing.T) {
	var s io.ByteScanner
	m := &ByteScannerMock{}
	s = m
	m.On("ReadByte").Return(byte('_'), nil)
	b, err := s.ReadByte()
	assert.Equal(t, byte('_'), b)
	assert.Equal(t, nil, err)
}

type unexported interface {
	io.Reader
}

//go:generate go run ../cmd/goautomock/main.go unexported

func TestUnexported(t *testing.T) {
	m := &unexportedMock{}
	m.On("Read", mock.Anything).Return(1, nil)
	m.Read([]byte{})
}
