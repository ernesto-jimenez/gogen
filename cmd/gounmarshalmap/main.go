package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/ernesto-jimenez/gogen/unmarshalmap"
)

var (
	out  = flag.String("o", "", "what file to write")
	tOut = flag.String("o-test", "", "what file to write the test to")
	pkg  = flag.String("pkg", ".", "what package to get the interface from")
)

func main() {
	flag.Parse()
	log.SetFlags(0)

	st := flag.Arg(0)

	if st == "" {
		log.Fatal("need to specify a struct name")
	}

	gen, err := unmarshalmap.NewGenerator(*pkg, st)
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer

	log.Printf("Generating func (*%s) UnmarshalMap(map[string]interface{}) error", st)
	err = gen.Write(&buf)
	if err != nil {
		log.Fatal(err)
	}

	if *out != "" {
		err := ioutil.WriteFile(*out, buf.Bytes(), 0666)
		if err != nil {
			log.Fatal(err)
		}
		if *tOut == "" {
			*tOut = fmt.Sprintf("%s_test.go", strings.TrimRight(*out, ".go"))
		}
	} else {
		fmt.Println(buf.String())
	}

	buf.Reset()

	err = gen.WriteTest(&buf)
	if err != nil {
		log.Fatal(err)
	}

	if *tOut != "" {
		err := ioutil.WriteFile(*tOut, buf.Bytes(), 0666)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(buf.String())
	}
}
