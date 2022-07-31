package main

import (
	"errors"
	"flag"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
)

const (
	formatAsIs       = "as_is"
	formatSnakeCase  = "snake_case"
	formatCamelCase  = "camel_case"
	formatPascalCase = "pascal_case"
)

type config struct {
	Pkg              string
	Type             string
	Suffix           string
	Tag              string
	Format           string
	Embedded         bool
	Excluded         []string
	TemplateFilename string
	OutputFilename   string
}

func initConfig() (config, error) {
	pkg := flag.String("pkg", ".", "Package name to extract type from")
	typ := flag.String("type", "", "Type to extract fields from")
	suffix := flag.String("suffix", "Field", "Suffix for the generated struct")
	tag := flag.String("tag", "", `Use tag to extract field values from (default "")`)
	format := flag.String("format", formatAsIs, "Format of the generated type values")
	embedded := flag.Bool("embedded", false, "Extract embedded fields (default false)")
	excluded := flag.String("excluded", "", `Comma separated list of excluded fields (default "")`)
	templateFilename := flag.String("tpl", "", `Set template filename (default "")`)
	outputFilename := flag.String("output", "", `Set output filename (default "<src_dir>/<type>_fielder.go")`)

	flag.Parse()

	if *typ == "" {
		return config{}, errors.New("type is required")
	}

	switch *format {
	case formatAsIs, formatSnakeCase, formatCamelCase, formatPascalCase:
	default:
		return config{}, fmt.Errorf("invalid format %s", *format)
	}

	conf := config{}
	conf.Pkg = *pkg
	conf.Type = *typ
	conf.Suffix = *suffix
	conf.Tag = *tag
	conf.Format = *format
	conf.Embedded = *embedded
	conf.TemplateFilename = *templateFilename

	if *excluded != "" {
		conf.Excluded = strings.Split(*excluded, ",")
	}

	conf.OutputFilename = filepath.Join(".", fmt.Sprintf("%s_fielder.go", strcase.ToSnake(conf.Type)))
	if *outputFilename != "" {
		conf.OutputFilename = *outputFilename
	}

	return conf, nil
}
