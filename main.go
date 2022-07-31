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

	fields, err := parseFields(typ, conf.Tag, conf.Embedded, conf.Excluded, conf.Format, 0)
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

	err = writeToFile(conf.OutputFilename, output)
	if err != nil {
		log.Fatalf("cannot write to file %s: %v", conf.OutputFilename, err)
	}
}
