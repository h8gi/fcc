# fcc
`fcc` is a filename case converter.

## Installation

```
go get -u github.com/h8gi/fcc/...
```

## Usage

```shell
$ ls ./*.go
./case_converter.go      ./case_converter_test.go ./case_style.go          ./guess_case_style.go    ./parse.go
$ fcc -v -o pascal ./*.go
copy case_converter.go -> CaseConverter.go
copy case_converter_test.go -> CaseConverterTest.go
copy case_style.go -> CaseStyle.go
copy guess_case_style.go -> GuessCaseStyle.go
copy parse.go -> Parse.go
```

```
usage: fcc [<flags>] [<filenames>...]

A filename case converter.

Flags:
  -h, --help           Show context-sensitive help (also try --help-long and --help-man).
  -i, --incase=INCASE  Input case style. (default: auto guess)
  -o, --outcase=kebab  Output case style.
  -r, --rename         Rename files (not copy)
  -d, --dry-run        Dry run
  -v, --verbose        Verbose
      --version        Show application version.

Args:
  [<filenames>]  filenames to convert
```

## Supported Cases

- kebab-case
- snake_case
- camelCase
- PascalCase
