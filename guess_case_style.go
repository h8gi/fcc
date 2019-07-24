package fcc

import (
	"regexp"
	"strings"
)

var (
	lower = regexp.MustCompile("^[[:lower:]]*$")
	upper = regexp.MustCompile("^[[:upper:]]*")
)

// guessCaseStyle guesses case style of a string.
func guessCaseStyle(s string) CaseStyle {
	switch {
	case strings.Contains(s, "-"):
		return KebabCase
	case strings.Contains(s, "_"):
		return SnakeCase
	case lower.MatchString(s):
		return LowerCase
	case upper.MatchString(s):
		return UpperCase
	default:
		return PascalCase
	}
}
