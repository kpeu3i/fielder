package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var parseFieldsData = []struct {
	pkg            string
	typ            string
	format         string
	tag            string
	tagRegex       string
	tagFormat      string
	tagStrict      bool
	embedded       bool
	excluded       []string
	expectedFields []field
}{
	{
		pkg:       "./internal/testdata",
		typ:       "User",
		format:    formatAsIs,
		tag:       "db",
		tagFormat: formatAsIs,
		embedded:  true,
		excluded:  []string{"FullName"},
		expectedFields: []field{
			{Name: "CreatedAt", Alias: "created_at", Depth: 2},
			{Name: "UpdatedAt", Alias: "updated_at", Depth: 2},
			{Name: "FirstUpdatedAt", Alias: "FirstUpdatedAt", Depth: 2},
			{Name: "DeletedAt", Alias: "deleted_at", Depth: 3},
			{Name: "ID", Alias: "id", Depth: 1},
			{Name: "FirstName", Alias: "name", Depth: 1},
			{Name: "LastName", Alias: "surname", Depth: 1},
			{Name: "Email", Alias: "email", Depth: 1},
			{Name: "Password", Alias: "password", Depth: 1},
		},
	},
	{
		pkg:       "./internal/testdata",
		typ:       "User",
		format:    formatAsIs,
		tag:       "db",
		tagRegex:  `(\w+),?.*`,
		tagFormat: formatAsIs,
		embedded:  true,
		excluded:  []string{"FullName"},
		expectedFields: []field{
			{Name: "CreatedAt", Alias: "created_at", Depth: 2},
			{Name: "UpdatedAt", Alias: "updated_at", Depth: 2},
			{Name: "FirstUpdatedAt", Alias: "FirstUpdatedAt", Depth: 2},
			{Name: "DeletedAt", Alias: "deleted_at", Depth: 3},
			{Name: "ID", Alias: "id", Depth: 1},
			{Name: "FirstName", Alias: "name", Depth: 1},
			{Name: "LastName", Alias: "surname", Depth: 1},
			{Name: "Email", Alias: "email", Depth: 1},
			{Name: "Password", Alias: "password", Depth: 1},
		},
	},
	{
		pkg:      "./internal/testdata",
		typ:      "User",
		format:   formatCamelCase,
		tag:      "",
		embedded: true,
		excluded: []string{},
		expectedFields: []field{
			{Name: "CreatedAt", Alias: "createdAt", Depth: 2},
			{Name: "UpdatedAt", Alias: "updatedAt", Depth: 2},
			{Name: "FirstUpdatedAt", Alias: "firstUpdatedAt", Depth: 2},
			{Name: "DeletedAt", Alias: "deletedAt", Depth: 3},
			{Name: "ID", Alias: "id", Depth: 1},
			{Name: "FirstName", Alias: "firstName", Depth: 1},
			{Name: "LastName", Alias: "lastName", Depth: 1},
			{Name: "Email", Alias: "email", Depth: 1},
			{Name: "Password", Alias: "password", Depth: 1},
			{Name: "FullName", Alias: "fullName", Depth: 1},
		},
	},
	{
		pkg:      "./internal/testdata",
		typ:      "User",
		format:   formatPascalCase,
		tag:      "",
		embedded: true,
		excluded: []string{},
		expectedFields: []field{
			{Name: "CreatedAt", Alias: "CreatedAt", Depth: 2},
			{Name: "UpdatedAt", Alias: "UpdatedAt", Depth: 2},
			{Name: "FirstUpdatedAt", Alias: "FirstUpdatedAt", Depth: 2},
			{Name: "DeletedAt", Alias: "DeletedAt", Depth: 3},
			{Name: "ID", Alias: "Id", Depth: 1},
			{Name: "FirstName", Alias: "FirstName", Depth: 1},
			{Name: "LastName", Alias: "LastName", Depth: 1},
			{Name: "Email", Alias: "Email", Depth: 1},
			{Name: "Password", Alias: "Password", Depth: 1},
			{Name: "FullName", Alias: "FullName", Depth: 1},
		},
	},
	{
		pkg:      "./internal/testdata",
		typ:      "User",
		format:   formatAsIs,
		tag:      "",
		embedded: false,
		excluded: []string{},
		expectedFields: []field{
			{Name: "ID", Alias: "ID", Depth: 1},
			{Name: "FirstName", Alias: "FirstName", Depth: 1},
			{Name: "LastName", Alias: "LastName", Depth: 1},
			{Name: "Email", Alias: "Email", Depth: 1},
			{Name: "Password", Alias: "Password", Depth: 1},
			{Name: "FullName", Alias: "FullName", Depth: 1},
		},
	},
}

func TestParseFields(t *testing.T) {
	for _, test := range parseFieldsData {
		t.Run(test.typ, func(t *testing.T) {
			_, typ, err := parseType(test.pkg, test.typ)
			require.NoError(t, err)

			parseParams := parseFieldsParams{
				format:    test.format,
				tag:       test.tag,
				tagRegex:  test.tagRegex,
				tagFormat: test.tagFormat,
				tagStrict: test.tagStrict,
				embedded:  test.embedded,
				excluded:  test.excluded,
			}

			actualFields, err := parseFields(typ, parseParams, 0)
			require.NoError(t, err)

			assert.Equal(t, test.expectedFields, actualFields)
		})
	}
}
