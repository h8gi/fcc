package fcc

import (
	"regexp"
	"strings"
)

// parse reads a string and splits it into words.
func parse(s string, cs CaseStyle) []string {

	switch cs {
	case KebabCase:
		return strings.Split(s, "-")
	case SnakeCase:
		return strings.Split(s, "_")
	case UpperCase, LowerCase:
		return []string{s}
	case CamelCase, PascalCase:
		var r = regexp.MustCompile("[[:upper:]][[:lower:]]*")
		return r.FindAllString(strings.Title(s), -1)
	default:
		return strings.Split(s, " ")
	}
}
