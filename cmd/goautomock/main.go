package main

import (
	"flag"
	"log"
	"os"
	"path"

	"github.com/ernesto-jimenez/gogen/automock"
)

var (
	out      = flag.String("o", "", "what file to write")
	mockName = flag.String("mock-name", "", "name for the mock")
	mockPkg  = flag.String("mock-pkg", "", "package name for the mock")
	pkg      = flag.String("pkg", ".", "what package to get the interface from")
	inPkg    = flag.Bool("in-pkg", false, "whether the mock is internal to the package")
)

func main() {
	flag.Parse()
	log.SetFlags(0)

	iface := flag.Arg(0)

	if iface == "" {
		log.Fatal("need to specify an interface name")
	}

	gen, err := automock.NewGenerator(*pkg, iface)
	if err != nil {
		log.Fatal(err)
	}

	if *mockName != "" {
		gen.SetName(*mockName)
	}
	if *mockPkg != "" {
		gen.SetPackage(*mockPkg)
	}
	if *pkg == "." && path.Dir(*out) == "." {
		*inPkg = true
	}
	gen.SetInternal(*inPkg)

	w := os.Stdout
	if *out != "" {
		log.Printf("Generating mock for %s in %s", iface, *out)
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
