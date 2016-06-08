package strconv

import (
	"regexp"
	"strings"
)

var (
	upper = regexp.MustCompile("([A-Z])")
	under = regexp.MustCompile("_+")
)

// SnakeCase converst camel case to snake case
func SnakeCase(s string) string {
	res := upper.ReplaceAllString(s, "_$1")
	res = strings.ToLower(res)
	res = under.ReplaceAllString(res, "_")
	return strings.TrimFunc(res, func(r rune) bool {
		return r == '_'
	})
}
