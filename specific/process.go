// Package specific copies the source from a package and generates a second
// package replacing some of the types used. It's aimed at taking generic
// packages that rely on interface{} and generating packages that use a
// specific type.
package specific

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"path"
)

type Options struct {
	SkipTestFiles bool
}

var DefaultOptions = Options{
	SkipTestFiles: false,
}

// Process creates a specific package from the generic specified in pkg
func Process(pkg, outdir string, newType string, opts *Options) error {
	if opts == nil {
		opts = &DefaultOptions
	}

	p, err := findPackage(pkg)
	if err != nil {
		return err
	}

	if outdir == "" {
		outdir = path.Base(pkg)
	}

	if err := os.MkdirAll(outdir, os.ModePerm); err != nil {
		return err
	}

	files := make([]processedFile, 0)

	for _, f := range p.GoFiles {
		res, err := processFile(p, f, newType)
		if err != nil {
			return err
		}
		files = append(files, res)
	}

	if opts.SkipTestFiles {
		return write(outdir, files)
	}

	for _, f := range p.TestGoFiles {
		res, err := processFile(p, f, newType)
		if err != nil {
			return err
		}
		files = append(files, res)
	}

	return write(outdir, files)
}

func processFile(p Package, filename string, newType string) (processedFile, error) {
	res := processedFile{filename: filename}

	in, err := os.Open(path.Join(p.Dir, filename))
	if err != nil {
		return res, FileError{Package: p.Dir, File: filename, Err: err}
	}
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return res, FileError{Package: p.Dir, File: filename, Err: err}
	}

	res.fset = token.NewFileSet()
	res.file, err = parser.ParseFile(res.fset, res.filename, src, parser.ParseComments|parser.AllErrors|parser.DeclarationErrors)
	if err != nil {
		return res, FileError{Package: p.Dir, File: filename, Err: err}
	}

	ast.Walk(visitor{newType: newType}, res.file)

	return res, err
}

type visitor struct {
	newType string
}

func (v visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return v
	}
	switch n := node.(type) {
	case *ast.ArrayType:
		switch t := n.Elt.(type) {
		case *ast.InterfaceType:
			if t.Methods.NumFields() == 0 {
				str := ast.NewIdent(v.newType)
				str.NamePos = t.Pos()
				n.Elt = str
			}
		}
	case *ast.MapType:
		switch t := n.Key.(type) {
		case *ast.InterfaceType:
			if t.Methods.NumFields() == 0 {
				str := ast.NewIdent(v.newType)
				str.NamePos = t.Pos()
				n.Key = str
			}
		}
		switch t := n.Value.(type) {
		case *ast.InterfaceType:
			if t.Methods.NumFields() == 0 {
				str := ast.NewIdent(v.newType)
				str.NamePos = t.Pos()
				n.Value = str
			}
		}
	case *ast.Field:
		switch t := n.Type.(type) {
		case *ast.InterfaceType:
			if t.Methods.NumFields() == 0 {
				str := ast.NewIdent(v.newType)
				str.NamePos = t.Pos()
				n.Type = str
			}
		}
	}
	return v
}

func write(outdir string, files []processedFile) error {
	for _, f := range files {
		out, err := os.Create(path.Join(outdir, f.filename))
		if err != nil {
			return FileError{Package: outdir, File: f.filename, Err: err}
		}

		printer.Fprint(out, f.fset, f.file)
	}
	return nil
}

type FileError struct {
	Package string
	File    string
	Err     error
}

func (ferr FileError) Error() string {
	return fmt.Sprintf("error in %s: %s", path.Join(ferr.Package, ferr.File), ferr.Err.Error())
}

type processedFile struct {
	filename string
	fset     *token.FileSet
	file     *ast.File
}
