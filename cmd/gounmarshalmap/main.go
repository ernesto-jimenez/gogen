package main

import (
	"flag"
	"log"
	"os"

	"github.com/ernesto-jimenez/gogen/unmarshalmap"
)

var (
	out = flag.String("o", "", "what file to write")
	pkg = flag.String("pkg", ".", "what package to get the interface from")
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

	w := os.Stdout
	if *out != "" {
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
