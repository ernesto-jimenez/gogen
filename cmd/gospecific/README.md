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
//go:generate -pkg=container/list interface{}:CustomType
```

Generate the code

```go
go generate
```

Now you will have your own list package to store a list of `CustomType`
elements.
