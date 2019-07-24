package fcc

import (
	"strings"
)

type Converter struct {
	// all words are lower case.
	words []string
}

func (c *Converter) Convert(s string, ics, ocs CaseStyle) string {
	c.SetWords(s, ics)
	return c.JoinWords(ocs)
}

func (c *Converter) ConvertGuess(s string, ocs CaseStyle) string {
	return c.Convert(s, guessCaseStyle(s), ocs)
}

func (c *Converter) SetWords(s string, cs CaseStyle) {
	c.words = parse(s, cs)
	c.Normalize()
}

// Normalize converts `words` to lowercase.
func (c *Converter) Normalize() {
	for i, w := range c.words {
		c.words[i] = strings.ToLower(w)
	}
}

// JoinWords generates a string according to cs.
func (c *Converter) JoinWords(cs CaseStyle) string {
	switch cs {
	case KebabCase:
		return c.KebabCase()
	case SnakeCase:
		return c.SnakeCase()
	case CamelCase:
		return c.CamelCase()
	case PascalCase:
		return c.PascalCase()
	default:
		return strings.Join(c.words, "")
	}
}

// KebabCase returns "kebab-case".
func (c *Converter) KebabCase() string {
	return strings.Join(c.words, "-")
}

// SnakeCase returns "snake_case"
func (c *Converter) SnakeCase() string {
	return strings.Join(c.words, "_")
}

// CamelCase returns "camelCase".
func (c *Converter) CamelCase() string {
	if len(c.words) < 1 {
		return ""
	}
	words := make([]string, len(c.words))
	words[0] = c.words[0]
	for i, w := range c.words[1:] {
		words[i+1] = strings.Title(w)
	}
	return strings.Join(words, "")
}

// PascalCase returns "PascalCase".
func (c *Converter) PascalCase() string {
	words := make([]string, len(c.words))
	for i, w := range c.words {
		words[i] = strings.Title(w)
	}
	return strings.Join(words, "")
}
