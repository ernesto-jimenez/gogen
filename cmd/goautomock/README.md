# goautomock

Automatically generate mocks

# Usage

Creating an interface in your code to mock a dependency

```go
type server interface {
  Serve(string) ([]byte, error)
}

func request(s server, path string) ([]byte, error) {
  return s.Serve(path)
}

//go:generate goautomock server

// Dummy test
func TestRequestReturnsServerError(t *testing.T) {
  m := &requestMock{}
  m.On("Serve", "/something").Return(nil, errors.New("failure"))
  _, err := request(m, "/something")
  assert.Error(t, err)
}
```

Mocking an interface from the standard library

```go
//go:generate goautomock io.Writer

// Dummy test using the generated mock
func TestWriter(t *testing.T) {
  m := &WriterMock{}
  expected := []byte("hello world")
  m.On("Write", expected).Return(11, nil)

  n, err := m.Write(expected)
  assert.Equal(t, 11, n)
  assert.Equal(t, nil, err)
}
```

Printing the generated code:

```go
$ goautomock -o=- io.ReadCloser
/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/gogen/automock
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package gogen

import (
  "fmt"
  mock "github.com/stretchr/testify/mock"
)

// ReadCloserMock mock
type ReadCloserMock struct {
  mock.Mock
}

// Close mocked method
func (m *ReadCloserMock) Close() error {

  ret := m.Called()

  var r0 error
  switch res := ret.Get(0).(type) {
  case nil:
  case error:
    r0 = res
  default:
    panic(fmt.Sprintf("unexpected type: %v", res))
  }

  return r0

}

// Read mocked method
func (m *ReadCloserMock) Read(p0 []byte) (int, error) {

  ret := m.Called(p0)

  var r0 int
  switch res := ret.Get(0).(type) {
  case nil:
  case int:
    r0 = res
  default:
    panic(fmt.Sprintf("unexpected type: %v", res))
  }

  var r1 error
  switch res := ret.Get(1).(type) {
  case nil:
  case error:
    r1 = res
  default:
    panic(fmt.Sprintf("unexpected type: %v", res))
  }

  return r0, r1

}
```
