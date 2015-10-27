package specific

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

func TestProcessContainerRingBuildsWithString(t *testing.T) {
	out, err := ioutil.TempDir("", "container-ring")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(out)

	Process("container/ring", out, "string", func(opts *Options) {
		opts.SkipTestFiles = true
	})

	if err := build(out); err != nil {
		t.Fatalf("failed to build resulting package\n%s", err.Error())
	}
}

func TestProcessContainerRingBuildsWithOsFile(t *testing.T) {
	out, err := ioutil.TempDir("", "container-ring")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(out)

	Process("container/ring", out, "*os.File", func(opts *Options) {
		opts.SkipTestFiles = true
	})

	if err := build(out); err != nil {
		t.Fatalf("failed to build resulting package\n%s", err.Error())
	}
}

func TestProcessContainerList(t *testing.T) {
	out, err := ioutil.TempDir("", "container-list")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(out)

	Process("container/list", out, "string", func(opts *Options) {
		opts.SkipTestFiles = true
	})

	if err := build(out); err != nil {
		t.Fatalf("failed to build resulting package\n%s", err.Error())
	}
}

func TestProcessTestPkg(t *testing.T) {
	out, err := ioutil.TempDir("", "testpkg")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(out)

	Process("github.com/ernesto-jimenez/gogen/specific/_testpkg", out, "string", func(opts *Options) {
		opts.SkipTestFiles = true
	})

	if err := build(out); err != nil {
		t.Fatalf("failed to build resulting package\n%s", err.Error())
	}
}

func build(dir string) error {
	os.Chdir(dir)
	cmd := exec.Command("go", "build")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s\n%s", err.Error(), out)
	}
	return nil
}
