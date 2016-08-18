package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/ernesto-jimenez/gogen/exportdefault"
	"github.com/ernesto-jimenez/gogen/strconv"
)

var (
	out     = flag.String("o", "", "specify the name of the generated code. Default value is by generated based on the name of the variable, e.g.: DefaultClient -> default_client_funcs.go (use \"-\" to print to stdout)")
	pref    = flag.String("prefix", "", "prefix for the exported function names")
	include = flag.String("include", "", "export only those methods that match this regexp")
	exclude = flag.String("exclude", "", "exclude those methods that match this regexp")
)

func main() {
	flag.Parse()
	log.SetFlags(0)

	// Variable whose methods we want to wrap
	v := flag.Arg(0)

	gen, err := exportdefault.New(".", v)
	if err != nil {
		log.Fatal(err)
	}

	gen.FuncNamePrefix = *pref

	if expr, err := regexp.Compile(*include); *include != "" && err == nil {
		gen.Include = expr
	} else if *include != "" {
		log.Fatalf("-include contains a invalid regular expression: %s", err.Error())
	}

	if expr, err := regexp.Compile(*include); *exclude != "" && err == nil {
		gen.Exclude = expr
	} else if *exclude != "" {
		log.Fatalf("-exclude contains a invalid regular expression: %s", err.Error())
	}

	w := os.Stdout
	if *out == "" {
		*out = fmt.Sprintf("%s_funcs.go", v)
	}
	if *out != "-" {
		log.Printf("Generating funcs for %s", v)
		*out = strconv.SnakeCase(*out)
		w, err = os.OpenFile(*out, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = gen.Write(w)
	if err != nil {
		log.Fatal(err)
	}
}
