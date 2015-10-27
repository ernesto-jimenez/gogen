/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/gogen/automock
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package automock

import (
	"fmt"
	mock "github.com/stretchr/testify/mock"
)

// WriterMock mock
type WriterMock struct {
	mock.Mock
}

// Write mocked method
func (m *WriterMock) Write(p0 []byte) (int, error) {

	ret := m.Called(p0)

	var r0 int
	switch res := ret.Get(0).(type) {
	case nil:
	case int:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}
