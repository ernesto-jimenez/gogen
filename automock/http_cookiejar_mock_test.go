/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/gogen/automock
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package automock

import (
	"fmt"
	mock "github.com/stretchr/testify/mock"

	http "net/http"
	url "net/url"
)

// CookieJarMock mock
type CookieJarMock struct {
	mock.Mock
}

// Cookies mocked method
func (m *CookieJarMock) Cookies(p0 *url.URL) []*http.Cookie {

	ret := m.Called(p0)

	var r0 []*http.Cookie
	switch res := ret.Get(0).(type) {
	case nil:
	case []*http.Cookie:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// SetCookies mocked method
func (m *CookieJarMock) SetCookies(p0 *url.URL, p1 []*http.Cookie) {

	m.Called(p0, p1)

}
