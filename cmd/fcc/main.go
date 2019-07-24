package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/h8gi/fcc"
	"gopkg.in/alecthomas/kingpin.v2"
)

var caseOptions = map[string]fcc.CaseStyle{
	"kebab":  fcc.KebabCase,
	"snake":  fcc.SnakeCase,
	"camel":  fcc.CamelCase,
	"pascal": fcc.PascalCase,
}

func keys(m map[string]fcc.CaseStyle) []string {
	ks := []string{}
	for k, _ := range m {
		ks = append(ks, k)
	}
	return ks
}

var c = &fcc.Converter{}

func dstFilename(path, input, output string) string {
	dirname := filepath.Dir(path)
	filename := filepath.Base(path)
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)

	ocs := caseOptions[output]
	ics, ok := caseOptions[input]

	if ok {
		name = c.Convert(name, ics, ocs)
	} else {
		name = c.ConvertGuess(name, ocs)
	}
	return filepath.Join(dirname, name+ext)
}

func copyOrRename(srcFile, dstFile string, rename bool) error {
	if rename {
		return os.Rename(srcFile, dstFile)
	}

	input, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dstFile, input, 0644)
}

var (
	app = kingpin.New("fcc", "A filename case converter.")

	input = app.Flag("incase", "Input case style. (default: auto guess)").Short('i').
		Enum(keys(caseOptions)...)
	output = app.Flag("outcase", "Output case style.").Short('o').Default("kebab").
		Enum(keys(caseOptions)...)
	rename    = app.Flag("rename", "Rename files (not copy)").Short('r').Bool()
	dryrun    = app.Flag("dry-run", "Dry run").Short('d').Bool()
	verbose   = app.Flag("verbose", "Verbose").Short('v').Bool()
	filenames = app.Arg("filenames", "filenames to convert").Strings()
)

func main() {
	app.Version("0.0.1")
	app.HelpFlag.Short('h')
	kingpin.MustParse(app.Parse(os.Args[1:]))

	for _, srcFile := range *filenames {
		dstFile := dstFilename(srcFile, *input, *output)

		if *verbose || *dryrun {
			if *rename {
				fmt.Printf("rename %s -> %s\n", srcFile, dstFile)
			} else {
				fmt.Printf("copy %s -> %s\n", srcFile, dstFile)
			}
		}

		if !*dryrun {
			copyOrRename(srcFile, dstFile, *rename)
		}
	}
}
