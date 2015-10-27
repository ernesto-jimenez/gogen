/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/gogen/automock
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package automock

import (
	"fmt"
	mock "github.com/stretchr/testify/mock"

	io "io"
)

// GeneratorMock mock
type GeneratorMock struct {
	mock.Mock
}

// Imports mocked method
func (m *GeneratorMock) Imports() map[string]string {

	ret := m.Called()

	var r0 map[string]string
	switch res := ret.Get(0).(type) {
	case nil:
	case map[string]string:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// Methods mocked method
func (m *GeneratorMock) Methods() []Method {

	ret := m.Called()

	var r0 []Method
	switch res := ret.Get(0).(type) {
	case nil:
	case []Method:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// Name mocked method
func (m *GeneratorMock) Name() string {

	ret := m.Called()

	var r0 string
	switch res := ret.Get(0).(type) {
	case nil:
	case string:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// Package mocked method
func (m *GeneratorMock) Package() string {

	ret := m.Called()

	var r0 string
	switch res := ret.Get(0).(type) {
	case nil:
	case string:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// SetInternal mocked method
func (m *GeneratorMock) SetInternal(p0 bool) {

	m.Called(p0)

}

// SetName mocked method
func (m *GeneratorMock) SetName(p0 string) {

	m.Called(p0)

}

// SetPackage mocked method
func (m *GeneratorMock) SetPackage(p0 string) {

	m.Called(p0)

}

// Write mocked method
func (m *GeneratorMock) Write(p0 io.Writer) error {

	ret := m.Called(p0)

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
