package fcc

import (
	"fmt"
	"testing"
)

func TestConverter(t *testing.T) {
	c := &Converter{}

	var tests = []struct {
		input       string
		inputCase   CaseStyle
		parseWant   []string
		convertWant map[CaseStyle]string
	}{
		{
			"good-bY-WoRld",
			KebabCase,
			[]string{"good", "by", "world"},
			map[CaseStyle]string{
				KebabCase:  "good-by-world",
				SnakeCase:  "good_by_world",
				CamelCase:  "goodByWorld",
				PascalCase: "GoodByWorld",
				UpperCase:  "GOODBYWORLD",
				LowerCase:  "goodbyworld",
			},
		},
		{
			"GooD_By_WoRld",
			SnakeCase,
			[]string{"good", "by", "world"},
			map[CaseStyle]string{
				KebabCase:  "good-by-world",
				SnakeCase:  "good_by_world",
				CamelCase:  "goodByWorld",
				PascalCase: "GoodByWorld",
				UpperCase:  "GOODBYWORLD",
				LowerCase:  "goodbyworld",
			},
		},
		{
			"GOODBY",
			UpperCase,
			[]string{"goodby"},
			map[CaseStyle]string{
				KebabCase:  "goodby",
				SnakeCase:  "goodby",
				CamelCase:  "goodby",
				PascalCase: "Goodby",
				UpperCase:  "GOODBY",
				LowerCase:  "goodby",
			},
		},
		{
			"helloWorld",
			CamelCase,
			[]string{"hello", "world"},
			map[CaseStyle]string{
				KebabCase:  "hello-world",
				SnakeCase:  "hello_world",
				CamelCase:  "helloWorld",
				PascalCase: "HelloWorld",
				UpperCase:  "HELLOWORLD",
				LowerCase:  "helloworld",
			},
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("c.SetWords(%v)", test.input)
		c.SetWords(test.input, test.inputCase)
		for i, v := range test.parseWant {
			if c.words[i] != v {
				t.Errorf("%s = %q, want %q", descr, c.words[i], v)
			}
		}

		for cs, want := range test.convertWant {
			descr = fmt.Sprintf("c.JoinWords(%v)", cs)
			result := c.JoinWords(cs)
			if result != want {
				t.Errorf("%s = %q, want %q", descr, result, want)
			}
		}
	}
}
