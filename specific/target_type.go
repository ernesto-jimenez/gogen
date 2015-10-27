package specific

import (
	"path"
	"strings"
)

type targetType struct {
	newPkg    string
	newType   string
	isPointer bool
}

func parseTargetType(fullType string) targetType {
	t := targetType{}
	deref := strings.TrimLeft(fullType, "*")
	if len(deref) != len(fullType) {
		t.isPointer = true
	}
	t.newType = path.Base(deref)
	t.newPkg = pkg(deref)
	return t
}

func pkg(fullType string) string {
	for i := len(fullType) - 1; i >= 0; i-- {
		if fullType[i] == '.' {
			return fullType[0:i]
		}
	}
	return ""
}
