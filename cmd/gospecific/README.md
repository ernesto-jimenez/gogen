# gospecific

Avoid using generic packages with `interface{}` by generating specific
packages that can be used with safe types.

# Usage

Install gospecific

```go
go get github.com/ernesto-jimenez/gogen/cmd/gospecific
```

Add a go generate comment to generate a package

```go
//go:generate gospecific -pkg=container/list -specific-type=string
```

Generate the code

```go
go generate
```

Now you will have your own `list` package to store strings rather than
`interface{}`

```sh
% godoc github.com/ernesto-jimenez/gogen/list | egrep 'func.+string'
```

```go
func (l *List) InsertAfter(v string, mark *Element) *Element
func (l *List) InsertBefore(v string, mark *Element) *Element
func (l *List) PushBack(v string) *Element
func (l *List) PushFront(v string) *Element
func (l *List) Remove(e *Element) string
```
