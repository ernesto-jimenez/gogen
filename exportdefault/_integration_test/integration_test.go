package integration

import (
	"os/exec"
	"testing"
)

func TestGenerateAndBuildTestPackage(t *testing.T) {
	cmd := exec.Command("go", "generate", "github.com/ernesto-jimenez/gogen/exportdefault/_testpkg")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("error generating wrappers: %s\nOutput:\n%s", err.Error(), out)
	}

	cmd = exec.Command("go", "build", "github.com/ernesto-jimenez/gogen/exportdefault/_testpkg")
	out, err = cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("error buildinjg package: %s\nOutput:\n%s", err.Error(), out)
	}

	cmd = exec.Command("go", "test", "github.com/ernesto-jimenez/gogen/exportdefault/_testpkg")
	out, err = cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("error testing package: %s\nOutput:\n%s", err.Error(), out)
	}
}
