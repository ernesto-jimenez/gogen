# gosimplemock

Automatically generate simple mocks with zero dependencies

# Usage

Creating an interface in your code to mock a dependency

```go
type server interface {
  Serve(string) ([]byte, error)
}

func request(s server, path string) ([]byte, error) {
  return s.Serve(path)
}

//go:generate gosimplemock server

// Dummy test
func TestRequestReturnsServerError(t *testing.T) {
  m := &requestMock{
    ServeFunc: func(path string) ([]byte, error) {
      return nil, errors.New("failure")
    },
  }
  _, err := request(m, "/something")
  assert.Error(t, err)
}
```

Mocking an interface from the standard library

```go
//go:generate gosimplemock io.Writer

// Dummy test using the generated mock
func TestWriter(t *testing.T) {
  expected := []byte("hello world")
  m := &WriterMock{
    WriteFunc: func(actual []byte) (int, error) {
      assert.Equal(t, expected, actual)
      return len(actual), nil
    },
  }
  n, err := m.Write(expected)
  assert.Equal(t, 11, n)
  assert.Equal(t, nil, err)
}
```

Printing the generated code:

```go
$ gosimplemock -o=- io.ReadCloser
/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/gogen/cmd/gosimplemock
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package gogen

// ReadCloserMock mock
type ReadCloserMock struct {
        CloseFunc func() error

        ReadFunc func([]byte) (int, error)
}

// Close mocked method
func (m *ReadCloserMock) Close() error {
        if m.CloseFunc == nil {
                panic("unexpected call to mocked method Close")
        }
        return m.CloseFunc()
}

// Read mocked method
func (m *ReadCloserMock) Read(p0 []byte) (int, error) {
        if m.ReadFunc == nil {
                panic("unexpected call to mocked method Read")
        }
        return m.ReadFunc(p0)
}
```
