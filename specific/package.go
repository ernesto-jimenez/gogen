// Package specific copies the source from a package and generates a second
// package replacing some of the types used. It's aimed at taking generic
// packages that rely on interface{} and generating packages that use a
// specific type.
package specific

import (
	"encoding/json"
	"errors"
	"os/exec"
)

func findPackage(pkg string) (Package, error) {
	var p Package
	data, err := exec.Command("go", "list", "-json", pkg).CombinedOutput()
	if err != nil && len(data) > 0 {
		return p, errors.New(string(data))
	} else if err != nil {
		return p, err
	}
	err = json.Unmarshal(data, &p)
	return p, err
}

type Package struct {
	Dir         string
	GoFiles     []string
	TestGoFiles []string
}
