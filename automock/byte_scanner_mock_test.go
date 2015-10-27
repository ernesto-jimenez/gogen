/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/gogen/automock
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package automock

import (
	"fmt"
	mock "github.com/stretchr/testify/mock"
)

// ByteScannerMock mock
type ByteScannerMock struct {
	mock.Mock
}

// ReadByte mocked method
func (m *ByteScannerMock) ReadByte() (byte, error) {

	ret := m.Called()

	var r0 byte
	switch res := ret.Get(0).(type) {
	case nil:
	case byte:
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

// UnreadByte mocked method
func (m *ByteScannerMock) UnreadByte() error {

	ret := m.Called()

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}
