package fcc

type CaseStyle int

const (
	UnknownCase CaseStyle = iota
	// kebab-case
	KebabCase
	// snake_case
	SnakeCase
	// camelCase
	CamelCase
	// PascalCase
	PascalCase
	// UPPERCASE
	UpperCase
	// lowercase
	LowerCase
)

func (cs CaseStyle) String() string {
	switch cs {
	case KebabCase:
		return "kebab-case"
	case SnakeCase:
		return "snake_case"
	case CamelCase:
		return "camelCase"
	case PascalCase:
		return "PascalCase"
	case UpperCase:
		return "UPPERCASE"
	case LowerCase:
		return "lowercase"
	default:
		return "UNKNOWN CASE"
	}
}
