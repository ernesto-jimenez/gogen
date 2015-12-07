package gogenutil

import (
	"os"
	"path"
	"strings"
)

// StripGopath teks the directory to a package and remove the gopath to get the
// cannonical package name
func StripGopath(p string) string {
	for _, gopath := range strings.Split(os.Getenv("GOPATH"), ":") {
		p = strings.Replace(p, path.Join(gopath, "src")+"/", "", 1)
	}
	return p
}
