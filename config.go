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
	Format           string
	Tag              string
	TagRegex         string
	TagFormat        string
	TagStrict        bool
	Embedded         bool
	Excluded         []string
	TemplateFilename string
	OutputFilename   string
}

func initConfig() (config, error) {
	pkg := flag.String("pkg", ".", "Package name to extract type from")
	typ := flag.String("type", "", "Type to extract fields from")
	suffix := flag.String("suffix", "Field", "Suffix for the generated struct")
	format := flag.String("format", formatAsIs, "Format of the generated type values extracted from the struct field")
	tag := flag.String("tag", "", `Tag to extract field values from (default "")`)
	tagRegex := flag.String("tag_regex", "", `Regular expression to parse field value from a tag (default "")`)
	tagFormat := flag.String("tag_format", formatAsIs, "Format of the generated type values extracted from the tag")
	tagStrict := flag.Bool("tag_strict", false, "Strict mode for tag parsing (returns error if tag is not found)")
	embedded := flag.Bool("embedded", false, "Extract embedded fields (default false)")
	excluded := flag.String("excluded", "", `Comma separated list of excluded fields (default "")`)
	templateFilename := flag.String("tpl", "", `Set template filename (default "")`)
	outputFilename := flag.String("output", "", `Set output filename (default "<src_dir>/<type>_fielder.go")`)

	flag.Parse()

	if *typ == "" {
		return config{}, errors.New("type is required")
	}

	if *tag == "" {
		if *tagRegex != "" {
			return config{}, errors.New("tag is required when tag_regex is set")
		}

		if *tagFormat != formatAsIs {
			return config{}, errors.New("tag is required when tag_format is set")
		}

		if *tagStrict != false {
			return config{}, errors.New("tag is required when tag_strict is set")
		}
	}

	if *tagStrict && *format != formatAsIs {
		return config{}, errors.New("format is not supported when tag_strict is set")
	}

	switch *format {
	case formatAsIs, formatSnakeCase, formatCamelCase, formatPascalCase:
	default:
		return config{}, fmt.Errorf("invalid format %s", *format)
	}

	switch *tagFormat {
	case formatAsIs, formatSnakeCase, formatCamelCase, formatPascalCase:
	default:
		return config{}, fmt.Errorf("invalid tag_format %s", *format)
	}

	conf := config{}
	conf.Pkg = *pkg
	conf.Type = *typ
	conf.Suffix = *suffix
	conf.Format = *format
	conf.Tag = *tag
	conf.TagRegex = *tagRegex
	conf.TagFormat = *tagFormat
	conf.TagStrict = *tagStrict
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
