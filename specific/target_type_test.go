package specific

import (
	"testing"
)

var targetTypeTests = []struct {
	src      string
	expected targetType
}{
	{
		src: "string",
		expected: targetType{
			newPkg:    "",
			newType:   "string",
			isPointer: false,
		},
	},
	{
		src: "*string",
		expected: targetType{
			newPkg:    "",
			newType:   "string",
			isPointer: true,
		},
	},
	{
		src: "*os.File",
		expected: targetType{
			newPkg:    "os",
			newType:   "os.File",
			isPointer: true,
		},
	},
	{
		src: "os.File",
		expected: targetType{
			newPkg:    "os",
			newType:   "os.File",
			isPointer: false,
		},
	},
	{
		src: "dummyexample.com/whatever/pkg.Type",
		expected: targetType{
			newPkg:    "dummyexample.com/whatever/pkg",
			newType:   "pkg.Type",
			isPointer: false,
		},
	},
}

func TestParseTargetType(t *testing.T) {
	errStr := "%s mismatch for %q\n  expected: %q\n  returned: %q"
	for _, test := range targetTypeTests {
		res := parseTargetType(test.src)
		if res.isPointer != test.expected.isPointer {
			t.Errorf(errStr, "isPointer", test.src, test.expected.isPointer, res.isPointer)
		}
		if res.newPkg != test.expected.newPkg {
			t.Errorf(errStr, "newPkg", test.src, test.expected.newPkg, res.newPkg)
		}
		if res.newType != test.expected.newType {
			t.Errorf(errStr, "newType", test.src, test.expected.newType, res.newType)
		}
	}
}
