package exportdefault

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	tests := []struct {
		name    string
		include *regexp.Regexp
		exclude *regexp.Regexp
	}{
		{
			name: "simple_example",
		},
		{
			name:    "filtered",
			include: regexp.MustCompile("Wrapped.*"),
		},
		{
			name:    "excluded",
			exclude: regexp.MustCompile("Variadric"),
		},
		{
			name:    "filter_and_exclude",
			include: regexp.MustCompile("Wrapped"),
			exclude: regexp.MustCompile("Variadric"),
		},
	}
	pkg := "./_testpkg"
	variable := "ExportedDefaultExportedInterface"
	for _, test := range tests {
		g, err := New(pkg, variable)
		if err != nil {
			t.Fatalf("%s: failed initializing generator %s", test.name, err.Error())
		}
		g.Include = test.include
		g.Exclude = test.exclude
		var buf bytes.Buffer
		g.Write(&buf)
		code, err := ioutil.ReadFile(fmt.Sprintf("testdata/%s.go", test.name))
		if err != nil {
			t.Fatalf("%s: %s", test.name, err.Error())
		}
		exp := string(code)
		if buf.String() != exp {
			t.Fatalf("%s\nexpected: %s\nreturned: %s", test.name, exp, buf.String())
		}
	}
}
