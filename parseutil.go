package main

import (
	"errors"
	"fmt"
	"go/types"

	"github.com/fatih/structtag"
	"github.com/iancoleman/strcase"
	"github.com/samber/lo"
	"golang.org/x/tools/go/packages"
)

type field struct {
	Name  string
	Alias string
	Depth int
}

func parseType(path, name string) (string, types.Type, error) {
	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, path)
	if err != nil {
		return "", nil, err
	}

	err = validatePackages(pkgs)
	if err != nil {
		return "", nil, err
	}

	obj := pkgs[0].Types.Scope().Lookup(name)
	if obj == nil || obj.Type() == nil {
		return "", nil, fmt.Errorf("type %s not found", name)
	}

	return obj.Pkg().Name(), obj.Type(), nil
}

func parseFields(
	typ types.Type,
	useTag string,
	useEmbedded bool,
	excluded []string,
	format string,
	depth int,
) ([]field, error) {
	strct, ok := typ.Underlying().(*types.Struct)
	if !ok {
		return nil, fmt.Errorf("type %s is not a struct", typ)
	}

	depth++

	var fields []field

	for i := 0; i < strct.NumFields(); i++ {
		if !strct.Field(i).IsField() && !strct.Field(i).Exported() {
			continue
		}

		isExcluded := lo.Contains(excluded, strct.Field(i).Name())
		if isExcluded {
			continue
		}

		if strct.Field(i).Embedded() {
			if !useEmbedded {
				continue
			}

			embeddedFields, err := parseFields(strct.Field(i).Type(), useTag, useEmbedded, excluded, format, depth)
			if err != nil {
				return nil, err
			}

			fields = append(fields, embeddedFields...)
		} else {
			alias := strct.Field(i).Name()
			if useTag != "" {
				tags, err := structtag.Parse(strct.Tag(i))
				if err != nil {
					return nil, err
				}

				found := false
				for _, t := range tags.Tags() {
					if t.Key == useTag {
						alias = t.Name
						found = true

						break
					}
				}

				if !found {
					return nil, fmt.Errorf("tag %s not found for field %s", useTag, strct.Field(i).Name())
				}
			}

			switch format {
			case formatAsIs:
			case formatSnakeCase:
				alias = strcase.ToSnake(alias)
			case formatCamelCase:
				alias = strcase.ToLowerCamel(alias)
			case formatPascalCase:
				alias = strcase.ToCamel(alias)
			default:
				return nil, fmt.Errorf("invalid format %s", format)
			}

			fields = append(fields, field{
				Name:  strct.Field(i).Name(),
				Alias: alias,
				Depth: depth,
			})
		}
	}

	// The fields with the lowest depth are the most important.
	for l := 0; l < len(fields); l++ {
		for k := 0; k < len(fields); k++ {
			if fields[l].Name == fields[k].Name && l != k {
				if fields[l].Depth > fields[k].Depth {
					fields = append(fields[:l], fields[l+1:]...)
				} else {
					fields = append(fields[:k], fields[k+1:]...)
				}
			}
		}
	}

	return fields, nil
}

func validatePackages(pkgs []*packages.Package) error {
	var err error
	packages.Visit(pkgs, nil, func(pkg *packages.Package) {
		for _, e := range pkg.Errors {
			err = e
		}
	})
	if err != nil {
		return err
	}

	if len(pkgs) == 0 {
		return errors.New("no packages found")
	}

	if len(pkgs) > 1 {
		return fmt.Errorf("multiple packages found")
	}

	return nil
}
