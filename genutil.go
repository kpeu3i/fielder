package main

import (
	"bytes"
	"embed"
	"fmt"
	goformat "go/format"
	"os"
	"strings"
	"text/template"
)

const (
	templateName = "output.tpl"
)

var (
	//go:embed output.tpl
	templatesFS embed.FS
)

func generateOutput(pkg string, typ, suffix, tplFilename string, fields []field) ([]byte, error) {
	var (
		err error
		buf bytes.Buffer
	)

	t := template.New(templateName)

	if tplFilename == "" {
		t, err = t.ParseFS(templatesFS, templateName)
		if err != nil {
			return nil, fmt.Errorf("could not parse template: %v", err)
		}
	} else {
		t, err = t.ParseFiles(tplFilename)
		if err != nil {
			return nil, fmt.Errorf("could not parse template: %v", err)
		}
	}

	err = t.Execute(&buf, map[string]any{
		"Package": pkg,
		"Type":    typ,
		"Suffix":  suffix,
		"Fields":  fields,
		"CMD":     fmt.Sprintf("fielder %s", strings.Join(os.Args[1:], " ")),
	})
	if err != nil {
		return nil, fmt.Errorf("could not execute template: %v", err)
	}

	formatted, err := goformat.Source(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("could not format source: %v", err)
	}

	return formatted, nil
}
