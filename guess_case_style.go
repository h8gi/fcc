package fcc

import "strings"

// guessCaseStyle guesses case style of a string.
func guessCaseStyle(s string) CaseStyle {
	switch {
	case strings.Contains(s, "-"):
		return KebabCase
	case strings.Contains(s, "_"):
		return SnakeCase
	default:
		return PascalCase
	}
}
