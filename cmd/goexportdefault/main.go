package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ernesto-jimenez/gogen/exportdefault"
	"github.com/ernesto-jimenez/gogen/strconv"
)

var (
	out  = flag.String("o", "", "specify the name of the generated code. Default value is by generated based on the name of the variable, e.g.: DefaultClient -> default_client_funcs.go (use \"-\" to print to stdout)")
	pref = flag.String("prefix", "", "prefix for the exported function names")
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
