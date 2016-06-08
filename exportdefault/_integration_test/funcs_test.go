package integration

import (
	"testing"

	"github.com/ernesto-jimenez/gogen/exportdefault/_testpkg"
)

func TestGeneratedFuncs(t *testing.T) {
	testpkg.EDEIEmbedded()
	testpkg.EDEIWrappedVariadric("a", "b")
	testpkg.EDESMethodPtr()
	testpkg.EDESMethodVal()
	testpkg.EDESPMethodPtr()
	testpkg.EDESPMethodVal()
	testpkg.EDUIEmbedded()
	testpkg.EDUIWrappedVariadric("a", "b")
	testpkg.EDUSMethodPtr()
	testpkg.EDUSMethodVal()
	testpkg.EDUSPMethodPtr()
	testpkg.EDUSPMethodVal()
	testpkg.UDEIEmbedded()
	testpkg.UDEIWrappedVariadric("a", "b")
	testpkg.UDESMethodPtr()
	testpkg.UDESMethodVal()
	testpkg.UDESPMethodPtr()
	testpkg.UDESPMethodVal()
	testpkg.UDUIEmbedded()
	testpkg.UDUIWrappedVariadric("a", "b")
	testpkg.UDUSMethodPtr()
	testpkg.UDUSMethodVal()
	testpkg.UDUSPMethodPtr()
	testpkg.UDUSPMethodVal()
	testpkg.EDEIWrapped("a")
	testpkg.EDUIWrapped("a")
	testpkg.UDEIWrapped("a")
	testpkg.UDUIWrapped("a")
}
