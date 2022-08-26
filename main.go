package main

import (
	"log"
)

func main() {
	conf, err := initConfig()
	if err != nil {
		log.Fatalf("could not init config: %v", err)
	}

	pkg, typ, err := parseType(conf.Pkg, conf.Type)
	if err != nil {
		log.Fatalf("cannot parse type %s: %v", conf.Type, err)
	}

	parseParams := parseFieldsParams{
		format:    conf.Format,
		tag:       conf.Tag,
		tagRegex:  conf.TagRegex,
		tagFormat: conf.TagFormat,
		tagStrict: conf.TagStrict,
		embedded:  conf.Embedded,
		excluded:  conf.Excluded,
	}

	fields, err := parseFields(typ, parseParams, 0)
	if err != nil {
		log.Fatalf("cannot parse fields for type %s: %v", conf.Type, err)
	}

	if len(fields) == 0 {
		log.Fatalf("no fields found for type %s", conf.Type)
	}

	output, err := generateOutput(pkg, conf.Type, conf.Suffix, conf.TemplateFilename, fields)
	if err != nil {
		log.Fatalf("cannot generate output: %v", err)
	}

	if len(output) == 0 {
		log.Fatalf("no output generated")
	}

	err = writeToFile(output, conf.OutputFilename)
	if err != nil {
		log.Fatalf("cannot write to file %s: %v", conf.OutputFilename, err)
	}
}
