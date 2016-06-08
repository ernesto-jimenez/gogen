# goexportdefault

It is a common pattern in Go packages to implement a default instance of an exported struct and export functions that call the underlying default instance.

A couple of examples from the stdlib:

 - `net/http` has `http.DefaultClient` and functions like `http.Get` just call the default `http.DefaultClient.Get`
 - `log` has `log.Logger` and functions like `log.Print` just call the default `log.std.Print`

The exported package functions simply call the corresponding methods from the default instance.

`goexportdefault` allows you to automatically generate a exported function for each method from a default struct.

# Usage

Given the following code:

```go
var DefaultClient = New()
//go:generate goexportdefault DefaultClient

type Client struct {}

func New() *Client {
  return &Client{}
}

// Do won't really do anything in this example
func (c *Client) Do(interface{}) error {
  return nil
}
```

The it will automatically generate a `default_client_funcs.go` file with the following contents:

```go
// Do is a wrapper around DefaultClient.Do
func Do(v interface{}) error {
  return DefaultClient.Do(v)
}
```
